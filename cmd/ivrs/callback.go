// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ivrs

import (
	"encoding/json"
	"fmt"

	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/matthieukhl/rgvr/internal/flags"
	"github.com/matthieukhl/rgvr/internal/models"
	"github.com/spf13/cobra"
)

// callbackCmd represents the callback command
var callbackCmd = &cobra.Command{
	Use: `callback <ivr_id>
		--from <from_number> --to <to_number>
		[--clir] [--timeout <seconds>]
		[--object <salesforce_object>] [--object-id <salesforce_record_id>]`,
	Short: "Triggers a callback through a specific IVR (Interactive Voice Response) queue.",
	Long: `Triggers a callback through a specific IVR (Interactive Voice Response) queue.
The specified phone number will be called, and when the recipient picks up, they are routed through
the IVR flow associated with the given ivrId. This is typically used to offer a "call me back" feature
to customers waiting in a queue.

IVR Callbacks are only available on the Business and Advanced plans.
Standard telephony charges apply based on the destination number.

Permission:

	Calls Write required.

Monitoring:
	
	Not required for this route. Any valid API key with Calls Write permission can initiate
	an IVR callback regardless of the Monitoring flag.

`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ivrID := args[0]

		// Retrieve required flags
		fromNumber, _ := cmd.Flags().GetInt64("from")
		toNumber, _ := cmd.Flags().GetInt64("to")
		clir, _ := cmd.Flags().GetBool("clir")
		timeout, _ := cmd.Flags().GetInt64("timeout")

		// Validate timeout value
		if timeout < 20 || timeout > 300 {
			return fmt.Errorf("timeout must be between 20 and 300 seconds, got %d", timeout)
		}

		// Create the base IVRCallback struct with the provided parameters to pass a request body to the API
		ivrCallback := models.IVRCallback{
			FromNumber: fromNumber,
			ToNumber:   toNumber,
			Clir:       clir,
			Timeout:    timeout,
		}

		// Retrieve optional flags
		object, _ := cmd.Flags().GetString("object")
		objectID, _ := cmd.Flags().GetString("object_id")

		// If the optional flags are provided, add them to the IVRCallback struct
		if object != "" {
			ivrCallback.Integrations = &models.Integrations{
				Salesforce: &models.SalesforceIntegration{
					Object:   object,
					ObjectID: objectID,
				},
			}
		}

		httpClient := cmd.Context().Value(client.ClientContextKey).(*client.Client)

		callbackResp, reqInfo, err := httpClient.PostCallback(ivrID, ivrCallback)
		if err != nil {
			return err
		}

		if err := flags.IsVerbose(cmd, reqInfo); err != nil {
			return err
		}

		// Print the response in JSON format
		jsonResp, err := json.MarshalIndent(callbackResp, "", "  ")
		if err != nil {
			return err
		}

		cmd.Printf("\nCallback triggered successfully.\n\n")
		cmd.Println(string(jsonResp))

		return nil
	},
}

func init() {
	ivrsCmd.AddCommand(callbackCmd)

	// Required Flags
	callbackCmd.Flags().Int64("from", 0, "The IVR Number to display")
	callbackCmd.Flags().Int64("to", 0, "The phone number to call to")
	callbackCmd.Flags().Bool("clir", false, "Calling line identification restriction: The CLIR service blocks calling party address information from being presented to the called user. Default false (i.e. the calling number will be displayed).")
	callbackCmd.Flags().Int64("timeout", 20, "Number of seconds before the call is aborted if the communication isn't etablished. Must be between 20 and 300. Default 20 seconds.")

	callbackCmd.MarkFlagRequired("from")
	callbackCmd.MarkFlagRequired("to")

	// Optional Flags
	callbackCmd.Flags().String("object", "", "Salesforce object type to link this callback to (e.g. \"Opportunity\", \"Lead\"). Requires --object-id.")
	callbackCmd.Flags().String("object_id", "", "Salesforce record ID to link this callback to. Requires --object.")
}

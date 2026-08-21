// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package flags

import (
	"fmt"
	"os"

	"github.com/matthieukhl/rgvr/internal/client"
	"github.com/spf13/cobra"
)

func IsVerbose(cmd *cobra.Command, requestInfo *client.RequestInfo) error {
	verbose, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		return err
	}

	if verbose {
		fmt.Fprintf(os.Stderr, "\nURL called: %s\n", requestInfo.URL)
		fmt.Fprintf(os.Stderr, "Query duration: %d ms\n", requestInfo.Duration.Milliseconds())
	}

	return nil
}

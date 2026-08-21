// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/matthieukhl/rgvr/internal/models"
)

func (c *Client) GetScenarios() (*models.ListResponse[models.Scenario], *RequestInfo, error) {
	path := "/scenarios"

	resp, reqInfo, err := c.Get(path)
	if err != nil {
		return nil, reqInfo, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 401 {
		return nil, reqInfo, fmt.Errorf("%s: invalid API key or missing 'IVRs Read' permission.", resp.Status)
	}

	// This has been commented out because although the documentation
	// of Ringover's public API mentions a 204 in case no scenarios are found,
	// in practice it is not the case. It returns a 200 with an empty body.
	// I decided to keep the code commented for when this will be fixed and meanwhile,
	// I'll simply check the body sent back by the API.
	// if resp.StatusCode == 204 {
	// 	return nil, reqInfo, nil
	// }
	// Temporary workaround
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, reqInfo, fmt.Errorf("reading response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, reqInfo, fmt.Errorf("unexpected response from API: %s", resp.Status)
	}

	if len(bodyBytes) == 0 {
		return &models.ListResponse[models.Scenario]{}, reqInfo, nil
	}

	var scenariosResponse models.ListResponse[models.Scenario]
	if err := json.Unmarshal(bodyBytes, &scenariosResponse); err != nil {
		return nil, reqInfo, fmt.Errorf("decoding scenarios information: %w", err)
	}

	return &scenariosResponse, reqInfo, nil
}

func (c *Client) GetScenario(scenarioID string) (*models.Scenario, *RequestInfo, error) {
	path := fmt.Sprintf("/scenarios/%s", scenarioID)

	resp, reqInfo, err := c.Get(path)
	if err != nil {
		return nil, reqInfo, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return nil, reqInfo, fmt.Errorf("%s: invalid scenario ID format", resp.Status)
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, reqInfo, fmt.Errorf("%s: invalid API key or missing 'IVRs Read' permission", resp.Status)
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, reqInfo, fmt.Errorf("%s: scenario not found or not accessible with current monitoring level", resp.Status)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, reqInfo, fmt.Errorf("unexpected response from API: %s", resp.Status)
	}

	var scenario *models.Scenario

	err = json.NewDecoder(resp.Body).Decode(&scenario)
	if err != nil {
		return nil, reqInfo, fmt.Errorf("decoding scenario information: %w", err)
	}

	return scenario, reqInfo, nil
}

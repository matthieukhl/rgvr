/*
		rgvr - A CLI to interact with Ringover's public API.
	    Copyright (C) 2026  Matthieu Khairallah <matthieu.khairallah@proton.me>

	    This program is free software: you can redistribute it and/or modify
	    it under the terms of the GNU Affero General Public License as published by
	    the Free Software Foundation, either version 3 of the License, or
	    (at your option) any later version.

	    This program is distributed in the hope that it will be useful,
	    but WITHOUT ANY WARRANTY; without even the implied warranty of
	    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	    GNU Affero General Public License for more details.

	    You should have received a copy of the GNU Affero General Public License
	    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/matthieukhl/rgvr/internal/models"
)

func (c *Client) GetIVRs() (*models.ListResponse[models.IVR], *RequestInfo, error) {
	path := "/ivrs"

	resp, reqInfo, err := c.Get(path)
	if err != nil {
		return nil, reqInfo, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 401 {
		return nil, reqInfo, fmt.Errorf("unauthorized: invalid or missing API key, or missing 'IVRs read' permission")
	}

	// This has been commented out because although the documentation
	// of Ringover's public API mentions a 204 in case no IVRs are found,
	// in practice it is not the case. It returns a 200 with an empty body.
	// I decided to keep the code commented for when this will be fixed and meanwhile,
	// I'll simply check the body sent back by the API.
	// if resp.StatusCode == 204 {
	// 	return nil, reqInfo, nil
	// }
	// Temporary workaroud
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, reqInfo, fmt.Errorf("reading response body: %w", err)
	}

	if resp.StatusCode != 200 {
		return nil, reqInfo, fmt.Errorf("unexpected response from API: %s", resp.Status)
	}

	if len(bodyBytes) == 0 {
		return &models.ListResponse[models.IVR]{}, reqInfo, nil
	}

	var ivrsResponse models.ListResponse[models.IVR]
	if err := json.Unmarshal(bodyBytes, &ivrsResponse); err != nil {
		return nil, reqInfo, fmt.Errorf("decoding IVR information: %w", err)
	}

	return &ivrsResponse, reqInfo, nil
}

func (c *Client) GetIVR(ivrId string) (*models.IVR, *RequestInfo, error) {
	path := fmt.Sprintf("/ivrs/%s", ivrId)

	resp, reqInfo, err := c.Get(path)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 400 {
		return nil, nil, fmt.Errorf("invalid IVR identifier: %s", resp.Status)
	}

	if resp.StatusCode == 401 {
		return nil, nil, fmt.Errorf("unauthorized: invalid or missing API token, or missing 'IVRs read' permission: %w", err)
	}

	if resp.StatusCode == 404 {
		return nil, nil, fmt.Errorf("IVR not found, or not accessible with current monitoring level: %s", resp.Status)
	}

	if resp.StatusCode != 200 {
		return nil, nil, fmt.Errorf("unexpected response from API: %s", resp.Status)
	}

	var ivr *models.IVR

	err = json.NewDecoder(resp.Body).Decode(&ivr)
	if err != nil {
		return nil, nil, fmt.Errorf("decoding number information: %w", err)
	}

	return ivr, reqInfo, nil
}

func (c *Client) GetIVRScenarios(ivrId string) (*models.ListResponse[models.Scenario], *RequestInfo, error) {
	path := fmt.Sprintf("/ivrs/%s/scenarios", ivrId)

	resp, reqInfo, err := c.Get(path)
	if err != nil {
		return nil, reqInfo, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 401 {
		return nil, reqInfo, fmt.Errorf("unauthorized: invalid or missing API key, or missing 'IVRs read' permission")
	}

	if resp.StatusCode == 404 {
		return nil, reqInfo, fmt.Errorf("IVR not found, or not accessible with current monitoring level: %s", resp.Status)
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

	if resp.StatusCode != 200 {
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

func (c *Client) GetIVRScenario(ivrID, scenarioID string) (*models.Scenario, *RequestInfo, error) {
	path := fmt.Sprintf("/ivrs/%s/scenarios/%s", ivrID, scenarioID)

	resp, reqInfo, err := c.Get(path)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return nil, nil, fmt.Errorf("invalid IVR or scenario identifier: %s", resp.Status)
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, nil, fmt.Errorf("unauthorized: invalid or missing API token, or missing 'IVRs Read' permission: %s", resp.Status)
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil, fmt.Errorf("Scenario not found, or not accessible with current monitoring level.")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("unexpected response from API: %w", err)
	}

	var scenario *models.Scenario

	err = json.NewDecoder(resp.Body).Decode(&scenario)
	if err != nil {
		return nil, reqInfo, fmt.Errorf("deocding scenario information: %w", err)
	}

	return scenario, reqInfo, nil
}

func (c *Client) PostCallback(ivrID string, ivrCallback models.IVRCallback) (*models.IVRCallbackResponse, *RequestInfo, error) {
	path := fmt.Sprintf("/ivrs/%s/callback", ivrID)

	body, err := json.Marshal(ivrCallback)
	if err != nil {
		return nil, nil, fmt.Errorf("encoding request body: %w", err)
	}

	resp, reqInfo, err := c.Post(path, body)
	if err != nil {
		return nil, reqInfo, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, reqInfo, fmt.Errorf("%s: missing or invalid token, or insufficient 'Calls write' permission")
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, reqInfo, fmt.Errorf("%s: the specified IVR does not exist")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, reqInfo, fmt.Errorf("unexpected response from API: %w", err)
	}

	var ivrCallbackResp *models.IVRCallbackResponse

	if err := json.NewDecoder(resp.Body).Decode(&ivrCallbackResp); err != nil {
		return nil, reqInfo, fmt.Errorf("deocding scenario information: %w", err)
	}

	return ivrCallbackResp, reqInfo, nil
}

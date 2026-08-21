// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/matthieukhl/rgvr/internal/models"
)

func (c *Client) GetConferences() (*models.ListResponse[models.Conference], *RequestInfo, error) {
	path := "/conferences"

	resp, reqInfo, err := c.Get(path)
	if err != nil {
		return nil, reqInfo, fmt.Errorf("reading response body: %w", err)
	}
	defer resp.Body.Close()

	// This endpoint actually returnss a 204 in case of an empty response from the server!
	if resp.StatusCode == http.StatusNoContent {
		return nil, reqInfo, nil
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, reqInfo, fmt.Errorf("%s: invalid or missing API token, or Monitoring is OFF.", resp.Status)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, reqInfo, fmt.Errorf("unexpected response from API: %w", err)
	}

	var conferences *models.ListResponse[models.Conference]

	err = json.NewDecoder(resp.Body).Decode(&conferences)
	if err != nil {
		return nil, reqInfo, fmt.Errorf("decoding conferences information: %w", err)
	}

	return conferences, reqInfo, nil
}

func (c *Client) GetConference(conferenceID int) (*models.Conference, *RequestInfo, error) {
	path := fmt.Sprintf("/conferences/%d", conferenceID)

	resp, reqInfo, err := c.Get(path)
	if err != nil {
		return nil, reqInfo, fmt.Errorf("reading response bodyy: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return nil, reqInfo, fmt.Errorf("%s: invalid conference identifier", resp.Status)
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, reqInfo, fmt.Errorf("%s: invalid or missing API token, or Monitoring is OFF.", resp.Status)
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, reqInfo, fmt.Errorf("%s: conference not found", resp.Status)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, reqInfo, fmt.Errorf("unexpected response from API: %w", err)
	}

	var conference *models.Conference

	if err := json.NewDecoder(resp.Body).Decode(&conference); err != nil {
		return nil, reqInfo, fmt.Errorf("decoding conferences information: %w", err)
	}

	return conference, reqInfo, nil
}

func (c *Client) SetConferencePin(conferenceID, pinCode int) (*RequestInfo, error) {
	path := fmt.Sprintf("/conferences/%d/pincode/%d", conferenceID, pinCode)

	resp, reqInfo, err := c.Patch(path, nil)
	if err != nil {
		return reqInfo, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotModified {
		return reqInfo, nil
	}

	if resp.StatusCode == http.StatusBadRequest {
		return reqInfo, fmt.Errorf("%s: invalid conference or pincode identifier", resp.Status)
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return reqInfo, fmt.Errorf("%s: invalid or missing API token, or Monitoring is OFF", resp.Status)
	}

	if resp.StatusCode == http.StatusNotFound {
		return reqInfo, fmt.Errorf("%s: conference not found", resp.Status)
	}

	if resp.StatusCode != http.StatusOK {
		return reqInfo, fmt.Errorf("unexpected reponse from API: %s", resp.Status)
	}

	return reqInfo, nil
}

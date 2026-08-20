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

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

func (c *Client) InviteUser(invitations *models.UserInvitationPayload) ([]models.UserInvitationResponse, *RequestInfo, error) {
	path := "/users/invite"

	body, err := json.Marshal(invitations)
	if err != nil {
		return nil, nil, fmt.Errorf("encoding request payload: %w", err)
	}

	resp, reqInfo, err := c.Post(path, body)
	if err != nil {
		return nil, reqInfo, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return nil, reqInfo, fmt.Errorf("%s: the request body is malformed or missing required fields", resp.Status)
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, reqInfo, fmt.Errorf("%s: missing, invalid or expired API key", resp.Status)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, reqInfo, fmt.Errorf("unexpected response from API: %s", resp.Status)
	}

	var result []models.UserInvitationResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, reqInfo, fmt.Errorf("decoding response body: %w", err)
	}

	return result, reqInfo, nil
}

func (c *Client) DeleteUser(userID int, deletionType string) (*RequestInfo, error) {
	path := fmt.Sprintf("/users/%d?type=%s", userID, deletionType)

	resp, reqInfo, err := c.Delete(path, nil)
	if err != nil {
		return reqInfo, fmt.Errorf("deleting user: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return reqInfo, fmt.Errorf("%s: the provided user ID is invalid: %d", resp.Status, userID)
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return reqInfo, fmt.Errorf("%s: missing, invalid or expired API key", resp.Status)
	}

	if resp.StatusCode == http.StatusNotFound {
		return reqInfo, fmt.Errorf("%s: no user found with the specified ID: %d", resp.Status, userID)
	}

	if resp.StatusCode != http.StatusOK {
		return reqInfo, fmt.Errorf("unexpected response from API: %s", resp.Status)
	}

	return reqInfo, nil
}

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

	// rgvr users delete passes delete as the deletionType, which is the name of the command.
	// For user deletion, the API expects the deletionType to be "DELETED".
	if deletionType == "delete" {
		deletionType = "DELETED"
	}

	// rgvr users archive passes archive as the deletionType, which is the name of the command.
	// For user archiving, the API expects the deletionType to be "ARCHIVED".
	if deletionType == "archive" {
		deletionType = "ARCHIVED"
	}

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

	// As of 2026-08-21, the API returns 204 No Content for successful deletion,
	// even though the documentation states it should return 200 OK. This is a known issue.
	if resp.StatusCode != http.StatusNoContent {
		return reqInfo, fmt.Errorf("unexpected response from API: %s", resp.Status)
	}

	return reqInfo, nil
}

// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/matthieukhl/rgvr/internal/models"
)

func (c *Client) GetGroupByID(groupID string, params url.Values) (*models.Group, *RequestInfo, error) {
	path := fmt.Sprintf("/groups/%s", groupID)

	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	resp, reqInfo, err := c.Get(path)
	if err != nil {
		return nil, nil, fmt.Errorf("unexpected response from API: %s", resp.Status)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, nil, fmt.Errorf("unexpected response from API: %s", resp.Status)
	}

	var group models.Group

	if err = json.NewDecoder(resp.Body).Decode(&group); err != nil {
		return nil, nil, fmt.Errorf("decoding groups information: %w", err)
	}

	return &group, reqInfo, nil
}

func (c *Client) PatchGroupAccess(groupID string) (*RequestInfo, error) {
	path := fmt.Sprintf("/groups/%s/isjumper/toggle", groupID)

	resp, reqInfo, err := c.Patch(path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected response from API: %s", resp.Status)
	}

	return reqInfo, nil
}

func (c *Client) PatchGroupRingduration(groupID string, userID string, ringDuration int) (*RequestInfo, error) {
	path := fmt.Sprintf("/groups/%s/users/%s/ring_duration/%d", groupID, userID, ringDuration)

	resp, reqInfo, err := c.Patch(path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected response from API (%s): %s", resp.Status, string(bodyBytes))
	}

	return reqInfo, nil

}

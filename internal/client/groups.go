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

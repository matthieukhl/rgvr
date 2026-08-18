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

func (c *Client) GetTags() (*models.ListResponse[models.Tag], *RequestInfo, error) {
	path := "/tags"

	resp, reqInfo, err := c.Get(path)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, reqInfo, fmt.Errorf("%s: invalid or missing API token, or missing 'IVRs read' permission", resp.Status)
	}

	// This has been commented out because although the documentation
	// of Ringover's public API mentions a 204 in case no tags are found,
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
		return nil, reqInfo, fmt.Errorf("unexpected response from API: %w", err)
	}

	var tagsResponse models.ListResponse[models.Tag]
	if err := json.Unmarshal(bodyBytes, &tagsResponse); err != nil {
		return nil, reqInfo, fmt.Errorf("decoding tag information: %w", err)
	}

	return &tagsResponse, reqInfo, nil
}

func (c *Client) GetTag(tagID string) (*models.Tag, *RequestInfo, error) {
	path := fmt.Sprintf("/tags/%s", tagID)

	resp, reqInfo, err := c.Get(path)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return nil, reqInfo, fmt.Errorf("%s: invalid tag identifier", resp.Status)
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, reqInfo, fmt.Errorf("%s: invalid or missing API token, or missing 'IVRs Read' permission", resp.Status)
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, reqInfo, nil // We'll handle this gratefully when called in cmd/tags/get.go
	}

	if resp.StatusCode != http.StatusOK {
		return nil, reqInfo, fmt.Errorf("unexpected response from API: %w", err)
	}

	var tag *models.Tag

	if err := json.NewDecoder(resp.Body).Decode(&tag); err != nil {
		return nil, reqInfo, fmt.Errorf("decoding tag information: %w", err)
	}

	return tag, reqInfo, nil
}

func (c *Client) CreateTag(tag *models.Tag) (*RequestInfo, error) {
	path := "/tags"

	bodyBytes, err := json.Marshal(tag)
	if err != nil {
		return nil, fmt.Errorf("encoding tag information: %w", err)
	}

	resp, reqInfo, err := c.Post(path, bodyBytes)
	if err != nil {
		return reqInfo, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return reqInfo, fmt.Errorf("%s: invalid or missing API token, or missing 'IVRs Write' permission", resp.Status)
	}

	if resp.StatusCode != http.StatusOK {
		return reqInfo, fmt.Errorf("unexpected response from API: %w", err)
	}

	return reqInfo, nil
}

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

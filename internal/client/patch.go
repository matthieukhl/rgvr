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
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func (c *Client) Patch(path string, body []byte) (*http.Response, *RequestInfo, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, path)
	reqInfo := &RequestInfo{}
	req, err := http.NewRequest("PATCH", url, bytes.NewReader(body))
	if err != nil {
		return nil, nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Authorization", c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	start := time.Now()
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("making request: %w", err)
	}

	duration := time.Since(start)

	reqInfo.URL = resp.Request.URL.String()
	reqInfo.Duration = duration

	return resp, reqInfo, nil
}

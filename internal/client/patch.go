// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

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

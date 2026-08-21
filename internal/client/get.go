// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package client

import (
	"fmt"
	"net/http"
	"time"
)

// Get makes a GET request to the specified path using the client's API key for authorization.
func (c *Client) Get(path string) (*http.Response, *RequestInfo, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, path)
	requestInfo := &RequestInfo{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, requestInfo, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Authorization", c.APIKey)

	start := time.Now()
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, requestInfo, fmt.Errorf("making request: %w", err)
	}
	duration := time.Since(start)

	requestInfo.URL = resp.Request.URL.String()
	requestInfo.Duration = duration

	return resp, requestInfo, nil
}

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
	"fmt"
	"net/http"
	"time"

	"github.com/matthieukhl/rgvr/internal/api"
	"github.com/spf13/viper"
)

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
	APIKey     string
}

type RequestInfo struct {
	URL      string
	Duration time.Duration
}

// NewClient creates a new Client instance with the API key and region from the configuration.
func NewClient() (*Client, error) {
	apiKey, err := api.GetAPIKey()
	if err != nil {
		return nil, err
	}

	region := viper.GetString("region")

	var baseURL string
	switch region {
	case "eu":
		baseURL = api.EUBaseURL
	case "us":
		baseURL = api.USBaseURL
	default:
		return nil, fmt.Errorf("invalid region: %s. Valid regions are 'eu' and 'us'", region)
	}

	return &Client{
		HTTPClient: &http.Client{},
		BaseURL:    baseURL,
		APIKey:     apiKey,
	}, nil
}

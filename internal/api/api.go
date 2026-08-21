// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package api

import (
	"errors"

	"github.com/spf13/viper"
)

const (
	EUBaseURL = "https://public-api.ringover.com/v2"
	USBaseURL = "https://public-api-us.ringover.com/v2"
)

var ErrAPIKeyNotSet = errors.New("API key is not set. Please run `rgvr auth login` to set your API key.")

func GetAPIKey() (string, error) {
	apiKey := viper.GetString("api_key")

	if apiKey == "" {
		return "", ErrAPIKeyNotSet
	}

	return apiKey, nil
}

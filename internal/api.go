/*
		rgvr - A CLI to interact with Ringover's public API.
	    Copyright (C) 2026  Matthieu Khairallah

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

package internal

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

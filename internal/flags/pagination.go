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

package flags

import (
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
)

func BuildPaginationParams(cmd *cobra.Command) (url.Values, error) {
	params := url.Values{}

	limit, err := cmd.Flags().GetInt("limit")
	if err != nil {
		return nil, fmt.Errorf("retrieving 'limit' flag: %w", err)
	}

	if limit > 0 {
		params.Set("limit_count", fmt.Sprintf("%d", limit))
	}

	offset, err := cmd.Flags().GetInt("offset")
	if err != nil {
		return nil, fmt.Errorf("retrieving 'offset' flag: %w", err)
	}

	if offset > 0 {
		params.Set("limit_offset", fmt.Sprintf("%d", offset))
	}

	return params, nil
}

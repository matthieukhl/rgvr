// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

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

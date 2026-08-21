// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package convert

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func StringIDsToJSON(ids []string) ([]byte, error) {
	intIDs := make([]int, len(ids))
	for i, id := range ids {
		n, err := strconv.Atoi(id)
		if err != nil {
			return nil, fmt.Errorf("converting string ID to int: %w", err)
		}
		intIDs[i] = n
	}
	return json.Marshal(intIDs)
}

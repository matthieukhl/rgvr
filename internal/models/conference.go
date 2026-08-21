// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import "fmt"

type Conference struct {
	ID      int    `json:"conference_id"`
	Name    string `json:"name"`
	Numbers []Number
}

func (c Conference) TableHeader() []string {
	return []string{
		"Conference ID",
		"Name",
		"Numbers",
	}
}

func (c Conference) TableRow() []string {
	numbers := listNumbers(c.Numbers)

	return []string{
		fmt.Sprintf("%d", c.ID),
		c.Name,
		numbers,
	}
}

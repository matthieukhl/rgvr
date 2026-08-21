// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"fmt"
	"strings"
)

type Group struct {
	GroupID         int    `json:"group_id"`
	Name            string `json:"name"`
	TotalUsersCount int    `json:"total_users_count"`
	Color           string `json:"color"`
	IsJumper        bool   `json:"is_jumper"`
	Users           []User `json:"users,omitempty"`
}

func (g Group) TableHeader() []string {
	return []string{
		"GroupID",
		"Name",
		"TotalUsersCount",
		"Color",
		"IsJumper",
		"Users",
	}
}

func (g Group) TableRow() []string {
	users := parseUsersSlice(g.Users)

	return []string{
		fmt.Sprintf("%d", g.GroupID),
		g.Name,
		fmt.Sprintf("%d", g.TotalUsersCount),
		g.Color,
		fmt.Sprintf("%t", g.IsJumper),
		users,
	}
}

// Helper function that parses a slice of Users and returns a string.
func parseUsersSlice(users []User) string {
	var userConcatNames []string

	if len(users) == 0 {
		return ""
	}

	for _, user := range users {
		userConcatNames = append(userConcatNames, user.ConcatName)
	}

	return strings.Join(userConcatNames, "\n")
}

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

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

// User represents a user in the Ringover system.
// It implements the Tabler interface.
type User struct {
	UserID       int      `json:"user_id"`
	TeamID       int      `json:"team_id"`
	Initial      string   `json:"initial"`
	Color        string   `json:"color"`
	Firstname    string   `json:"firstname"`
	Lastname     string   `json:"lastname"`
	Company      string   `json:"company"`
	Email        string   `json:"email"`
	Picture      string   `json:"picture"`
	ConcatName   string   `json:"concat_name"`
	RingDuration int      `json:"ring_duration,omitempty"`
	Order        int      `json:"order,omitempty"`
	Numbers      []Number `json:"numbers,omitempty"`
	Plan         Plan     `json:"plan"`
}

// UserInvite represents a single user entry in an invitation request payload.
type UserInvite struct {
	Number    int    `json:"number"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	PlanID    int    `json:"plan_id"`
}

// UserInvitationPayload represents the request body for POST /users/invite.
type UserInvitationPayload struct {
	InvitedBy int          `json:"invited_by"` // ID of the user initiating the invitation
	Users     []UserInvite `json:"users"`
}

// UserInvitationResponse represents the response received after inviting a user
// to the Ringover system using POST /users/invite endpoint.
type UserInvitationResponse struct {
	ID                        int    `json:"id"`
	Email                     string `json:"email"`
	Firstname                 string `json:"firstname"`
	Lastname                  string `json:"lastname"`
	IsAdmin                   bool   `json:"is_admin"`
	IsBilling                 bool   `json:"is_billing"`
	IsSupervisorCurrentscalls bool   `json:"is_supervisor_currentscalls"`
	IsSupervisorLogs          bool   `json:"is_supervisor_logs"`
	IsSupervisorStats         bool   `json:"is_supervisor_stats"`
	IsSupervisorCampaign      bool   `json:"is_supervisor_campaign"`
	IsTechnical               bool   `json:"is_technical"`
	PlanID                    int    `json:"plan_id"`
	PlanName                  string `json:"plan_name"`
	Superadmin                bool   `json:"superadmin"`
	UserStatus                string `json:"user_status"`
	Avatar                    string `json:"avatar"`
	Admin                     string `json:"admin"`
	Waiting                   bool   `json:"waiting"`
}

func (u User) TableHeader() []string {
	return []string{
		"UserID",
		"TeamID",
		"Initial",
		"Color",
		"Firstname",
		"Lastname",
		"Company",
		"Email",
		"Picture",
		"ConcatName",
		"Numbers",
	}
}

func (u User) TableRow() []string {
	return []string{
		fmt.Sprintf("%d", u.UserID),
		fmt.Sprintf("%d", u.TeamID),
		u.Initial,
		u.Color,
		u.Firstname,
		u.Lastname,
		u.Company,
		u.Email,
		u.Picture,
		u.ConcatName,
		listNumbers(u.Numbers),
	}
}

// Helper function to list numbers for into a comma-separated string.
func listNumbers(numbers []Number) string {
	var numbersList []string

	for _, number := range numbers {
		numbersList = append(numbersList, fmt.Sprintf("%d", number.Number))
	}

	return strings.Join(numbersList, ", ")
}

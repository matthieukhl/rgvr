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

package models

import (
	"fmt"
	"time"
)

// Team represents a team in the Ringover system, including its associated numbers, users, IVRs, conferences, tags, and groups.
// It implements the Table interface.
type Team struct {
	TeamID            int      `json:"team_id"`
	Name              string   `json:"name"`
	TotalNumbersCount int      `json:"total_numbers_count"`
	Numbers           []Number `json:"numbers"`
	TotalUsersCount   int      `json:"total_users_count"`
	Users             []User   `json:"users"`
	TotalIvrsCount    int      `json:"total_ivrs_count"`
	Ivrs              []struct {
		IvrID     int      `json:"ivr_id"`
		Name      string   `json:"name"`
		Color     string   `json:"color"`
		Numbers   []Number `json:"numbers"`
		Scenarios []struct {
			ScenarioID   int    `json:"scenario_id"`
			IvrID        int    `json:"ivr_id"`
			Name         string `json:"name"`
			Color        string `json:"color"`
			ScenarioType string `json:"scenario_type"`
			IsDefault    bool   `json:"is_default"`
		} `json:"scenarios"`
		IsOpen bool `json:"is_open"`
	} `json:"ivrs"`
	TotalConferencesCount int `json:"total_conferences_count"`
	Conferences           []struct {
		ConferenceID int      `json:"conference_id"`
		Name         string   `json:"name"`
		Numbers      []Number `json:"numbers"`
	} `json:"conferences"`
	TotalTagsCount int `json:"total_tags_count"`
	Tags           []struct {
		TagID        int       `json:"tag_id"`
		Name         string    `json:"name"`
		Color        string    `json:"color"`
		Description  string    `json:"description"`
		CreationDate time.Time `json:"creation_date"`
	} `json:"tags"`
	TotalGroupsCount int `json:"total_groups_count"`
	Groups           []struct {
		GroupID         int    `json:"group_id"`
		Name            string `json:"name"`
		TotalUsersCount int    `json:"total_users_count"`
		Color           string `json:"color"`
		IsJumper        bool   `json:"is_jumper"`
	} `json:"groups"`
}

func (t Team) TableHeader() []string {
	return []string{
		"TeamID",
		"Name",
		"TotalNumbersCount",
		"TotalUsersCount",
		"TotalIvrsCount",
		"TotalConferencesCount",
		"TotalTagsCount",
		"TotalGroupsCount",
	}
}

func (t Team) TableRow() []string {
	return []string{
		fmt.Sprintf("%d", t.TeamID),
		t.Name,
		fmt.Sprintf("%d", t.TotalNumbersCount),
		fmt.Sprintf("%d", t.TotalUsersCount),
		fmt.Sprintf("%d", t.TotalIvrsCount),
		fmt.Sprintf("%d", t.TotalConferencesCount),
		fmt.Sprintf("%d", t.TotalTagsCount),
		fmt.Sprintf("%d", t.TotalGroupsCount),
	}
}

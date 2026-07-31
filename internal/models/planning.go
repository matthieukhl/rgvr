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
	"time"
)

type Planning struct {
	TeamID         int         `json:"team_id"`
	UserID         int         `json:"user_id"`
	TzIdentifier   string      `json:"tz_identifier"`
	TzNow          time.Time   `json:"tz_now"`
	PlanningEnable bool        `json:"planning_enable"`
	IsPlanning     bool        `json:"is_planning"`
	IsSnoozed      bool        `json:"is_snoozed"`
	TimeRanges     []TimeRange `json:"time_ranges"`
}

type TimeRange struct {
	Day   int `json:"day"`
	Start int `json:"start"`
	End   int `json:"end"`
}

func (p Planning) TableHeader() []string {
	return []string{
		"TeamID",
		"UserID",
		"TzIdentifier",
		"TzNow",
		"PlanningEnable",
		"IsSnoozed",
		"IsPlanning",
		"TimeRanges",
	}
}

func (p Planning) TableRow() []string {
	timeRanges := parseTimeRangeSlice(p.TimeRanges)

	return []string{
		fmt.Sprintf("%d", p.TeamID),
		fmt.Sprintf("%d", p.UserID),
		p.TzIdentifier,
		p.TzNow.String(),
		fmt.Sprintf("%t", p.PlanningEnable),
		fmt.Sprintf("%t", p.IsPlanning),
		fmt.Sprintf("%t", p.IsSnoozed),
		timeRanges,
	}
}

// Helper function that parses a slice of TimeRange into a string.
func parseTimeRangeSlice(timeRanges []TimeRange) string {
	var timeRangesList []string

	for _, timeRange := range timeRanges {
		timeRangesList = append(timeRangesList, fmt.Sprintf("Day: %d - Start: %d - End: %d\n", timeRange.Day, timeRange.Day, timeRange.End))
	}

	return strings.Join(timeRangesList, " ")
}

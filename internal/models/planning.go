// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

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

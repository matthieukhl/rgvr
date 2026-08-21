// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import "fmt"

type Presence struct {
	TeamID           int    `json:"team_id"`
	UserID           int    `json:"user_id"`
	InCall           int    `json:"in_call"`
	ConnectedDevices int    `json:"connected_devices"`
	PlanningEnable   bool   `json:"planning_enable"`
	IsPlanning       bool   `json:"is_planning"`
	IsPlannedSnooze  bool   `json:"is_planned_snooze"`
	IsSnoozed        bool   `json:"is_snoozed"`
	SnoozeType       string `json:"snooze_type"`
	SnoozeEnd        string `json:"snooze_end"`
}

func (p Presence) TableHeader() []string {
	return []string{
		"TeamID",
		"UserID",
		"InCall",
		"ConnectedDevices",
		"PlanningEnable",
		"IsPlanning",
		"IsPlannedSnooze",
		"IsSnoozed",
		"SnoozeType",
		"SnoozeEnd",
	}
}

func (p Presence) TableRow() []string {
	return []string{
		fmt.Sprintf("%d", p.TeamID),
		fmt.Sprintf("%d", p.UserID),
		fmt.Sprintf("%d", p.InCall),
		fmt.Sprintf("%d", p.ConnectedDevices),
		fmt.Sprintf("%t", p.PlanningEnable),
		fmt.Sprintf("%t", p.IsPlanning),
		fmt.Sprintf("%t", p.IsPlannedSnooze),
		fmt.Sprintf("%t", p.IsSnoozed),
		p.SnoozeType,
		p.SnoozeEnd,
	}
}

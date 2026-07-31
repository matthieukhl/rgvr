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

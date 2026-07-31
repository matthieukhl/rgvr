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

import "time"

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

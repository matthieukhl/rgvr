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
	"time"
)

type SnoozeLog struct {
	UserID              int      `json:"user_id"`
	TeamID              int      `json:"team_id"`
	LastIDOffsetSetted  int      `json:"last_id_offset_setted"`
	LimitCountSetted    int      `json:"limit_count_setted"`
	SnoozeLogListCount  int      `json:"snooze_log_list_count"`
	SnoozeLogList       []Snooze `json:"snooze_log_list"`
	TotalSnoozeLogCount int      `json:"total_snooze_log_count"`
}

func (sl SnoozeLog) TableHeader() []string {
	return []string{
		"UserID",
		"TeamID",
		"LastIDOffsetSetted",
		"LimitCountSetted",
		"SnoozeLogListCount",
		"TotalSnoozeLogCount",
	}
}

func (sl SnoozeLog) TableRow() []string {
	return []string{
		fmt.Sprintf("%d", sl.UserID),
		fmt.Sprintf("%d", sl.TeamID),
		fmt.Sprintf("%d", sl.LastIDOffsetSetted),
		fmt.Sprintf("%d", sl.LimitCountSetted),
		fmt.Sprintf("%d", sl.SnoozeLogListCount),
		fmt.Sprintf("%d", sl.TotalSnoozeLogCount),
	}
}

type Snooze struct {
	State        string    `json:"state"`
	Name         string    `json:"name"`
	Label        string    `json:"label"`
	LabelFr      string    `json:"label_fr"`
	LabelEn      string    `json:"label_en"`
	LabelEs      string    `json:"label_es"`
	Comment      string    `json:"comment"`
	CreationDate time.Time `json:"creation_date"`
	EndDate      time.Time `json:"end_date"`
}

func (s Snooze) TableHeader() []string {
	return []string{
		"State",
		"Name",
		"Label",
		"LabelFr",
		"LabelEn",
		"LabelEs",
		"Comment",
		"CreationDate",
		"EndDate",
	}
}

func (s Snooze) TableRow() []string {
	return []string{
		s.State,
		s.Name,
		s.Label,
		s.LabelFr,
		s.LabelEn,
		s.LabelEs,
		s.Comment,
		s.CreationDate.String(),
		s.EndDate.String(),
	}
}

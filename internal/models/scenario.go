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

type Scenario struct {
	Id        int    `json:"scenario_id"`
	IvrId     int    `json:"ivr_id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	Type      string `json:"scenario_type"`
	IsDefault bool   `json:"is_default"`
}

func (s Scenario) TableHeader() []string {
	return []string{
		"Scenario ID",
		"IVR ID",
		"Name",
		"Color",
		"Scenario Type",
		"Is Default",
	}
}

func (s Scenario) TableRow() []string {
	return []string{
		fmt.Sprintf("%d", s.Id),
		fmt.Sprintf("%d", s.IvrId),
		s.Name,
		s.Color,
		s.Type,
		fmt.Sprintf("%t", s.IsDefault),
	}
}

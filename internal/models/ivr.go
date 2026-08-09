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

type IVR struct {
	IvrID     int        `json:"ivr_id"`
	Name      string     `json:"name"`
	Color     string     `json:"color"`
	Numbers   []Number   `json:"numbers,omitempty"` // omitempty tag necessary because of the difference of the IVR objects sent back by GET /ivrs and GET /ivrs/{ivr_id}
	Scenarios []Scenario `json:"scenarios"`
	IsOpen    bool       `json:"is_open"`
}

func (i IVR) TableHeader() []string {
	return []string{
		"Ivr ID",
		"Name",
		"Color",
		"Numbers",
		"Scenarios",
		"Is Open",
	}
}

func (i IVR) TableRow() []string {
	return []string{
		fmt.Sprintf("%d", i.IvrID),
		i.Name,
		i.Color,
		listNumbers(i.Numbers),
		listScernarios(i.Scenarios),
		fmt.Sprintf("%t", i.IsOpen),
	}
}

// Helper function that parses a slice of
func listScernarios(scenarios []Scenario) string {
	var scenariosList []string

	for _, scenario := range scenarios {
		scenariosList = append(scenariosList, scenario.Name)
	}

	return strings.Join(scenariosList, ", ")
}

// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

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

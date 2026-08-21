// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

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

type SalesforceIntegration struct {
	Object   string `json:"object"`
	ObjectID string `json:"object_id"`
}

type Integrations struct {
	Salesforce *SalesforceIntegration `json:"salesforce,omitempty"`
}

type IVRCallback struct {
	FromNumber   int64         `json:"from_number"`
	Clir         bool          `json:"clir"`
	ToNumber     int64         `json:"to_number"`
	Timeout      int64         `json:"timeout"`
	Integrations *Integrations `json:"integrations,omitempty"`
}

type IVRCallbackResponse struct {
	CallID    uint64 `json:"call_id"`
	ChannelID uint64 `json:"channel_id"`
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

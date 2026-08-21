// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"fmt"
)

type Number struct {
	Number       int64  `json:"number"`
	Label        string `json:"label"`
	Type         string `json:"type"`
	UserID       int    `json:"user_id"`
	IvrID        int    `json:"ivr_id"`
	ConferenceID int    `json:"conference_id"`
	IsSms        bool   `json:"is_sms"`
	IsSmsWrite   bool   `json:"is_sms_write"`
	IsCallable   bool   `json:"is_callable"`
	NumberFormat `json:"format"`
}

type NumberAssignment struct {
	UserID       int64 `json:"user_id,omitempty"`
	IvrID        int64 `json:"ivr_id,omitempty"`
	ConferenceID int64 `json:"conference_id,omitempty"`
}

func (n Number) TableHeader() []string {
	return []string{
		"Number",
		"Label",
		"Type",
		"User ID",
		"IVR ID",
		"Conference ID",
		"Is SMS",
		"Is SMS Write",
		"Is Callable",
	}
}

func (n Number) TableRow() []string {
	return []string{
		fmt.Sprintf("%d", n.Number),
		n.Label,
		n.Type,
		fmt.Sprintf("%d", n.UserID),
		fmt.Sprintf("%d", n.IvrID),
		fmt.Sprintf("%d", n.ConferenceID),
		fmt.Sprintf("%t", n.IsSms),
		fmt.Sprintf("%t", n.IsSmsWrite),
		fmt.Sprintf("%t", n.IsCallable),
	}
}

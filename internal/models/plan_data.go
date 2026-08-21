// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import "fmt"

type PlanData struct {
	PlanID         int    `json:"plan_id"`
	PlanName       string `json:"plan_name"`
	NbLicencesUsed int    `json:"nb_licences_used"`
	NbLicences     int    `json:"nb_licences"`
}

func (pd PlanData) TableHeader() []string {
	return []string{
		"Plan ID",
		"Plan Name",
		"Number of Licences Used",
		"Total Number of Licences",
	}
}

func (pd PlanData) TableRow() []string {
	return []string{
		fmt.Sprintf("%d", pd.PlanID),
		pd.PlanName,
		fmt.Sprintf("%d", pd.NbLicencesUsed),
		fmt.Sprintf("%d", pd.NbLicences),
	}
}

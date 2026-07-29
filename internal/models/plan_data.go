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

package models

type PlanDataRaw []struct {
	PlanID         int    `json:"plan_id"`
	PlanName       string `json:"plan_name"`
	NbLicencesUsed int    `json:"nb_licences_used"`
	NbLicences     int    `json:"nb_licences"`
}

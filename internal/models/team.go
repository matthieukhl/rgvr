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

type Team struct {
	TeamID            int    `json:"team_id"`
	Name              string `json:"name"`
	TotalNumbersCount int    `json:"total_numbers_count"`
	Numbers           []struct {
		Number       int64  `json:"number"`
		Label        string `json:"label"`
		Type         string `json:"type"`
		UserID       int    `json:"user_id"`
		IvrID        int    `json:"ivr_id"`
		ConferenceID int    `json:"conference_id"`
		IsSms        bool   `json:"is_sms"`
		IsSmsWrite   bool   `json:"is_sms_write"`
		IsCallable   bool   `json:"is_callable"`
		Format       struct {
			Raw              int    `json:"raw"`
			CountryCode      string `json:"country_code"`
			Country          string `json:"country"`
			E164             string `json:"e164"`
			International    string `json:"international"`
			InternationalAlt string `json:"international_alt"`
			National         string `json:"national"`
			NationalAlt      string `json:"national_alt"`
			Rfc3966          string `json:"rfc3966"`
			IsShortCode      bool   `json:"is_short_code"`
		} `json:"format"`
	} `json:"numbers"`
	TotalUsersCount int `json:"total_users_count"`
	Users           []struct {
		UserID     int    `json:"user_id"`
		TeamID     int    `json:"team_id"`
		Initial    string `json:"initial"`
		Color      string `json:"color"`
		Firstname  string `json:"firstname"`
		Lastname   string `json:"lastname"`
		Company    string `json:"company"`
		Email      string `json:"email"`
		Picture    string `json:"picture"`
		ConcatName string `json:"concat_name"`
		Numbers    []struct {
			Number       int64  `json:"number"`
			Label        string `json:"label"`
			Type         string `json:"type"`
			UserID       int    `json:"user_id"`
			IvrID        int    `json:"ivr_id"`
			ConferenceID int    `json:"conference_id"`
			IsSms        bool   `json:"is_sms"`
			IsSmsWrite   bool   `json:"is_sms_write"`
			IsCallable   bool   `json:"is_callable"`
			Format       struct {
				Raw              int    `json:"raw"`
				CountryCode      string `json:"country_code"`
				Country          string `json:"country"`
				E164             string `json:"e164"`
				International    string `json:"international"`
				InternationalAlt string `json:"international_alt"`
				National         string `json:"national"`
				NationalAlt      string `json:"national_alt"`
				Rfc3966          string `json:"rfc3966"`
				IsShortCode      bool   `json:"is_short_code"`
			} `json:"format"`
		} `json:"numbers"`
	} `json:"users"`
	TotalIvrsCount int `json:"total_ivrs_count"`
	Ivrs           []struct {
		IvrID   int    `json:"ivr_id"`
		Name    string `json:"name"`
		Color   string `json:"color"`
		Numbers []struct {
			Number       int64  `json:"number"`
			Label        string `json:"label"`
			Type         string `json:"type"`
			UserID       int    `json:"user_id"`
			IvrID        int    `json:"ivr_id"`
			ConferenceID int    `json:"conference_id"`
			IsSms        bool   `json:"is_sms"`
			IsSmsWrite   bool   `json:"is_sms_write"`
			IsCallable   bool   `json:"is_callable"`
			Format       struct {
				Raw              int    `json:"raw"`
				CountryCode      string `json:"country_code"`
				Country          string `json:"country"`
				E164             string `json:"e164"`
				International    string `json:"international"`
				InternationalAlt string `json:"international_alt"`
				National         string `json:"national"`
				NationalAlt      string `json:"national_alt"`
				Rfc3966          string `json:"rfc3966"`
				IsShortCode      bool   `json:"is_short_code"`
			} `json:"format"`
		} `json:"numbers"`
		Scenarios []struct {
			ScenarioID   int    `json:"scenario_id"`
			IvrID        int    `json:"ivr_id"`
			Name         string `json:"name"`
			Color        string `json:"color"`
			ScenarioType string `json:"scenario_type"`
			IsDefault    bool   `json:"is_default"`
		} `json:"scenarios"`
		IsOpen bool `json:"is_open"`
	} `json:"ivrs"`
	TotalConferencesCount int `json:"total_conferences_count"`
	Conferences           []struct {
		ConferenceID int    `json:"conference_id"`
		Name         string `json:"name"`
		Numbers      []struct {
			Number       int64  `json:"number"`
			Label        string `json:"label"`
			Type         string `json:"type"`
			UserID       int    `json:"user_id"`
			IvrID        int    `json:"ivr_id"`
			ConferenceID int    `json:"conference_id"`
			IsSms        bool   `json:"is_sms"`
			IsSmsWrite   bool   `json:"is_sms_write"`
			IsCallable   bool   `json:"is_callable"`
			Format       struct {
				Raw              int    `json:"raw"`
				CountryCode      string `json:"country_code"`
				Country          string `json:"country"`
				E164             string `json:"e164"`
				International    string `json:"international"`
				InternationalAlt string `json:"international_alt"`
				National         string `json:"national"`
				NationalAlt      string `json:"national_alt"`
				Rfc3966          string `json:"rfc3966"`
				IsShortCode      bool   `json:"is_short_code"`
			} `json:"format"`
		} `json:"numbers"`
	} `json:"conferences"`
	TotalTagsCount int `json:"total_tags_count"`
	Tags           []struct {
		TagID        int       `json:"tag_id"`
		Name         string    `json:"name"`
		Color        string    `json:"color"`
		Description  string    `json:"description"`
		CreationDate time.Time `json:"creation_date"`
	} `json:"tags"`
	TotalGroupsCount int `json:"total_groups_count"`
	Groups           []struct {
		GroupID         int    `json:"group_id"`
		Name            string `json:"name"`
		TotalUsersCount int    `json:"total_users_count"`
		Color           string `json:"color"`
		IsJumper        bool   `json:"is_jumper"`
	} `json:"groups"`
}

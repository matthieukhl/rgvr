// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

type NumberFormat struct {
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
}

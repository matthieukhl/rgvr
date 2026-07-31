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

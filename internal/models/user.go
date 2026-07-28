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

// User represents a user in the Ringover system.
type User struct {
	UserID     int      `json:"user_id"`
	TeamID     int      `json:"team_id"`
	Initial    string   `json:"initial"`
	Color      string   `json:"color"`
	Firstname  string   `json:"firstname"`
	Lastname   string   `json:"lastname"`
	Company    string   `json:"company"`
	Email      string   `json:"email"`
	Picture    string   `json:"picture"`
	ConcatName string   `json:"concat_name"`
	Numbers    []Number `json:"numbers"`
}

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

import (
	"fmt"
	"strings"
	"time"
)

// TagColor represents a valid Ringover tag color, identified by its hex code.
// These are the canonical hex codes from Ringover's current design system.
type TagColor string

const (
	ColorYellow TagColor = "FFD54F"
	ColorOrange TagColor = "FFB74D"
	ColorBrown  TagColor = "A1887F"
	ColorRed    TagColor = "FF6B6B"
	ColorPink   TagColor = "F06292"
	ColorPurple TagColor = "BA68C8"
	ColorBlue   TagColor = "64B5F6"
	ColorGreen  TagColor = "81C784"
	ColorGrey   TagColor = "E0E0E0"
)

func (tc TagColor) String() string {
	return string(tc)
}

var colorNames = map[string]TagColor{
	"yellow": ColorYellow,
	"orange": ColorOrange,
	"brown":  ColorBrown,
	"red":    ColorRed,
	"pink":   ColorPink,
	"purple": ColorPurple,
	"blue":   ColorBlue,
	"green":  ColorGreen,
	"grey":   ColorGrey,
}

func ParseTagColor(colorName string) (TagColor, error) {
	normalized := strings.ToLower(colorName)

	if color, exists := colorNames[normalized]; exists {
		return color, nil
	}

	return "", fmt.Errorf("invalid tag color: %s", colorName)
}

type Tag struct {
	TagID        int       `json:"tag_id"`
	Name         string    `json:"name"`
	Color        TagColor  `json:"color"`
	Description  string    `json:"description"`
	CreationDate time.Time `json:"creation_date"`
}

func (t Tag) TableHeader() []string {
	return []string{
		"Tag ID",
		"Name",
		"Color",
		"Description",
		"Creation Date",
	}
}

func (t Tag) TableRow() []string {
	return []string{
		fmt.Sprintf("%d", t.TagID),
		t.Name,
		t.Color.String(),
		t.Description,
		t.CreationDate.String(),
	}
}

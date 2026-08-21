// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

type Plan struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

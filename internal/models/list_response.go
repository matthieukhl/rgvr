// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

// Ringover's API returns a list of items in a paginated format.
// The ListResponse struct is a generic representation of such a response,
// where T can be any type that represents the items in the list.
type ListResponse[T any] struct {
	ListCount int `json:"list_count"`
	List      []T `json:"list"`
}

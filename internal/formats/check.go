// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package formats

import "errors"

var ErrIncorrectFormat = errors.New("Incorrect format. Available formats: 'table', 'json'.")

func Check(f string) error {
	if f != "json" && f != "table" {
		return ErrIncorrectFormat
	}

	return nil
}

// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/matthieukhl/rgvr/cmd"
	_ "github.com/matthieukhl/rgvr/cmd/auth"
	_ "github.com/matthieukhl/rgvr/cmd/conferences"
	_ "github.com/matthieukhl/rgvr/cmd/groups"
	_ "github.com/matthieukhl/rgvr/cmd/ivrs"
	_ "github.com/matthieukhl/rgvr/cmd/numbers"
	_ "github.com/matthieukhl/rgvr/cmd/scenarios"
	_ "github.com/matthieukhl/rgvr/cmd/tags"
	_ "github.com/matthieukhl/rgvr/cmd/teams"
	_ "github.com/matthieukhl/rgvr/cmd/users"
)

func main() {
	cmd.Execute()
}

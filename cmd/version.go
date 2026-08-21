// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import "runtime/debug"

var (
	appVersion = "dev" // overridden by ldflags during `make build`; falls back to build info for `go install`
)

func init() {
	if appVersion == "dev" {
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "" && info.Main.Version != "(devel)" {
			appVersion = info.Main.Version
		}
	}
}

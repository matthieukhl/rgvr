// Copyright (c) 2026 Matthieu Khairallah. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package formats

import (
	"io"

	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/tw"
)

type Tabler interface {
	TableHeader() []string
	TableRow() []string
}

// Table function renders API output in table format
// leveraging Tabler interface.
func Table[T Tabler](w io.Writer, items []T) error {
	if len(items) == 0 {
		return nil
	}

	table := tablewriter.NewTable(w, tablewriter.WithHeaderAutoFormat(tw.Off))
	table.Header(items[0].TableHeader())

	for _, item := range items {
		table.Append(item.TableRow())
	}

	table.Render()

	return nil
}

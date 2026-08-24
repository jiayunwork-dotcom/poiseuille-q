package report

import (
	"io"
	"strings"
)

// Row is a single formatted report line.
type Row struct {
	Key   string
	Value string
}

// Table collects key-value rows for ordered output.
type Table struct {
	Rows []Row
}

// Add appends a row.
func (t *Table) Add(key, value string) {
	t.Rows = append(t.Rows, Row{Key: key, Value: value})
}

// AddFloat appends a numeric row.
func (t *Table) AddFloat(key string, v float64) {
	t.Add(key, FormatNumber(v))
}

// Render writes rows as aligned lines.
func (t *Table) Render(w io.Writer) error {
	width := 0
	for _, r := range t.Rows {
		if len(r.Key) > width {
			width = len(r.Key)
		}
	}
	for _, r := range t.Rows {
		line := PadRight(r.Key, width) + "  " + r.Value + "\n"
		if _, err := io.WriteString(w, line); err != nil {
			return err
		}
	}
	return nil
}

// CSV writes the table as comma-separated values.
func (t *Table) CSV(w io.Writer) error {
	parts := make([]string, 0, len(t.Rows))
	for _, r := range t.Rows {
		parts = append(parts, r.Key+","+r.Value)
	}
	_, err := io.WriteString(w, strings.Join(parts, "\n")+"\n")
	return err
}

// Len returns the number of rows.
func (t *Table) Len() int {
	return len(t.Rows)
}

package reporter

import (
	"errors"
	"strconv"
)

// Reporter FromTable - make reporter from table.
// If 'hpos' < 0, then has no header.
// Else header position always a first row
func FromTable(rows [][]interface{}, hpos int) *ReportTable {
	return &ReportTable{
		report: &report{
			maxRows: len(rows),
			hpos:    hpos,
		},
		rows: rows,
	}
}

func (r *ReportTable) Next() bool {
	return next(r.report)
}

func (r *ReportTable) Columns() ([]string, error) {
	if len(r.rows) == 0 {
		return []string{}, errors.New("empty rows")
	}

	if r.hpos < 0 {
		return []string{}, nil
	}

	defer func() {
		r.currentRow += 1
	}()

	//first line is always the header
	header := []string{}
	row := r.rows[0]
	var name string
	for _, cell := range row {
		switch cell.(type) {
		case string:
			name = cell.(string)
		case int64:
			name = strconv.FormatInt(cell.(int64), 10)
		case float64:
			name = strconv.FormatFloat(cell.(float64), 'f', -1, 32)
		case bool:
			name = strconv.FormatBool(cell.(bool))
		default:
			name = "column?"
		}
		header = append(header, name)
	}
	return header, nil
}

func (r *ReportTable) SliceScan() ([]interface{}, error) {
	defer func() {
		r.currentRow += 1
	}()
	return r.rows[r.currentRow], nil
}

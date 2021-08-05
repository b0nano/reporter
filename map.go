package reporter

// Reporter FromMap - make reporter from map
func (r *Reporter) FromMap(rows []map[string]interface{}) {
	r.store = &ReportFromMap{
		report: &report{
			maxRows: len(rows),
		},
		rows: rows,
	}
}

func (r *ReportFromMap) Next() bool {
	return next(r.report)
}

func (r *ReportFromMap) Columns() ([]string, error) {
	columns := []string{}
	colNames := map[string]struct{}{}
	for _, row := range r.rows {
		for col := range row {
			colNames[col] = struct{}{}
		}
	}

	for col := range colNames {
		columns = append(columns, col)
	}

	return columns, nil
}

func (r *ReportFromMap) SliceScan() ([]interface{}, error) {
	// TODO::
	return []interface{}{}, nil
}

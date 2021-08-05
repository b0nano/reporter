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

	r.cols = columns

	return columns, nil
}

func (r *ReportFromMap) SliceScan() ([]interface{}, error) {

	res := []interface{}{}

	row := r.rows[r.currentRow]
	for _, col := range r.cols {
		if val, ok := row[col]; ok {
			res = append(res, val)
		} else {
			res = append(res, "")
		}
	}

	defer func() {
		r.currentRow += 1
	}()

	return res, nil
}

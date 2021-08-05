package reporter

// Reporter From Rows - make reporter from query rows
func (r *Reporter) FromRows(rows Rowsy) {
	r.store = &ReportFromRows{
		rows: rows,
	}
}

func (r *ReportFromRows) Next() bool {
	return r.rows.Next()
}

func (r *ReportFromRows) Columns() ([]string, error) {
	return r.rows.Columns()
}

func (r *ReportFromRows) SliceScan() ([]interface{}, error) {
	return r.rows.SliceScan()
}

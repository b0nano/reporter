package reporter

// Reporter From Rows - make reporter from query rows
func FromRows(rows Rowsy) *ReportRows {
	return &ReportRows{
		rows: rows,
	}
}

func (r *ReportRows) Next() bool {
	return r.rows.Next()
}

func (r *ReportRows) Columns() ([]string, error) {
	return r.rows.Columns()
}

func (r *ReportRows) SliceScan() ([]interface{}, error) {
	return r.rows.SliceScan()
}

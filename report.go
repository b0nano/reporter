package reporter

type report struct {
	currentRow int
	maxRows    int
	hpos       int
}

type ReportFromRows struct {
	rows Rowsy
}

type ReportFromTable struct {
	*report
	rows [][]interface{}
}

type ReportFromMap struct {
	*report
	rows []map[string]interface{}
}

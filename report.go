package reporter

type report struct {
	currentRow int
	maxRows    int
	hpos       int
}

type ReportRows struct {
	rows Rowsy
}

type ReportTable struct {
	*report
	rows [][]interface{}
}

type ReportMap struct {
	*report
	cols []string
	rows []map[string]interface{}
}

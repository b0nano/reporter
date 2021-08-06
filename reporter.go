package reporter

import (
	"errors"
	"io"
	"time"

	"github.com/tealeg/xlsx"
)

type Rowsy interface {
	Next() bool
	Columns() ([]string, error)
	SliceScan() ([]interface{}, error)
}

type Reporter struct {
	store         Rowsy
	totalStore    Rowsy
	name          string
	Title         string
	totalPosition int
	file          *xlsx.File
	sh            *xlsx.Sheet
}

func ReporterNew(name string) *Reporter {
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("sheet1")
	return &Reporter{
		name: name,
		file: file,
		sh:   sheet,
	}
}

func (r *Reporter) SetTitle(title string) {
	r.Title = title
}

func (r *Reporter) SetTotalPosition(position int) {
	r.totalPosition = position
}

func (r *Reporter) SetStores(rows, totalRows Rowsy) {
	r.store = rows
	r.totalStore = totalRows
}

func (r *Reporter) Exec(w io.Writer) error {

	r.renderTitle(r.sh)
	r.renderHeader(r.sh)
	r.renderRows(r.sh)
	r.renderTotal(r.sh)

	return r.file.Write(w)
}

func (r *Reporter) renderTitle(sh *xlsx.Sheet) {
	if r.Title == "" {
		return
	}

	sh.AddRow()
	row := sh.AddRow()
	AddTitle(row, r.Title)
	sh.AddRow()
}

func (r *Reporter) renderHeader(sh *xlsx.Sheet) {
	col, err := r.store.Columns()
	if err != nil {
		return
	}

	if len(col) == 0 {
		return
	}

	sh.AddRow()
	row := sh.AddRow()
	row.AddCell()
	curIdx, err := GetRowIndex(sh, row)
	if err != nil {
		return
	}
	numCols := row.WriteSlice(&col, -1)

	//set colls style
	if numCols > 0 {
		for i := 1; i <= numCols; i++ {
			cell := sh.Cell(curIdx, i)
			cell.SetStyle(HeaderCellStyle)
		}
	}
}

func RenderRows(store Rowsy, sh *xlsx.Sheet, style *xlsx.Style) {
	if store == nil {
		return
	}

	for store.Next() {
		container, err := store.SliceScan()
		if err != nil {
			return
		}
		row := sh.AddRow()
		row.AddCell()
		for _, v := range container {
			cell := row.AddCell()
			switch v := v.(type) {
			case string:
				cell.SetString(v)
			case []byte:
				cell.SetString(string(v))
			case int64:
				cell.SetInt64(v)
			case float64:
				cell.SetFloat(v)
			case bool:
				cell.SetBool(v)
			case time.Time:
				cell.SetDateTime(v)
			default:
				cell.SetValue(v)
			}
			cell.SetStyle(style)
		}

	}
}

func (r *Reporter) renderTotal(sh *xlsx.Sheet) {
	RenderRows(r.totalStore, sh, TotalCellStyle)
}

func (r *Reporter) renderRows(sh *xlsx.Sheet) {
	RenderRows(r.store, sh, CellStyle)
}

func next(r *report) bool {
	if r.maxRows == 0 {
		return false
	}
	canNext := r.currentRow < r.maxRows
	return canNext
}

func AddTitle(row *xlsx.Row, title string) {
	row.AddCell()
	cell := row.AddCell()
	cell.SetString(title)
	cell.SetStyle(TitleStyle)
}

func GetRowIndex(sheet *xlsx.Sheet, row *xlsx.Row) (int, error) {
	for idx, r := range sheet.Rows {
		if r == row {
			return idx, nil
		}
	}

	return -1, errors.New("can not find row")
}

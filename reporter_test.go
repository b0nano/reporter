package reporter

import (
	"bytes"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestExec(t *testing.T) {

	buf := bytes.NewBuffer([]byte{})
	payloadTable := [][]interface{}{
		{"num", "first", "second", "third", "name", "surname", "phone"},
		{1, 1, 150, "test string 1", "John", "Brown", "78889999999"},
		{2, "second string", 200.00, "test string 2", "Samanta", "Fox", "71234567890"},
		{3, "third string", -59.10, 300, "Bob", "Uncle", "49500000000"},
	}

	rep := ReporterNew("")
	rep.SetTitle("Test from table")
	rep.FromTable(payloadTable, 0)

	err := rep.Exec(buf)
	if err != nil {
		fmt.Println(err)
	}

	if file, err := os.Create("test1.xlsx"); err != nil {
		fmt.Println(err)
	} else {
		file.Write(buf.Bytes())
		file.Close()
	}

	payloadMap := []map[string]interface{}{
		{
			"num":        1,
			"data":       "Some data",
			"Дата время": time.Now(),
		},
		{
			"num":        2,
			"data":       "Something else data",
			"Дата время": time.Now(),
		},
		{
			"num":        3,
			"data":       "Third data",
			"Дата время": time.Now(),
		},
		{
			"new column": "There's something here",
		},
	}

	reporter := ReporterNew("Test from map")
	reporter.SetTitle("test from map")
	reporter.FromMap(payloadMap)
	reporter.Exec(buf)

	if file, err := os.Create("test2.xlsx"); err != nil {
		fmt.Println(err)
	} else {
		file.Write(buf.Bytes())
		file.Close()
	}
}

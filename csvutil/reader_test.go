package csvutil

import (
	"bytes"
	"testing"

	"encoding/csv"
	"io/ioutil"
	"reflect"
)

func TestNewTolerantReader(t *testing.T) {

	// func main() {

	fileNewline, err := ioutil.ReadFile("test_newline.csv")
	if err != nil {
		t.Errorf("Failed to read test_newline.csv: %s ", err.Error())
	}
	fileCarriage, err := ioutil.ReadFile("test_carriage.csv")
	if err != nil {
		t.Errorf("Failed to read test_carriage.csv: %s ", err.Error())
	}

	csvrNewline := csv.NewReader(bytes.NewBuffer(fileNewline))
	csvrCarriage := csv.NewReader(bytes.NewBuffer(fileCarriage))
	csvrNewlineTol := NewTolerantReader(bytes.NewBuffer(fileNewline))
	csvrCarriageTol := NewTolerantReader(bytes.NewBuffer(fileCarriage))

	recNewline, err := csvrNewline.ReadAll()
	if err != nil {
		t.Errorf("Failed to read csv.NewReader datastream read from test_newline.csv: %s ", err.Error())
	}
	recCarriage, err := csvrCarriage.ReadAll()
	if err != nil {
		t.Errorf("Failed to read csv.NewReader datastream read from test_carriage.csv: %s ", err.Error())
	}
	recNewlineTol, err := csvrNewlineTol.ReadAll()
	if err != nil {
		t.Errorf("Failed to read NewTolerantReader datastream read from test_newline.csv: %s ", err.Error())
	}
	recCarriageTol, err := csvrCarriageTol.ReadAll()
	if err != nil {
		t.Errorf("Failed to read NewTolerantReader datastream read from test_carriage.csv: %s ", err.Error())
	}

	// TESTS:
	// Should be no difference for newline files.
	if !reflect.DeepEqual(recNewline, recNewlineTol) {
		t.Errorf("csv.NewReader and NewTolerantReader are reading test_newline.csv differently. These shold be read the same.")
		// fmt.Println(recNewline)
		// fmt.Println(recNewlineTol)
		// fmt.Println("")
	}
	// Carriage return files should be read differently.
	if reflect.DeepEqual(recCarriage, recCarriageTol) {
		t.Errorf("csv.NewReader and NewTolerantReader are reading test_carriage.csv similarly. These should be different due to carriage returns.")
		// fmt.Println(recCarriage)
		// fmt.Println(recCarriageTol)
		// fmt.Println("")
	}
	// Check that carriage and newline files are read the same by NewToleranceReader:
	if !reflect.DeepEqual(recNewlineTol, recCarriageTol) {
		t.Errorf("NewTolerantReader is reading test_newline.csv different from test_carriage.csv. These should be the same.")
		// fmt.Println(recNewlineTol)
		// fmt.Println(recCarriageTol)
		// fmt.Println("")
	}
	if !reflect.DeepEqual(recNewline, recCarriageTol) {
		t.Errorf("csv.NewReader is reading test_newline.csv differently than NewTolerantReader reads test_carriage.csv. These should be the same.")
		// fmt.Println(recNewline)
		// fmt.Println(recCarriageTol)
		// fmt.Println("")
	}
}

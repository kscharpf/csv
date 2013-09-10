package csv

import (
  "fmt"
  "testing"
)

func check_index(m CsvMatrix, expected int, name string, t *testing.T) {
  if m.ColumnIndex(name) != expected {
    t.Errorf("Incorrect index %d for %s expected %d", m.ColumnIndex(name), name, expected)
  }
}

func TestInit(t *testing.T) {
  cols := make([]string, 5)
  cols[0] = "Col0"
  cols[1] = "Col1"
  cols[2] = "Col2"
  cols[3] = "Col3"
  cols[4] = "Col4"
  m := NewCsvMatrix(cols)
  check_index(m, 0, "Col0", t)
  check_index(m, 1, "Col1", t)
  check_index(m, 2, "Col2", t)
  check_index(m, 3, "Col3", t)
  check_index(m, 4, "Col4", t)
}

func TestAppendFile(t *testing.T) {
  cols := make([]string, 2)
  cols[0] = "Col0"
  cols[1] = "Col1"
  m := NewCsvMatrix(cols)

  sNew := fmt.Sprintf("Col0,Col1\nv1,v2\nv3,v4")
 
  m.AppendCsv(sNew)  
  if sNew != m.DumpCsv() {
    t.Errorf("Appended CSV Expected %v Got %v\n", sNew, m.DumpCsv())
  }
}

func TestAppendFile3Cols(t *testing.T) {
  cols := make([]string, 3)
  cols[0] = "Col0"
  cols[1] = "Col1"
  cols[2] = "Col2"
  m := NewCsvMatrix(cols)

  sNew := fmt.Sprintf("Col0,Col1\nv1,v2\nv3,v4")

  sOut := fmt.Sprintf("Col0,Col1,Col2\nv1,v2,\nv3,v4,")
 
  m.AppendCsv(sNew)  
  if sOut != m.DumpCsv() {
    t.Errorf("Appended CSV Expected %v Got %v\n", sOut, m.DumpCsv())
  }

  sNew = fmt.Sprintf("Col0,Col2\nv1,v5\nv3,v6")
  m.AppendCsv(sNew)  
  sOut = fmt.Sprintf("Col0,Col1,Col2\nv1,v2,v5\nv3,v4,v6")
  if sOut != m.DumpCsv() {
    t.Errorf("Appended CSV Expected %v Got %v\n", sOut, m.DumpCsv())
  }
}

package csv

import (
  "strings"
  "github.com/kscharpf/stringmatrix"
  "fmt"
)

type CsvMatrix struct {
  M stringmatrix.StringMatrix
  ColumnNames []string
  ColName2Index map[string]int
}

func NewCsvMatrix(headers [] string) CsvMatrix {
  c := CsvMatrix{}
  c.ColumnNames = headers
  c.ColName2Index = make(map[string]int)
  for i := range c.ColumnNames {
    c.ColName2Index[c.ColumnNames[i]] = i
  }
  return c
}


func (m CsvMatrix) ColumnIndex(colName string) int{
  return m.ColName2Index[colName]
}


func (m CsvMatrix) Column(idx int) []string{
  return m.Column(idx)
}

func (m CsvMatrix) Row(idx int) []string{
  return m.Row(idx)
}

func NewCsvMatrixFromString(s string) CsvMatrix{
  c := CsvMatrix{}
  lines := strings.Split(s,"\n")  
  c.ColumnNames = strings.Split(lines[0], ",")
  c.ColName2Index = make(map[string]int)

  for i:= range c.ColumnNames {
    c.ColName2Index[c.ColumnNames[i]] = i
  }

  for i:= range lines[1:] {
    fields := strings.Split(lines[i], ",")
    c.M.AppendRow(fields)
  }
  return c
}
 

func (m *CsvMatrix) AppendCsv(s string) {
  c := NewCsvMatrixFromString(s)
  fmt.Printf("%s\n", c.DumpCsv()) 
  //for i := range c.ColumnNames {
    //colIdx := m.ColumnIndex(c.ColumnNames[i])
//
    //m.M.AppendArrayToColumn(c.Column(i), colIdx)
  //}
}

func (m *CsvMatrix) DumpCsv() string {
  rows := make([]string, m.M.NumRows())

  for i := 0; i<m.M.NumRows(); i++ {
    rows[i] = strings.Join(m.Row(i), ",")
  }
  return strings.Join(rows, "\n")
}

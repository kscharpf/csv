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
  RowKeys map[string]bool
}

func NewCsvMatrix(headers [] string) CsvMatrix {
  c := CsvMatrix{}

  c.ColumnNames = make([]string, len(headers))
  for i:= range headers {
    c.ColumnNames[i] = strings.TrimSpace(headers[i])
  }
  c.ColName2Index = make(map[string]int)
  for i := range c.ColumnNames {
    c.ColName2Index[c.ColumnNames[i]] = i
  }

  c.RowKeys = make(map[string]bool)

  c.M = stringmatrix.NewStringMatrixWithSize(0, len(c.ColumnNames))

  return c
}

func (m CsvMatrix) ColumnIndex(colName string) int {
  return m.ColName2Index[colName]
}


func (m CsvMatrix) Column(idx int) []string{
  return m.M.Column(idx)
}

func (m CsvMatrix) Row(idx int) []string{
  return m.M.Row(idx)
}

func NewCsvMatrixFromString(s string) CsvMatrix{
  c := CsvMatrix{}
  lines := strings.Split(s,"\n")  

  t := strings.Split(lines[0], ",")
  c.ColumnNames = make([]string, len(t))
  for i:= range t {
    c.ColumnNames[i] = strings.TrimSpace(t[i])
  }
  c.ColName2Index = make(map[string]int)

  for i:= range c.ColumnNames {
    c.ColName2Index[c.ColumnNames[i]] = i
  }

  for i:= range lines[1:] {
    tfields := strings.Split(lines[i+1], ",")
    fields := make([]string, len(tfields))
    if len(fields) > 0 {
      for j := range tfields {
        fields[j] = strings.TrimSpace(tfields[j])
      } 

      if len(fields) > 0 && fields[0] != "" {
        c.M.AppendRow(fields)
      }
    }
  }
  c.RowKeys = make(map[string]bool)

  return c
}

func (m *CsvMatrix) AppendCsv(s string) {

  c := NewCsvMatrixFromString(s)

  // Search for key field

  r1 := c.ColumnNames[0]

  m1 := m.ColumnNames[0]

  if r1 != m1 {
    panic(fmt.Sprintf("Unmatched key fields: got %v expected %v",
                         r1, m1))
  }

  newKeys := c.Column(0)

  startingLine := m.M.NumRows()

  for key := range newKeys {
    if ! m.RowKeys[newKeys[key]] {
      startingLine = key
      break
    }
  }  


  for idx := startingLine; idx < len(newKeys); idx++ {
    m.RowKeys[newKeys[idx]] = true
  }


  m.M.AppendArrayToColumn(c.Column(0)[startingLine:], 0)
  m.M.ReplaceArrayInColumn(c.Column(1), m.ColName2Index[strings.ToLower(c.ColumnNames[1])])
}

func (m *CsvMatrix) DumpCsv() string {
  rows := make([]string, m.M.NumRows()+1)

  rows[0] = strings.Join(m.ColumnNames, ",")
  for i := 1; i<m.M.NumRows()+1; i++ {
    rows[i] = strings.Join(m.Row(i-1), ",")
  }
  return strings.Join(rows, "\n")
}

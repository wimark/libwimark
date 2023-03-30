package libwimark

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/jung-kurt/gofpdf"
)

type timeBounds struct {
	Start int64 `json:"start" bson:"start"`
	Stop  int64 `json:"stop" bson:"stop"`
}

// StatReport struct for represent auto reports
type StatReport struct {
	ID string `json:"id" bson:"_id"`

	// basic fields
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`

	// location dat
	Location string `json:"-" bson:"location"`
	// subject of report (CPEs, Clients, Events, Custom)
	Subject ReportSubject `json:"subject" bson:"subject"`
	// CustomSubject

	// report type (current, summary)
	Type ReportType `json:"type" bson:"type"`

	// collect period (now, day, week, month)
	Period ReportPeriod `json:"collect_period" bson:"collect_period"`

	// period timebounds if Period == now
	TimeBounds timeBounds `json:"timebounds" bson:"timebounds"`

	// report format
	Format ReportFormat `json:"format" bson:"format"`

	// time to do post action (CRON string)
	ActionTime string `json:"action_time" bson:"action_time"`

	// type of post action (email, script)
	ActionType ReportActionType `json:"action_type" bson:"action_type"`

	// post action address dests
	ActionDest []string `json:"action_dest" bson:"action_dest"`

	// add charts
	Charts bool `json:"charts" bson:"charts"`
}

// StatReportResult struct for store collected reports
type StatReportResult struct {
	ID string `json:"id" bson:"_id"`

	Report StatReport  `json:"report_id" bson:"report_id"`
	Data   interface{} `json:"data" bson:"data"`

	CreateAt time.Time `json:"create_at" bson:"create_at"`
}

// GenerateFileReport return []byte as format: csv,pdf,xls
// format set as format variable
// set data as [][]string by table(columns and rows)
func GenerateFileReport(data [][]string, format string) ([]byte, error) {
	switch format {
	case "pdf":
		return generatePDFReport(data)
	case "csv":
		return generateCSVReport(data)
	case "xls":
		return generateXLSReport(data)
	default:
		return []byte{}, fmt.Errorf("unknown format %s", format)
	}
}

func generateCSVReport(data [][]string) ([]byte, error) {
	buffer := new(bytes.Buffer)
	writer := csv.NewWriter(buffer)
	writer.UseCRLF = true
	err := writer.WriteAll(data)
	return buffer.Bytes(), err
}

func generatePDFReport(data [][]string) ([]byte, error) {
	var buffer bytes.Buffer
	pdf := gofpdf.New("L", "mm", "Letter", "")

	pdf.AddPage()

	pdf.SetFont("Times", "", 20)
	pdf.Cell(40, 10, time.Now().Format("Mon Jan 2, 2006"))
	pdf.Ln(20)

	pdf.SetFont("Times", "", 10)
	pdf.SetFillColor(255, 255, 255)

	for _, line := range data {
		for _, str := range line {
			pdf.Cell(10, 7, str)
			//pdf.CellFormat(40, 7, str, "1", 0, "L", false, 0, "")
		}
		pdf.Ln(-1)
	}

	err := pdf.Output(&buffer)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func generateXLSReport(data [][]string) ([]byte, error) {
	f := excelize.NewFile()
	buffer := new(bytes.Buffer)
	for i, row := range data {
		for g, column := range row {
			f.SetCellValue("Sheet1", createXLSRowIndex(i+1, g+1), column)
		}
	}

	err := f.Write(buffer)
	return buffer.Bytes(), err
}

const letters = 26

var (
	xlsRowIndexes = [letters]string{"A", "B", "C", "D", "E", "F", "G", "H",
		"I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V",
		"W", "X", "Y", "Z"}
)

func createXLSRowIndex(row, column int) string {

	var indexColumn string

	for column > 0 {
		module := (column - 1) % letters
		indexColumn += xlsRowIndexes[module]
		column = (column - module) / letters
	}

	return fmt.Sprintf("%s%d", indexColumn, row)
}

package reports_controller

import (
	"AvitoInternship/internal/models"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type ReportsController struct{}

func CreateNewReportsController() *ReportsController {
	return &ReportsController{}
}

func (c *ReportsController) CreateFinancialReportCSV(serviceReport []models.Report, fileURL string) error {
	csvFile, err := os.Create(fileURL)
	if err == nil {
		defer csvFile.Close()
		writer := csv.NewWriter(csvFile)
		writer.Comma = 59
		for _, record := range serviceReport {
			err = writer.Write([]string{strconv.Itoa(record.ServiceId), fmt.Sprintf("%f", record.Sum)})
			if err != nil {
				break
			}
		}
		writer.Flush()
	} else {
		err = errors.New("Bad FilePath")
	}
	return err
}

package services

import (
	"encoding/csv"
	"io"
	"os"

	"flexera.com/purchaseManager/internal/common/errors"
	"flexera.com/purchaseManager/internal/common/logger"
	"flexera.com/purchaseManager/internal/common/utils"
)

var _ DataSetManager = new(csvDataSetManager)

type csvDataSetManager struct {
	log      logger.Logger
	filePath string
}

func NewCsvDataSetManager(filepath string, log logger.Logger) DataSetManager {
	return &csvDataSetManager{
		filePath: filepath,
	}
}

func (ds *csvDataSetManager) CreateReadStream(streamCh chan interface{}) error {
	filePath := ds.filePath

	f, err := os.Open(filePath)
	if err != nil {
		ds.log.Error("unable to read input file "+filePath, err)
		return errors.ErrOpenFile
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	// ignore the first line since it is the header
	csvReader.Read()

	for {
		row, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				close(streamCh)
				return nil
			}

			ds.log.Error("error occurred while reading the dataset row: ", err)
			return errors.ErrReadingCsvDataRow
		}

		parsedRow, err := utils.ParseAssetRow(row)
		if err != nil {
			ds.log.Error("error occurred while parsing csv data row")
			return errors.ErrParsingCSVDataRow
		}

		streamCh <- parsedRow
	}
}

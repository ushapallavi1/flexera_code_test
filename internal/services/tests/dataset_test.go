package tests

import (
	"testing"

	"flexera.com/purchaseManager/internal/common/logger"
	"flexera.com/purchaseManager/internal/models"
	"flexera.com/purchaseManager/internal/services"
	"github.com/stretchr/testify/suite"
)

type DataSetManagerTestSuite struct {
	suite.Suite
}

func TestDataSetTestSuite(t *testing.T) {
	suite.Run(t, new(DataSetManagerTestSuite))
}

func (s *DataSetManagerTestSuite) TestCreateReadStream() {

	log := logger.NewLogger()

	assetDataSetManager := services.NewCsvDataSetManager("./data/test1.csv", log)
	s.Require().NotNil(assetDataSetManager)

	streamCh := make(chan interface{})

	go assetDataSetManager.CreateReadStream(streamCh)

	for v := range streamCh {
		s.Require().IsType(&models.Asset{}, v)
	}
}

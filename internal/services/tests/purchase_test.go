package tests

import (
	"testing"

	"flexera.com/purchaseManager/internal/common/logger"
	"flexera.com/purchaseManager/internal/services"
	"github.com/stretchr/testify/suite"
)

type AssetPurchaseService struct {
	suite.Suite
}

func TestAssetPurchaseService(t *testing.T) {
	suite.Run(t, new(AssetPurchaseService))
}

func (s *AssetPurchaseService) TestGetMinimumPurchase() {

	log := logger.NewLogger()

	ps := services.NewAssetPurchaseService("374", log)
	s.Require().NotNil(ps)

	tests := []struct {
		name   string
		dm     services.DataSetManager
		output int
		err    error
	}{
		{
			name:   "simple data set",
			dm:     services.NewCsvDataSetManager("./data/test1.csv", log),
			output: 1,
			err:    nil,
		},
		{
			name:   "with latop and desktop for same user",
			dm:     services.NewCsvDataSetManager("./data/test2.csv", log),
			output: 3,
			err:    nil,
		},
		{
			name:   "with duplicate assets",
			dm:     services.NewCsvDataSetManager("./data/test3.csv", log),
			output: 2,
			err:    nil,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			result, err := ps.GetMinimumPurchase(tt.dm)
			s.Require().Equal(tt.err, err)
			s.Require().Equal(tt.output, result)
		})
	}
}


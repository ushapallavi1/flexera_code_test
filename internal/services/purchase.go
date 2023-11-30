package services

import (
	"flexera.com/purchaseManager/internal/common/errors"
	"flexera.com/purchaseManager/internal/common/logger"
	"flexera.com/purchaseManager/internal/common/utils"
	"flexera.com/purchaseManager/internal/models"
)

var _ PurchaseService = new(assetPurchaseService)

type assetPurchaseService struct {
	appId string
	log logger.Logger
}

func NewAssetPurchaseService(appId string, log logger.Logger) PurchaseService {
	return &assetPurchaseService{
		appId: appId,
		log: log,
	}
}

func (aps *assetPurchaseService) GetMinimumPurchase(dsManager DataSetManager) (int, error) {
	consumerCh := make(chan interface{}, 1000)

	go dsManager.CreateReadStream(consumerCh)

	processedAssets := make(map[string]struct{})
	count := make(map[string]int)

	numberOfCopiesRequired := 0

	for msg := range consumerCh {

		asset, ok := msg.(*models.Asset)
		if !ok {
			aps.log.Error("received invalid data after parsing data row")
			return -1, errors.ErrInvalidDataRow
		}

		if asset.ApplicationID != aps.appId {
			continue
		}

		key := utils.GenerateKey(asset.ComputerID, asset.UserID, asset.ApplicationID, string(asset.ComputerType))

		if _, isProcessed := processedAssets[key]; isProcessed {
			continue
		}

		processedAssets[key] = struct{}{}

		userAppkey := utils.GenerateKey(asset.UserID, asset.ApplicationID, string(asset.ComputerType))

		if asset.ComputerType == models.Desktop {
			laptopKey := utils.GenerateKey(asset.UserID, asset.ApplicationID, string(models.Laptop))
			if c, ok := count[laptopKey]; ok && c > 0 {
				count[laptopKey] -= 1
				continue
			}
		} else {
			desktopKey := utils.GenerateKey(asset.UserID, asset.ApplicationID, string(models.Desktop))
			if c, ok := count[desktopKey]; ok && c > 0 {
				count[desktopKey] -= 1
				continue
			}
		}

		count[userAppkey] += 1
		numberOfCopiesRequired += 1
	}

	return numberOfCopiesRequired, nil
}

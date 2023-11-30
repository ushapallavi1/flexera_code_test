package main

import (
	"flag"

	"flexera.com/purchaseManager/internal/common/logger"
	"flexera.com/purchaseManager/internal/services"
)

func main() {
	log := logger.NewLogger()

	csvDataFilePath := flag.String("csv", "", "csv dataset file path")
	appId := flag.String("appId", "", "application Id to calculate the purchase order")

	flag.Parse()

	if *csvDataFilePath == "" {
		log.Error("required -csv value")
		return
	}

	if *appId == "" {
		log.Error("required -appId value")
		return
	}

	assetDataSetManager := services.NewCsvDataSetManager(*csvDataFilePath, log)
	assetPurchaseService := services.NewAssetPurchaseService(*appId, log)

	log.Info("Processing data set from file: ", *csvDataFilePath)

	minimumPurchaseRequired, err := assetPurchaseService.GetMinimumPurchase(assetDataSetManager)
	if err != nil {
		log.Error("error occurred while calculating minimum purchase", err)
		return
	}

	log.Info("Data set processed successfully")
	log.Info("Minimum number of application copies required: ", minimumPurchaseRequired)

}

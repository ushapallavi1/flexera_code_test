package utils

import (
	"strings"

	"flexera.com/purchaseManager/internal/models"
)

func GenerateKey(values ...string) string {
	return strings.Join(values, "|")
}

func ParseAssetRow(row []string) (*models.Asset, error) {
	return models.NewAsset(row[0], row[1], row[2], row[3], row[4]), nil
}
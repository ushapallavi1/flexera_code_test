package models

import (
	"strings"
)

type AssetType string

const (
	Desktop AssetType = "desktop"
	Laptop  AssetType = "laptop"
)

type Asset struct {
	ComputerID    string
	UserID        string
	ApplicationID string
	ComputerType  AssetType
	Comment       string
}

func NewAsset(computerId, userId, appId, computerType, comment string) *Asset {
	return &Asset{
		ComputerID:    computerId,
		UserID:        userId,
		ApplicationID: appId,
		ComputerType:  AssetType(strings.ToLower(computerType)),
		Comment:       comment,
	}
}

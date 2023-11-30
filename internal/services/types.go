package services

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . DataSetManager
type DataSetManager interface {
	CreateReadStream(streamCh chan interface{}) error
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . PurchaseService
type PurchaseService interface {
	GetMinimumPurchase(dsManager DataSetManager) (int, error)
}

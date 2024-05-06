package services

import (
	"github.com/mikezzb/steam-trading-server/pkg/setting"
	"github.com/mikezzb/steam-trading-server/util"
	"github.com/mikezzb/steam-trading-shared/database/model"
)

type Transaction struct {
	ID string
}

func (s *Transaction) Count(filters *util.ItemFilters) (int64, error) {
	return transactionRepo.Count(filters.GetBsonFilters())
}

func (s *Transaction) GetTransactionsByDays(days int, filters *util.ItemFilters) ([]model.Transaction, error) {
	return transactionRepo.FindItemByDays(days, filters.GetBsonFilters())
}

func (s *Transaction) GetTransactionsByPage(page int, filters *util.ItemFilters) ([]model.Transaction, error) {
	bsonFilters := filters.GetBsonFilters()
	return transactionRepo.FindTransactionsByPage(page, setting.App.TransactionPageSize, bsonFilters)
}

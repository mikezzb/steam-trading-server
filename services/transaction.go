package services

import (
	"github.com/mikezzb/steam-trading-server/pkg/setting"
	"github.com/mikezzb/steam-trading-server/types"
	"github.com/mikezzb/steam-trading-shared/database/model"
)

type Transaction struct {
	ID   string
	Page int
	// filters
	ItemFilters *types.ItemFilters
}

func (s *Transaction) Count() (int64, error) {
	return transactionRepo.Count()
}

func (s *Transaction) GetTransactions() ([]model.Transaction, error) {
	return transactionRepo.FindTransactionsByPage(s.Page, setting.App.TransactionPageSize, s.ItemFilters.GetBsonFilters())
}

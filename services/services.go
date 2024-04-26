package services

import (
	"github.com/mikezzb/steam-trading-server/db"
	"github.com/mikezzb/steam-trading-shared/database/repository"
)

var itemRepo *repository.ItemRepository
var userRepo *repository.UserRepository
var listingRepo *repository.ListingRepository
var transactionRepo *repository.TransactionRepository
var subRepo *repository.SubscriptionRepository

func Setup() {
	repos := db.Repos
	itemRepo = repos.GetItemRepository()
	userRepo = repos.GetUserRepository()
	listingRepo = repos.GetListingRepository()
	transactionRepo = repos.GetTransactionRepository()
	subRepo = repos.GetSubscriptionRepository()
}

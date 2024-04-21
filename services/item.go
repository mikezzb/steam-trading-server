package services

import (
	"strconv"
	"strings"
	"time"

	"github.com/mikezzb/steam-trading-server/cache"
	"github.com/mikezzb/steam-trading-server/db"
	"github.com/mikezzb/steam-trading-server/pkg/setting"
	"github.com/mikezzb/steam-trading-shared/database/model"
	"github.com/mikezzb/steam-trading-shared/database/repository"
)

var itemRepo *repository.ItemRepository = nil

type ItemService struct {
	PageNum int
}

func getItemRepo() *repository.ItemRepository {
	if itemRepo == nil {
		itemRepo = db.Repos.GetItemRepository()
	}
	return itemRepo
}

func (s *ItemService) Count() (int64, error) {
	val, err := cache.UseCache(
		"ITEM_TOTAL",
		1*time.Minute,
		func() (interface{}, error) {
			return getItemRepo().Count()
		},
	)

	if err != nil {
		return 0, err
	}

	return val.(int64), nil
}

func (s *ItemService) GetItems() ([]model.Item, error) {
	return getItemRepo().GetItemsByPage(s.PageNum, setting.App.ItemPageSize, nil)
}

func (s *ItemService) GetItemByName(id string) (*model.Item, error) {
	return getItemRepo().GetItemByName(id)
}

func (s *ItemService) getCacheKey() string {
	keys := []string{
		"ITEM",
		strconv.Itoa(s.PageNum),
	}

	return strings.Join(keys, "_")
}

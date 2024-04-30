package services

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/mikezzb/steam-trading-server/cache"
	"github.com/mikezzb/steam-trading-shared/database/model"
)

type Item struct {
	PageNum int
}

func (s *Item) Count() (int64, error) {
	val, err := cache.UseCache(
		"ITEM_TOTAL",
		1*time.Minute,
		func() (interface{}, error) {
			return itemRepo.Count()
		},
	)

	if err != nil {
		return 0, err
	}

	return val.(int64), nil
}

func (s *Item) GetItem(id string) (*model.Item, error) {
	return itemRepo.FindItemById(id)
}

func (s *Item) GetItems(pageSize int) ([]model.Item, error) {
	log.Printf("GetItems: %v", s.PageNum)
	items, err := itemRepo.GetItemsByPage(s.PageNum, pageSize, nil)
	if err != nil || items == nil {
		return make([]model.Item, 0), err
	}
	return items, nil
}

func (s *Item) GetItemByName(id string) (*model.Item, error) {
	return itemRepo.GetItemByName(id)
}

func (s *Item) getCacheKey() string {
	keys := []string{
		"ITEM",
		strconv.Itoa(s.PageNum),
	}

	return strings.Join(keys, "_")
}

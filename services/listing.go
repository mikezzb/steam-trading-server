package services

import (
	"github.com/mikezzb/steam-trading-server/pkg/setting"
	"github.com/mikezzb/steam-trading-server/util"
	"github.com/mikezzb/steam-trading-shared/database/model"
)

type Listing struct {
}

func (s *Listing) GetListingsByPage(page int, filters *util.ItemFilters) ([]model.Listing, error) {
	return listingRepo.GetListingsByPage(page, setting.App.ListingPageSize, filters.GetBsonFilters())
}

func (s *Listing) Count(filters *util.ItemFilters) (int64, error) {
	return listingRepo.Count(filters.GetBsonFilters())
}

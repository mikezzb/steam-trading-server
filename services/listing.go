package services

import "github.com/mikezzb/steam-trading-shared/database/model"

type Listing struct {
	PageNum int
}

func (s *Listing) GetListings() ([]model.Listing, error) {
	return listingRepo.GetListingsByPage(s.PageNum, 10, nil)
}

func (s *Listing) Count() (int64, error) {
	return listingRepo.Count()
}

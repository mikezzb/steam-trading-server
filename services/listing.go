package services

import "github.com/mikezzb/steam-trading-shared/database/model"

type Listing struct {
	Page int
}

func (s *Listing) GetListings() ([]model.Listing, error) {
	return listingRepo.GetListingsByPage(s.Page, 10, nil)
}

func (s *Listing) Count() (int64, error) {
	return listingRepo.Count()
}

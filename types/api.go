package types

import (
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
)

type ItemFilters struct {
	Rarity    string
	PaintSeed string
	// item name
	Name string
}

func (s *ItemFilters) GetBsonFilters() bson.M {
	filters := bson.M{}

	if s.Rarity != "" {
		filters["rarity"] = s.Rarity
	}

	if s.PaintSeed != "" {
		filters["paintSeed"] = s.PaintSeed
	}

	return filters
}

func (s *ItemFilters) AddFilter(key, value string) {
	switch key {
	case "rarity":
		s.Rarity = value
	case "paintSeed":
		s.PaintSeed = value
	}
}

func NewItemFilters(params url.Values) *ItemFilters {
	s := &ItemFilters{}
	for k, v := range params {
		s.AddFilter(k, v[0])
	}
	return s
}

package types

import (
	"net/url"
	"strings"

	"github.com/mikezzb/steam-trading-server/constants"
	"go.mongodb.org/mongo-driver/bson"
)

type ItemFilters struct {
	Rarity    string
	PaintSeed string
	// item name
	Name string

	// optional
	Category  string
	Skin      string
	Exteriors []string
}

func (s *ItemFilters) GetBsonFilters() bson.M {
	filters := bson.M{}

	if s.Rarity != "" {
		filters["rarity"] = s.Rarity
	}

	if s.PaintSeed != "" {
		filters["paintSeed"] = s.PaintSeed
	}

	if s.Name != "" {
		filters["name"] = s.Name
	}

	if s.Category != "" {
		filters["category"] = s.Category
	}

	if s.Skin != "" {
		filters["skin"] = s.Skin
	}

	if len(s.Exteriors) > 0 {
		filters["exterior"] = bson.M{"$in": s.Exteriors}
	}

	return filters
}

func (s *ItemFilters) AddFilter(key, value string) {
	switch key {
	case "rarity":
		s.Rarity = value
	case "paintSeed":
		s.PaintSeed = value
	case "name":
		s.Name = value
	case "category":
		s.Category = value
	case "skin":
		s.Skin = value
	case "exterior":
		exts := strings.Split(value, ",")
		// transform abbv to full name
		for i, abbr := range exts {
			if full, ok := constants.ExteriorFull[abbr]; ok {
				exts[i] = full
			}
		}
		s.Exteriors = exts
	}
}

func NewItemFilters(params url.Values) *ItemFilters {
	s := &ItemFilters{}
	for k, v := range params {
		s.AddFilter(k, v[0])
	}
	return s
}

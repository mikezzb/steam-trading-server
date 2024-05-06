package services

import (
	"github.com/mikezzb/steam-trading-shared/database/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddSubForm struct {
	Name       string   `json:"name" valid:"Required;"`
	Rarities   []string `bson:"rarities,omitempty" json:"rarities"`
	PaintSeeds []int    `bson:"paintSeeds,omitempty" json:"paintSeeds"`
	MaxPremium string   `json:"maxPremium" valid:"Required;"`
	NotiType   string   `json:"notiType" valid:"Required;"`
	NotiId     string   `json:"notiId" valid:"Required;"`
}

type UpdateSubForm struct {
	ID string `form:"id" valid:"Required;"`

	Name       string   `json:"name" valid:"Required;"`
	Rarities   []string `bson:"rarities,omitempty" json:"rarities"`
	PaintSeeds []int    `bson:"paintSeeds,omitempty" json:"paintSeeds"`
	MaxPremium string   `json:"maxPremium" valid:"Required;"`
	NotiType   string   `json:"notiType" valid:"Required;"`
	NotiId     string   `json:"notiId" valid:"Required;"`
}

type Subscription struct {
	ID         string
	Name       string
	Rarities   []string
	PaintSeeds []int

	MaxPremium string
	NotiType   string
	NotiId     string
	OwnerId    primitive.ObjectID

	// Non model fields
}

func (s *Subscription) getModel() *model.Subscription {
	var id primitive.ObjectID = primitive.NilObjectID
	if s.ID != "" {
		id, _ = primitive.ObjectIDFromHex(s.ID)
	}
	return &model.Subscription{
		ID:         id,
		Name:       s.Name,
		Rarities:   s.Rarities,
		PaintSeeds: s.PaintSeeds,
		MaxPremium: s.MaxPremium,
		NotiType:   s.NotiType,
		NotiId:     s.NotiId,
		OwnerId:    s.OwnerId,
	}
}

func (s *Subscription) AddSub() (string, error) {
	id, err := subRepo.InsertSubscription(s.getModel())
	if err != nil {
		return "", err
	}
	return id.Hex(), nil
}

func (s *Subscription) UpdateSub() error {
	return subRepo.UpdateSubscription(s.getModel())
}

func (s *Subscription) DeleteSub() error {
	return subRepo.DeleteSubscriptionByName(s.Name, s.OwnerId)
}

func (s *Subscription) GetSubs() ([]model.Subscription, error) {
	return subRepo.GetAllByOwnerId(s.OwnerId)
}

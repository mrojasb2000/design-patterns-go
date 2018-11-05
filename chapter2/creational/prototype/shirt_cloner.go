package prototype

import (
	"errors"
	"fmt"
)

const (
	White = 1
	Black = 2
	Blue  = 3
)

// Package-level function
func GetShirtsCloner() ShirtCloner {
	return nil
}

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

type ShirtsCache struct{}

// Implement interface ShirtCloner
func (s *ShirtsCache) GetClone(m int) (ItemInfoGetter, error) {
	switch m {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("Shirt model not recognized")
	}
}

type ShirtColor byte

type ItemInfoGetter interface {
	GetInfo() string
}

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

// Implement interface ItemInfoGetter
func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt withSKU '%s' and Color id '%d' that costs %f\n", s.SKU, s.Color, s.Price)
}

var whitePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

var blackPrototype *Shirt = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: Black,
}

var bluePrototype *Shirt = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}

func (i *Shirt) GetPrice() float32 {
	return i.Price
}

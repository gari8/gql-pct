package domain

import (
	"fmt"
	"io"
	"strconv"
)

type Place struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	PlaceType PlaceType `json:"placeType"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
}

func (e PlaceType) ToInt() int {
	for i, v := range AllPlaceType {
		if v == e {
			return i
		}
	}
	return -1
}

type PlaceType string

const (
	PlaceTypeLiveStage  PlaceType = "LIVE_STAGE"
	PlaceTypeEventBooth PlaceType = "EVENT_BOOTH"
)

var AllPlaceType = []PlaceType{
	PlaceTypeLiveStage,
	PlaceTypeEventBooth,
}

func (e PlaceType) IsValid() bool {
	switch e {
	case PlaceTypeLiveStage, PlaceTypeEventBooth:
		return true
	}
	return false
}

func (e PlaceType) String() string {
	return string(e)
}

func (e *PlaceType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PlaceType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PlaceType", str)
	}
	return nil
}

func (e PlaceType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

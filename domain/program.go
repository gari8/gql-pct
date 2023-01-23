package domain

import (
	"fmt"
	"io"
	"strconv"
)

type Program struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Image       string      `json:"image"`
	ProgramType ProgramType `json:"programType"`
	PlaceID     string      `json:"placeId"`
}

func (e ProgramType) ToInt() *int {
	for i, v := range AllProgramType {
		if v == e {
			return &i
		}
	}
	return nil
}

type ProgramType string

const (
	ProgramTypeWorkshop   ProgramType = "WORKSHOP"
	ProgramTypeConference ProgramType = "CONFERENCE"
	ProgramTypeMeeting    ProgramType = "MEETING"
	ProgramTypeEvent      ProgramType = "EVENT"
)

var AllProgramType = []ProgramType{
	ProgramTypeWorkshop,
	ProgramTypeConference,
	ProgramTypeMeeting,
	ProgramTypeEvent,
}

func (e ProgramType) IsValid() bool {
	switch e {
	case ProgramTypeWorkshop, ProgramTypeConference, ProgramTypeMeeting, ProgramTypeEvent:
		return true
	}
	return false
}

func (e ProgramType) String() string {
	return string(e)
}

func (e *ProgramType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProgramType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProgramType", str)
	}
	return nil
}

func (e ProgramType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

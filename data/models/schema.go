package models

import (
	"errors"
	"github.com/elos/server/data"
)

/*



*/

var UndefinedKindError = errors.New("Error: undefined kind")
var UndefinedLinkError = errors.New("Error: undefined link")
var UndefinedLinkKindError = errors.New("Error: undefined link kind")

type LinkKind string

const MulLink LinkKind = "MANY"
const OneLink LinkKind = "ONE"

type Schema map[data.Kind]map[data.Kind]LinkKind

// add a function to check validity of a schema

func PossibleLink(s *Schema, this Model, other Model) (bool, error) {
	thisKind := this.Kind()

	links, ok := (*s)[thisKind]

	if !ok {
		return false, UndefinedKindError
	}

	otherKind := other.Kind()

	_, linkPossible := links[otherKind]

	if !linkPossible {
		return false, UndefinedLinkError
	}

	return true, nil
}

func (s *Schema) LinkType(this Model, other Model) (LinkKind, error) {
	_, err := PossibleLink(s, this, other)
	if err != nil {
		return "", err
	}

	return (*s)[this.Kind()][other.Kind()], nil
}

func LinkWith(lk LinkKind, this Model, that Model) error {
	switch lk {
	case MulLink:
		// this.LinkMul(that)
	case OneLink:
		// this.LinkOne(that)
	default:
		return UndefinedLinkKindError
	}

	return nil
}

func UnlinkWith(ln LinkKind, this Model, that Model) error {
	switch ln {
	case MulLink:
		// this.UnlinkMul(that)
	case OneLink:
		// this.UnlinkOne(that)
	default:
		return UndefinedLinkKindError
	}

	return nil
}

func (s *Schema) Link(this Model, that Model) error {
	thisLinkType, err := s.LinkType(this, that)

	if err != nil {
		return err
	} else {
		if err = LinkWith(thisLinkType, this, that); err != nil {
			return err
		}
	}

	thatLinkType, err := s.LinkType(this, that)

	if err != nil {
		return err
	} else {
		if err = LinkWith(thatLinkType, this, that); err != nil {
			return err
		}
	}

	return nil
}

func (s *Schema) Unlink(this Model, that Model) error {
	thisLinkType, err := s.LinkType(this, that)
	if err != nil {
		return err
	} else {
		if err = UnlinkWith(thisLinkType, this, that); err != nil {
			return err
		}
	}

	thatLinkType, err := s.LinkType(this, that)

	if err != nil {
		return err
	} else {
		if err = UnlinkWith(thatLinkType, this, that); err != nil {
			return err
		}
	}

	return nil
}

var ElosSchema Schema = map[data.Kind]map[data.Kind]LinkKind{
	UserKind: {
		EventKind: MulLink,
	},
	EventKind: {
		UserKind: OneLink,
	},
}

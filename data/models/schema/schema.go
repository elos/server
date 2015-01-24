package schema

import (
	"errors"
	"github.com/elos/server/data"
	"github.com/elos/server/data/models"
)

var UndefinedKindError = errors.New("Error: undefined kind")
var UndefinedLinkError = errors.New("Error: undefined link")
var UndefinedLinkKindError = errors.New("Error: undefined link kind")
var InvalidSchemaError = errors.New("Error: invalid schema")

type LinkKind string

const MulLink LinkKind = "MANY"
const OneLink LinkKind = "ONE"

type SchemaMap map[data.Kind]map[data.Kind]LinkKind

func (s *SchemaMap) Valid() bool {
	for outerKind, links := range *s {
		for innerKind, _ /*linkKind*/ := range links {
			innerLinks, ok := (*s)[innerKind]
			if !ok {
				return false
			}

			_ /*matchingLink*/, ok = innerLinks[outerKind]

			if !ok {
				return false
			}
		}
	}

	return true
}

func PossibleLink(s *SchemaMap, this models.Model, other models.Model) (bool, error) {
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

func (s *SchemaMap) LinkType(this models.Model, other models.Model) (LinkKind, error) {
	_, err := PossibleLink(s, this, other)
	if err != nil {
		return "", err
	}

	return (*s)[this.Kind()][other.Kind()], nil
}

func LinkWith(lk LinkKind, this models.Model, that models.Model) error {
	switch lk {
	case MulLink:
		this.LinkMul(that)
	case OneLink:
		this.LinkOne(that)
	default:
		return UndefinedLinkKindError
	}

	return nil
}

func UnlinkWith(ln LinkKind, this models.Model, that models.Model) error {
	switch ln {
	case MulLink:
		this.UnlinkMul(that)
	case OneLink:
		this.UnlinkOne(that)
	default:
		return UndefinedLinkKindError
	}

	return nil
}

func (s *SchemaMap) Link(this models.Model, that models.Model) error {
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
		if err = LinkWith(thatLinkType, that, this); err != nil {
			return err
		}
	}

	return nil
}

func (s *SchemaMap) Unlink(this models.Model, that models.Model) error {
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
		if err = UnlinkWith(thatLinkType, that, this); err != nil {
			return err
		}
	}

	return nil
}

type VersionedSchema struct {
	*SchemaMap
	version int
}

func NewSchema(sm *SchemaMap, version int) (models.Schema, error) {
	s := &VersionedSchema{
		SchemaMap: sm,
		version:   version,
	}

	if !s.Valid() {
		return nil, InvalidSchemaError
	}

	return s, nil
}

func (s *VersionedSchema) GetVersion() int {
	return s.version
}

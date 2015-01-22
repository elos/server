package user

import (
	"github.com/elos/server/data"
	"github.com/elos/server/data/models"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type MongoUser struct {
	LoadedAt time.Time `json:"-" bson:"-"`

	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`

	// Properties
	Name string `json:"name"`
	Key  string `json:"key"`

	// Links
	EventIds []data.ID `json:"event_ids", bson:"event_ids"`
}

func (u *MongoUser) Kind() data.Kind {
	return models.UserKind
}

func (u *MongoUser) SetID(id data.ID) {
	vid, ok := id.(bson.ObjectId)
	if ok {
		u.ID = vid
	}
}

func (u *MongoUser) GetID() data.ID {
	return u.ID
}

func (u *MongoUser) SetName(name string) {
	u.Name = name
}

func (u *MongoUser) GetName() string {
	return u.Name
}

func (u *MongoUser) Save() error {
	return db.Save(u)
}

func (u *MongoUser) Concerned() []data.ID {
	a := make([]data.ID, 1)
	a[0] = u.ID
	return a
}

func (u *MongoUser) AddEvent(eventId data.ID) error {
	if err := data.CheckID(eventId); err != nil {
		return err
	}

	if !u.EventIdsHash()[eventId] {
		u.EventIds = append(u.EventIds, eventId)
		return u.Save()
	}

	return nil
}

func (u *MongoUser) RemoveEvent(eventId data.ID) error {
	if err := data.CheckID(eventId); err != nil {
		return err
	}

	eventIds := u.EventIdsHash()

	if eventIds[eventId] {
		eventIds[eventId] = false
		ids := make([]data.ID, 0)
		for id := range eventIds {
			if eventIds[id] {
				ids = append(ids, id)
			}
		}

		u.EventIds = ids
		return u.Save()
	}

	return nil
}

func (u *MongoUser) EventIdsHash() map[data.ID]bool {
	hash := make(map[data.ID]bool, len(u.EventIds))

	for _, id := range u.EventIds {
		hash[id] = true
	}

	return hash
}

func (u *MongoUser) SetCreatedAt(t time.Time) {
	u.CreatedAt = t
}

func (u *MongoUser) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *MongoUser) SetUpdatedAt(t time.Time) {
	u.UpdatedAt = t
}

func (u *MongoUser) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}

func (u *MongoUser) GetLoadedAt() time.Time {
	return u.LoadedAt
}

func (u *MongoUser) SetKey(s string) {
	u.Key = s
}

func (u *MongoUser) GetKey() string {
	return u.Key
}

package models

import (
	"time"

	"github.com/elos/data"
	"gopkg.in/mgo.v2/bson"
)

/*
	Note these structs have random
	letters before fields, this is so that
	they are visible and are marshalled by json and bson

	definitely a hack -- working on it
*/

type Based struct {
	Bid        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	BcreatedAt time.Time     `json:"created_at" bson:"created_at"`
	BupdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}

func (b *Based) ID() data.ID {
	return b.Bid
}

func (b *Based) SetID(id data.ID) {
	vid, ok := id.(bson.ObjectId)
	if ok {
		b.Bid = vid
	}
}

func (b *Based) CreatedAt() time.Time {
	return b.BcreatedAt
}

func (b *Based) SetCreatedAt(t time.Time) {
	b.BcreatedAt = t
}

func (b *Based) UpdatedAt() time.Time {
	return b.BupdatedAt
}

func (b *Based) SetUpdatedAt(t time.Time) {
	b.BupdatedAt = t
}

type Named struct {
	Nname string `json:"name"`
}

func (n *Named) Name() string {
	return n.Nname
}

func (n *Named) SetName(name string) {
	n.Nname = name
}

type Keyed struct {
	key string `json:"key"`
}

func (k *Keyed) Key() string {
	return k.key
}

func (k *Keyed) SetKey(key string) {
	k.key = key
}

type Timed struct {
	startTime time.Time `json:"start_time" bson:"start_time"`
	endTime   time.Time `json:"end_time" bson:"end_time"`
}

func (t *Timed) EndTime() time.Time {
	return t.endTime
}

func (t *Timed) SetEndTime(eTime time.Time) {
	t.endTime = eTime
}

func (t *Timed) StartTime() time.Time {
	return t.startTime
}

func (t *Timed) SetStartTime(sTime time.Time) {
	t.startTime = sTime
}

package user_test

import (
	. "github.com/elos/server/models/user"
	"time"

	"github.com/elos/server/data"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/mgo.v2/bson"
)

var _ = Describe("Type", func() {
	It("Defines it's kind property", func() {
		var k data.Kind
		Expect(Kind).ToNot(BeNil())
		Expect(Kind).To(BeAssignableToTypeOf(k))
	})

	Describe("Definition", func() {
		It("Defines required fields", func() {
			n := "hello"
			key := "asdlfkjasdjkf"
			time := time.Now()

			u := &User{
				Id:        bson.NewObjectId(),
				CreatedAt: time,
				Name:      n,
				Key:       key,
			}

			Expect(u).ToNot(BeNil())
			Expect(u.Id).ToNot(BeNil())
			Expect(u.Id.Valid()).To(BeTrue())
			Expect(u.CreatedAt).To(Equal(time))
			Expect(u.Name).To(Equal(n))
			Expect(u.Key).To(Equal(key))

			eventIds := make([]bson.ObjectId, 2)
			eventIds[0] = bson.NewObjectId()
			eventIds[1] = bson.NewObjectId()

			// Test that we can assign these
			u.EventIds = eventIds
			Expect(u.EventIds).To(Equal(eventIds))
		})
	})

	Describe("Constructors", func() {
		Describe("New", func() {
			It("Returns a new user model", func() {
				u := New()
				Expect(u).To(BeAssignableToTypeOf(&User{}))
				var emptyId bson.ObjectId
				Expect(u.Id).To(Equal(emptyId))
				var emptyTime time.Time
				Expect(u.CreatedAt).To(Equal(emptyTime))
				var emptyString string
				Expect(u.Key).To(Equal(emptyString))
				Expect(u.Name).To(Equal(emptyString))
			})
		})

		Describe("Create", func() {
		})
	})
})

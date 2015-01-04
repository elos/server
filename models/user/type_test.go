package user_test

/*
import (
	. "github.com/elos/server/models/user"
	"time"

	"github.com/elos/server/data"
	"github.com/elos/server/data/test"
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

	PDescribe("Constructors", func() {
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

		PDescribe("Create", func() {
			var (
				db  *test.TestDB
				err error
			)

			BeforeEach(func() {
				// Get a new test database
				db, err = test.NewDB()
				It("Should create a database and not error", func() {
					Expect(err).ToNot(HaveOccurred())
					Expect(db).ToNot(BeNil())
				})
			})

			AfterEach(func() {
				db.Reset()
			})

			PContext("Correct model, should save", func() {
				// Set the user database to be our test database
				DB = db

				n := "This is my name"
				u, err := Create(n)

				PIt("Doesn't error", func() {
					Expect(err).ToNot(HaveOccurred())
					Expect(u).ToNot(BeNil())
				})

				PIt("Sets up proper fields", func() {
					Expect(u.GetId().Valid()).To(BeTrue())
					Expect(u.CreatedAt).To(BeAssignableToTypeOf(time.Now()))
					Expect(u.Name).To(Equal(n))
					Expect(u.Key).ToNot(Equal("akdljfasjdkf"))
					Expect(u.Key).To(HaveLen(64))
				})

				PIt("Saves to the database", func() {
					Expect(u).To(Equal(db.Saved[0]))
				})
			})

			PContext("Database set to error, should set up but not save", func() {
				// Set the user database to be our test database
				DB = db
				db.Error = true

				n := "this user will fail creation"
				u, err := Create(n)

				PIt("Error", func() {
					Expect(err).To(HaveOccurred())
				})

				PIt("Still returns a valid user", func() {
					Expect(u).ToNot(BeNil())
				})

				PIt("Still sets up fields", func() {
					Expect(err).To(HaveOccurred())
					Expect(u.Name).To(Equal(n))
					Expect(u.GetId().Valid()).To(BeTrue())
					Expect(u.CreatedAt).To(BeAssignableToTypeOf(time.Now()))
				})

				PIt("Should not save", func() {
					Expect(db.Saved).To(HaveLen(0))
				})

				db.Error = false
			})

		})
	})
})
*/

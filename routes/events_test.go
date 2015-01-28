package routes_test

import (
	. "github.com/elos/server/routes"

	"net/http"
	"net/http/httptest"
	"net/url"
	"time"

	"github.com/elos/data/test"
	"github.com/elos/server/models/event"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/mgo.v2/bson"
)

var _ = Describe("Events", func() {

	It("Exposes a EventsPost HAndler", func() {
		Expect(EventsPost).NotTo(BeNil())
	})

	Describe("EventsPostHandler", func() {
		db := test.NewDB()
		event.SetDB(db)

		w := httptest.NewRecorder()

		values := url.Values{}
		name := "this is the name"
		userId := bson.NewObjectId()
		start_time := time.Now()
		end_time := time.Now()
		values.Add("name", name)
		values.Add("start_time", start_time.String())
		values.Add("end_time", end_time.String())
		values.Add("user_id", userId.Hex())

		r := &http.Request{
			Form: values,
		}

		n1 := NewNullHandler()
		n2 := NewNullHandler()

		var re error

		var errHandlerC = func(err error) http.Handler {
			re = err
			return n1
		}

		var (
			rc int
			rr interface{}
		)

		var resourceHandlerC = func(c int, r interface{}) http.Handler {
			rc = c
			rr = r
			return n2
		}

		It("Handles a database error", func() {
			db.Error = true

			EventsPostHandler(w, r, errHandlerC, resourceHandlerC)

			By("Calling the error handler")
			Expect(n1.Handled).To(HaveKeyWithValue(r, true))
			By("Using the db error")
			Expect(re).To(Equal(test.TestDBError))
			By("Not touching the resource response")
			Expect(n2.Handled).To(BeEmpty())

			db.Reset()
			n1.Reset()
		})

		It("Successfully creates an event", func() {
			EventsPostHandler(w, r, errHandlerC, resourceHandlerC)
			By("Calling the resource response handler")
			Expect(n2.Handled).To(HaveKeyWithValue(r, true))
			Expect(rc).To(Equal(201))
			Expect(db.Saved).To(HaveLen(1))
			Expect(rr).To(Equal(db.Saved[0]))

			By("Not touching the error handler")
			Expect(n1.Handled).To(BeEmpty())

			By("Transferring the form value name to the user name")
			e := db.Saved[0]

			var ok bool
			e, ok = e.(*event.Event)
			Expect(ok).To(BeTrue())
			Expect(e.(*event.Event).Name).To(Equal(name))
			// Expect(e.(*event.Event).StartTime).To(Equal(start_time))
			// Expect(e.(*event.Event).EndTime).To(Equal(end_time))

			db.Reset()
			n2.Reset()
		})
	})

})

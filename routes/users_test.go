package routes_test

import (
	. "github.com/elos/server/routes"

	"github.com/elos/server/data/models/user"
	"github.com/elos/server/data/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"net/url"
)

var _ = Describe("Users", func() {
	It("Exposes a UsersPost Handler", func() {
		Expect(UsersPost).NotTo(BeNil())
	})

	Describe("UsersPostFunction", func() {
		db := test.NewDB()
		user.SetDB(db)

		w := httptest.NewRecorder()

		values := url.Values{}
		name := "this is the name"
		values.Add("name", name)
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

		It("Handles a database errors", func() {
			db.Error = true
			UsersPostFunction(w, r, errHandlerC, resourceHandlerC)

			By("Calling the error handler")
			Expect(n1.Handled).To(HaveKeyWithValue(r, true))
			By("Using the db error")
			Expect(re).To(Equal(test.TestDBError))
			By("Not touching the resource response")
			Expect(n2.Handled).To(BeEmpty())

			db.Reset()
			n1.Reset()
		})

		It("Handles a successfuly user creation", func() {
			UsersPostFunction(w, r, errHandlerC, resourceHandlerC)

			By("Calling the resource response handler")
			Expect(n2.Handled).To(HaveKeyWithValue(r, true))
			Expect(rc).To(Equal(201))
			Expect(db.Saved).To(HaveLen(1))
			Expect(rr).To(Equal(db.Saved[0]))

			By("Not touching the error handler")
			Expect(n1.Handled).To(BeEmpty())

			db.Reset()
			n2.Reset()
		})

	})
})

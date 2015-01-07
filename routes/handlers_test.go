package routes_test

import (
	. "github.com/elos/server/routes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"errors"
	"github.com/elos/server/data"
	"github.com/elos/server/data/models/user"
	"github.com/elos/server/util"
	"github.com/elos/server/util/auth"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Handlers", func() {

	// NullHandler {{{
	Describe("NullHandler", func() {
		h := NewNullHandler()

		Describe("NewHullHandler", func() {
			It("Allocates and returns a new NullHandler", func() {
				Expect(h).ToNot(BeNil())
				Expect(h).To(BeAssignableToTypeOf(&NullHandler{}))
				Expect(h.Handled).ToNot(BeNil())
				Expect(h.Handled).To(BeEmpty())
			})
		})

		Describe("ServeHTTP", func() {
			It("Adds a request to it's handled map when asked to serve it", func() {
				r := &http.Request{Method: "FOOBAR"}
				w := httptest.NewRecorder()
				h.ServeHTTP(w, r)
				Expect(h.Handled).To(HaveKeyWithValue(r, true))
			})
		})

		Describe("Reset()", func() {
			It("Wipes it Handled map", func() {
				h.Reset()
				Expect(h.Handled).To(BeEmpty())
			})
		})
	})
	// NullHandler }}}

	// ErrorHandler {{{
	Describe("ErrorHandler", func() {
		err := errors.New("This is a test error")
		h := NewErrorHandler(err)

		Describe("NewErrorHandler", func() {
			It("Allocates and returns a new ErrorHandler", func() {
				Expect(h).ToNot(BeNil())
				Expect(h).To(BeAssignableToTypeOf(&ErrorHandler{}))
				Expect(h.(*ErrorHandler).Err).To(Equal(err))
			})
		})

		Describe("ServeHTTP", func() {
			w1 := httptest.NewRecorder()
			w2 := httptest.NewRecorder()
			util.ServerError(w1, err)
			h.ServeHTTP(w2, &http.Request{})
			It("Uses util to write the error response", func() {
				Expect(w1.Body).To(Equal(w2.Body))
			})
		})

	})
	// ErrorHandler }}}

	// ResourceHandler {{{
	Describe("ResoureHandler", func() {
		resource := map[string]interface{}{
			"hey": "ho",
		}

		code := 200

		h := NewResourceHandler(code, resource)

		Describe("NewResourceHandler", func() {
			It("Allocates and returns a new ResourceHandler", func() {
				Expect(h).ToNot(BeNil())
				Expect(h).To(BeAssignableToTypeOf(&ResourceHandler{}))
			})

			It("Sets up necessary information", func() {
				By("Setting status code")
				Expect(h.(*ResourceHandler).Code).To(Equal(code))
				By("Setting resource")
				Expect(h.(*ResourceHandler).Resource).To(Equal(resource))
			})
		})

		Describe("ServeHTTP", func() {
			w1 := httptest.NewRecorder()
			w2 := httptest.NewRecorder()
			util.WriteResourceResponse(w1, code, resource)
			h.ServeHTTP(w2, &http.Request{})
			It("Uses util.WriteResourceResponse", func() {
				Expect(w1.Body).To(Equal(w2.Body))
				Expect(w1.Code).To(Equal(w2.Code))
			})
		})

	})
	// ResourceHandler }}}

	// BadMethodHandler {{{
	Describe("BadMethodHandler", func() {
		r := &http.Request{Method: "BOOP"}
		h := NewBadMethodHandler(r)

		Describe("NewBadMethodHandler", func() {
			It("Should allocate and return a new BadMethodHandler", func() {
				Expect(h).ToNot(BeNil())
				Expect(h).To(BeAssignableToTypeOf(&BadMethodHandler{}))
			})

			It("Should set it's RequestedMethod field", func() {
				Expect(h.(*BadMethodHandler).RequestedMethod).To(Equal(r.Method))
			})
		})

		Describe("ServeHTTP", func() {
			w1 := httptest.NewRecorder()
			w2 := httptest.NewRecorder()
			util.InvalidMethod(w1)
			h.ServeHTTP(w2, &http.Request{})
			It("Uses util.InvalidMethod", func() {
				Expect(w1.Body).To(Equal(w2.Body))
				Expect(w1.Code).To(Equal(w2.Code))
			})
		})
	})
	// BadMethodHandler }}}

	// UnauthorizedHandler {{{
	Describe("UnauthorizedHandler", func() {
		reason := "asdf"
		h := NewUnauthorizedHandler(reason)

		Describe("NewUnauthorizedHandler", func() {
			It("Allocates and returns a new UnauthorizedHandler", func() {
				Expect(h).ToNot(BeNil())
				Expect(h).To(BeAssignableToTypeOf(&UnauthorizedHandler{}))
			})

			It("Sets the reason field", func() {
				Expect(h.(*UnauthorizedHandler).Reason).To(Equal(reason))
			})
		})

		Describe("ServeHTTP", func() {
			w1 := httptest.NewRecorder()
			w2 := httptest.NewRecorder()
			util.Unauthorized(w1)
			h.ServeHTTP(w2, &http.Request{})
			It("Uses util.Unauthorized", func() {
				Expect(w1.Body).To(Equal(w2.Body))
				Expect(w1.Code).To(Equal(w2.Code))
			})
		})

	})

	// UnauthorizedHandler {{{

	// AuthenticationHandler }}}
	// AuthenticationHandler }}}

	// Authenticators {{{
	Describe("Authenticators", func() {
		It("Defines DefaultAuthenticator", func() {
			Expect(DefaultAuthenticator).ToNot(BeNil())
		})
	})
	// Authenticators }}}

	// AuthenticationHandler {{{
	Describe("AuthenticationHandler", func() {
		a := user.New()
		a.SetId(bson.NewObjectId())
		authed := true
		var err error = nil

		var authenticator auth.RequestAuthenticator = func(r *http.Request) (data.Agent, bool, error) {
			return a, authed, err
		}

		n1 := NewNullHandler()

		var errHandlerC = func(e error) http.Handler {
			return n1
		}

		n2 := NewNullHandler()

		var unauthHandlerC = func(r string) http.Handler {
			return n2
		}

		transferCalledCount := 0

		var transferFunc AuthenticatedHandlerFunc = func(w http.ResponseWriter, r *http.Request, a data.Agent) {
			transferCalledCount = transferCalledCount + 1

		}

		h := NewAuthenticationHandler(authenticator, errHandlerC, unauthHandlerC, transferFunc)

		Describe("NewAuthenticationHandler", func() {
			It("Allocates and returns a new AuthenticationHandler", func() {
				Expect(h).NotTo(BeNil())
				Expect(h).To(BeAssignableToTypeOf(&AuthenticationHandler{}))
			})

			It("Sets fields", func() {
				By("sets Authenticator")
				h := h.(*AuthenticationHandler)
				Expect(h.Authenticator).NotTo(BeNil())
				Expect(h.NewErrorHandler).NotTo(BeNil())
				Expect(h.NewUnauthorizedHandler).NotTo(BeNil())
				Expect(h.TransferFunc).NotTo(BeNil())
			})
		})

		Describe("ServeHTTP", func() {
			Context("Authentication Error", func() {
				err = errors.New("This is an error during the authentication process")
			})

			Context("Authentication Failure", func() {
			})

			Context("Authentication Successful", func() {
			})
		})

	})
	// AuthenticationHandler }}}
})

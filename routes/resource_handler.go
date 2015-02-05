package routes

import (
	"net/http"

	"github.com/elos/data"
	"github.com/elos/server/util/auth"
	"github.com/elos/transfer"
)

// Authenticte

type Handler interface {
	Handle(http.ResponseWriter, *http.Request, data.Identifiable, data.Store)
}

type AuthHandler struct {
	data.Store
	Handler
}

func Authenticated(h Handler, s data.Store) *AuthHandler {
	return &AuthHandler{
		Handler: h,
		Store:   s,
	}
}

func (h *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	agent, authenticated, err := auth.AuthenticateRequest(h.Store, r)
	if err != nil {
		logf("An error occurred during authentication, err: %s", err)
		// h.NewErrorHandler(err).ServeHTTP(w, r)
		return
	}

	if authenticated {
		h.Handler.Handle(w, r, agent, h.Store)
		logf("Agent with id %s authenticated", agent.ID())
	} else {
		logf("Agent with id %s authenticated", agent.ID())
		//	h.NewUnauthorizedHandler("Not authenticated").ServeHTTP(w, r)
	}
}

type PostHandler struct {
	Kind       data.Kind
	FormValues []string
}

func (h *PostHandler) Handle(w http.ResponseWriter, r *http.Request, a data.Identifiable, s data.Store) {
	var attrs data.AttrMap

	for _, k := range h.FormValues {
		attrs[k] = r.FormValue(k)
	}

	c := transfer.NewHTTPConnection(w, r, a)
	e := transfer.New(c, transfer.POST, h.Kind, attrs)
	go transfer.Route(e, s)
	log("we are success")
}

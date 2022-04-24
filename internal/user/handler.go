package user

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rest-api-server/internal/handlers"
	"rest-api-server/pkg/logger"
)

const (
	userURL  = "/users/:uuid"
	usersURL = "/users/"
)

type handler struct {
	logger *logger.Logger
}

func NewHandler(logger *logger.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetUsersList)
	router.GET(userURL, h.GetUserByUUID)
	router.POST(usersURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartialUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}

func (h *handler) GetUsersList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is GetUsersList"))
}
func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is GetUserByUUID"))
}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is CreateUser"))
}
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is UpdateUser"))
}
func (h *handler) PartialUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is PartialUpdateUser"))
}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is DeleteUser"))
}

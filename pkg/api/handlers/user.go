package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ZdravkoGyurov/go-grader/pkg/api"
	"github.com/ZdravkoGyurov/go-grader/pkg/api/response"
	"github.com/ZdravkoGyurov/go-grader/pkg/api/router/paths"
	"github.com/ZdravkoGyurov/go-grader/pkg/controller"
	"github.com/ZdravkoGyurov/go-grader/pkg/errors"
	"github.com/ZdravkoGyurov/go-grader/pkg/log"
	"github.com/ZdravkoGyurov/go-grader/pkg/model"
	"github.com/gorilla/mux"
)

type User struct {
	Controller *controller.Controller
}

func (h *User) GetAll(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	body := struct {
		CourseID string `json:"course_id"`
	}{}
	if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
		err = errors.Wrap(err, "failed to decode course_id from request body")
		response.SendError(writer, http.StatusBadRequest, err)
		return
	}

	users, err := h.Controller.GetAllUsers(ctx, body.CourseID)
	if err != nil {
		response.SendError(writer, api.StatusCode(err), err)
		return
	}

	response.SendData(writer, http.StatusOK, users)
}

func (h *User) Patch(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	userID, ok := mux.Vars(request)[paths.IDParam]
	if !ok {
		err := errors.New("failed to get user id path parameter")
		response.SendError(writer, http.StatusInternalServerError, err)
		return
	}

	var updateUser model.User
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&updateUser); err != nil {
		err = errors.Wrap(err, "failed to decode user from request body")
		response.SendError(writer, http.StatusBadRequest, err)
		return
	}
	updateUser.ID = ""

	updatedUser, err := h.Controller.UpdateUser(ctx, userID, &updateUser)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal user json data")
		response.SendError(writer, api.StatusCode(err), err)
		return
	}

	log.Info().Printf("updated user with id %s\n", userID)
	response.SendData(writer, http.StatusOK, updatedUser)
}

func (h *User) Delete(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	userID, ok := mux.Vars(request)[paths.IDParam]
	if !ok {
		err := errors.New("failed to get user id path parameter")
		response.SendError(writer, http.StatusInternalServerError, err)
		return
	}

	if err := h.Controller.DeleteUser(ctx, userID); err != nil {
		response.SendError(writer, api.StatusCode(err), err)
		return
	}

	log.Info().Printf("deleted user with id %s\n", userID)
	response.SendData(writer, http.StatusNoContent, struct{}{})
}

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ZdravkoGyurov/go-grader/pkg/api"
	"github.com/ZdravkoGyurov/go-grader/pkg/api/response"
	"github.com/ZdravkoGyurov/go-grader/pkg/api/router/paths"
	"github.com/ZdravkoGyurov/go-grader/pkg/controller"
	"github.com/ZdravkoGyurov/go-grader/pkg/errors"
	"github.com/ZdravkoGyurov/go-grader/pkg/log"
	"github.com/ZdravkoGyurov/go-grader/pkg/model"
)

type Course struct {
	Controller *controller.Controller
}

func (h *Course) Post(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var course model.Course
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&course); err != nil {
		err = errors.Wrap(err, "failed to decode course from request body")
		response.SendError(writer, http.StatusBadRequest, err)
		return
	}

	if err := h.Controller.CreateCourse(ctx, &course); err != nil {
		response.SendError(writer, api.StatusCode(err), err)
		return
	}

	log.Info().Printf("created course with id %s\n", course.ID)
	response.SendData(writer, http.StatusOK, course)
}

func (h *Course) Get(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	courseID, ok := mux.Vars(request)[paths.IDParam]
	if !ok {
		err := errors.New("failed to get course id path parameter")
		response.SendError(writer, http.StatusInternalServerError, err)
		return
	}

	course, err := h.Controller.GetCourse(ctx, courseID)
	if err != nil {
		response.SendError(writer, api.StatusCode(err), err)
		return
	}

	response.SendData(writer, http.StatusOK, course)
}

func (h *Course) GetAll(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	courses, err := h.Controller.GetAllCourses(ctx)
	if err != nil {
		response.SendError(writer, api.StatusCode(err), err)
		return
	}

	response.SendData(writer, http.StatusOK, courses)
}

func (h *Course) Patch(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	courseID, ok := mux.Vars(request)[paths.IDParam]
	if !ok {
		err := errors.New("failed to get course id path parameter")
		response.SendError(writer, http.StatusInternalServerError, err)
		return
	}

	var course model.Course
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&course); err != nil {
		err = errors.Wrap(err, "failed to decode course from request body")
		response.SendError(writer, http.StatusBadRequest, err)
		return
	}

	updatedCourse, err := h.Controller.UpdateCourse(ctx, courseID, &course)
	if err != nil {
		response.SendError(writer, api.StatusCode(err), err)
		return
	}

	response.SendData(writer, http.StatusOK, updatedCourse)
}

func (h *Course) Delete(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	courseID, ok := mux.Vars(request)[paths.IDParam]
	if !ok {
		err := errors.New("failed to get course id path parameter")
		response.SendError(writer, http.StatusInternalServerError, err)
		return
	}

	if err := h.Controller.DeleteCourse(ctx, courseID); err != nil {
		response.SendError(writer, api.StatusCode(err), err)
		return
	}

	log.Info().Printf("deleted course with id %s\n", courseID)
	response.SendData(writer, http.StatusNoContent, struct{}{})
}

package tehub

import (
	"github.com/tenlavien/spec/nut"
	"log"
	"net/http"
	"time"
)

type stepHandler struct {
	app       *App
}

func newStepHandler(app *App) *stepHandler {
	return &stepHandler{
		app:       app,
	}
}

func (h *stepHandler) handleCreateStep(w http.ResponseWriter, r *http.Request) {
	api := apiCreateStep

	err := nut.ValidateRoute(api, r.Method, r.URL.Path)
	if err != nil {
		err = nut.WriteResponse(w, http.StatusNotFound, nut.Map{"message": "invalid route", "error": err.Error()})
		if err != nil {
			log.Println(err)
		}
		return
	}

	params, err := nut.ParseParams(r, nut.ParamBody)
	if err != nil {
		err = nut.WriteResponse(w, http.StatusBadRequest, nut.Map{"message": "error parsing body params"})
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = nut.ValidateParams(api, params, []string{"step_code", "status"})
	if err != nil {
		err = nut.WriteResponse(w, http.StatusBadRequest, nut.Map{"message": "invalid parameter", "error": err.Error()})
		if err != nil {
			log.Println(err)
		}
		return
	}

	stepCode, _ := params.GetString("step_code")
	caseID, _ := params.GetInt64("case_id")
	status, _ := params.GetString("status")
	description, _ := params.GetString("description")
	stepIn, _ := params.GetString("step_in")
	stepOut, _ := params.GetString("step_out")

	step := &DBStep{
		StepCode:    stepCode,
		CaseID:      caseID,
		Description: description,
		Status:      nut.Status(status),
		StepIn:      stepIn,
		StepOut:     stepOut,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = h.app.Store.CreateStep(step)
	if err != nil {
		err = nut.WriteResponse(w, http.StatusInternalServerError, nut.Map{"message": "create step error", "error": err.Error()})
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = nut.WriteResponse(w, http.StatusOK, nut.Map{"message": "create step success", "id": step.ID})
	if err != nil {
		log.Println(err)
	}
}

func (h *stepHandler) handleUpdateStep(w http.ResponseWriter, r *http.Request) {
	api :=  apiUpdateStep

	err := nut.ValidateRoute(api, r.Method, r.URL.Path)
	if err != nil {
		err = nut.WriteResponse(w, http.StatusNotFound, nut.Map{"message": "invalid route", "error": err.Error()})
		if err != nil {
			log.Println(err)
		}
		return
	}

	params, err := nut.ParseParams(r, nut.ParamBody)
	if err != nil {
		err = nut.WriteResponse(w, http.StatusBadRequest, nut.Map{"message": "error parsing body params"})
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = nut.ValidateParams(api, params, []string{"id"})
	if err != nil {
		err = nut.WriteResponse(w, http.StatusBadRequest, nut.Map{"message": "invalid parameter", "error": err.Error()})
		if err != nil {
			log.Println(err)
		}
		return
	}

	id, _ := params.GetInt64("id")
	status, _ := params.GetString("status")
	description, _ := params.GetString("description")
	stepIn, _ := params.GetString("step_in")
	stepOut, _ := params.GetString("step_out")
	elapsedMilliSeconds, _ := params.GetInt64("elapsed_milliseconds")

	step := &DBStep{
		ID:                  id,
		Description:         description,
		Status:              nut.Status(status),
		StepIn:              stepIn,
		StepOut:             stepOut,
		UpdatedAt:           time.Now(),
		ElapsedMilliSeconds: elapsedMilliSeconds,
	}

	err = h.app.Store.UpdateStep(step)
	if err != nil {
		err = nut.WriteResponse(w, http.StatusInternalServerError, nut.Map{"message": "update step error", "error": err.Error()})
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = nut.WriteResponse(w, http.StatusOK, nut.Map{"message": "update step success"})
	if err != nil {
		log.Println(err)
	}
}

func (h *stepHandler) handleListSteps(w http.ResponseWriter, r *http.Request) {
	api := apiListStep

	err := nut.ValidateRoute(api, r.Method, r.URL.Path)
	if err != nil {
		err = nut.WriteResponse(w, http.StatusNotFound, nut.Map{"message": "invalid route", "error": err.Error()})
		if err != nil {
			log.Println(err)
		}
		return
	}

	params, err := nut.ParseParams(r, nut.ParamBody)
	if err != nil {
		err = nut.WriteResponse(w, http.StatusBadRequest, nut.Map{"message": "error parsing body params"})
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = nut.ValidateParams(api, params, []string{"page_id", "per_page"})
	if err != nil {
		err = nut.WriteResponse(w, http.StatusBadRequest, nut.Map{"message": "invalid parameter", "error": err.Error()})
		if err != nil {
			log.Println(err)
		}
		return
	}

	pageID, _ := params.GetInt64("page_id")
	perPage, _ := params.GetInt64("per_page")
	caseID, _ := params.GetInt64("case_id")
	runID, _ := params.GetString("run_id")

	findStep := &DBStep{
		CaseID: caseID,
		RunID:  runID,
	}

	steps, err := h.app.Store.ListSteps(findStep, pageID, perPage)
	if err != nil {
		err = nut.WriteResponse(w, http.StatusInternalServerError, nut.Map{"message": "list steps error", "error": err.Error()})
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = nut.WriteResponse(w, http.StatusOK, nut.Map{"message": "list step success", "steps": steps})
	if err != nil {
		log.Println(err)
	}
}

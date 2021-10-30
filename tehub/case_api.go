package tehub

import (
	"github.com/tenlavien/spec/nut"
	"log"
	"net/http"
	"time"
)

type caseHandler struct {
	app       *App
}

func newCaseHandler(app *App) *caseHandler {
	return &caseHandler{
		app:       app,
	}
}

func (h *caseHandler) handleCreateCase(w http.ResponseWriter, r *http.Request) {
	api := apiCreateCase

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

	err = nut.ValidateParams(api, params, []string{"case_code", "status"})
	if err != nil {
		err = nut.WriteResponse(w, http.StatusBadRequest, nut.Map{"message": "invalid parameter", "error": err.Error()})
		if err != nil {
			log.Println(err)
		}
		return
	}

	caseCode, _ := params.GetString("case_code")
	status, _ := params.GetString("status")
	description, _ := params.GetString("description")
	caseIn, _ := params.GetString("case_in")
	caseOut, _ := params.GetString("case_out")

	tc := &DBCase{
		CaseCode:    caseCode,
		Description: description,
		Status:      nut.Status(status),
		CaseIn:      caseIn,
		CaseOut:     caseOut,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = h.app.Store.CreateCase(tc)
	if err != nil {
		err = nut.WriteResponse(w, http.StatusInternalServerError, nut.Map{"message": "create test case error", "error": err.Error()})
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = nut.WriteResponse(w, http.StatusOK, nut.Map{"message": "create test case success", "id": tc.ID})
	if err != nil {
		log.Println(err)
	}
}

func (h *caseHandler) handleUpdateCase(w http.ResponseWriter, r *http.Request) {
	api := apiUpdateCase

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
	caseIn, _ := params.GetString("case_in")
	caseOut, _ := params.GetString("case_out")

	tc := &DBCase{
		ID:          id,
		Description: description,
		Status:      nut.Status(status),
		CaseIn:      caseIn,
		CaseOut:     caseOut,
		UpdatedAt:   time.Now(),
	}

	err = h.app.Store.UpdateCase(tc)
	if err != nil {
		err = nut.WriteResponse(w, http.StatusInternalServerError, nut.Map{"message": "update test case error", "error": err.Error()})
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = nut.WriteResponse(w, http.StatusOK, nut.Map{"message": "update test case success"})
	if err != nil {
		log.Println(err)
	}
}

func (h *caseHandler) handleListCases(w http.ResponseWriter, r *http.Request) {
	api := apiListCases

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
	runID, _ := params.GetString("run_id")

	findTC := &DBCase{
		RunID:  runID,
	}

	testCases, err := h.app.Store.ListCases(findTC, pageID, perPage)
	if err != nil {
		err = nut.WriteResponse(w, http.StatusInternalServerError, nut.Map{"message": "list test cases error", "error": err.Error()})
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = nut.WriteResponse(w, http.StatusOK, nut.Map{"message": "list cases success", "test_cases": testCases})
	if err != nil {
		log.Println(err)
	}
}

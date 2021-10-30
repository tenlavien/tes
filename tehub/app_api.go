package tehub

import (
	"github.com/tenlavien/spec/nut"
	"log"
	"net/http"
)

var (
	apiPing = nut.API{
		Path: "/ping",
	}

	apiCleanHub = nut.API{
		Method:      http.MethodDelete,
		Path:        "/hub/delete",
	}

	apiCreateStep = nut.API{
		Method:    http.MethodPost,
		Path:      "/steps/create",
		ParamFrom: nut.ParamBody,
		Parameters: map[string]nut.Parameter{
			"step_code":   {""},
			"case_id":     {""},
			"description": {""},
			"status":      {""},
			"step_in":     {""},
			"step_out":    {""},
			"created_at":  {""},
			"updated_at":  {""},
		},
	}

	apiUpdateStep = nut.API{
		Method:    http.MethodPatch,
		Path:      "/steps/update",
		ParamFrom: nut.ParamBody,
		Parameters: map[string]nut.Parameter{
			"id":                   {""},
			"description":          {""},
			"status":               {""},
			"step_in":              {""},
			"step_out":             {""},
			"elapsed_milliseconds": {""},
			"updated_at":           {""},
		},
	}

	apiListStep = nut.API{
		Method:    http.MethodPost,
		Path:      "/steps/list",
		ParamFrom: nut.ParamBody,
		Parameters: map[string]nut.Parameter{
			"page_id":  {""},
			"per_page": {""},
			"run_id":   {""},
			"case_id":  {""},
		},
	}

	apiCreateCase = nut.API{
		Method:    http.MethodPost,
		Path:      "/cases/create",
		ParamFrom: nut.ParamBody,
		Parameters: map[string]nut.Parameter{
			"case_code":   {""},
			"description": {""},
			"status":      {""},
			"case_in":     {""},
			"case_out":    {""},
			"created_at":  {""},
			"updated_at":  {""},
		},
	}

	apiUpdateCase = nut.API{
		Method:    http.MethodPatch,
		Path:      "/cases/update",
		ParamFrom: nut.ParamBody,
		Parameters: map[string]nut.Parameter{
			"id":          {""},
			"description": {""},
			"status":      {""},
			"case_in":     {""},
			"case_out":    {""},
			"updated_at":  {""},
		},
	}

	apiListCases = nut.API{
		Method:    http.MethodPost,
		Path:      "/cases/list",
		ParamFrom: nut.ParamBody,
		Parameters: map[string]nut.Parameter{
			"page_id":  {""},
			"per_page": {""},
			"run_id":   {""},
		},
	}

)

func (app *App) InitAPIList() {
	mux := http.NewServeMux()

	stepHandler := newStepHandler(app)
	caseHandler := newCaseHandler(app)

	mux.HandleFunc(apiPing.Path, app.handlePing)
	mux.HandleFunc(apiCleanHub.Path, app.handleCleanHub)

	mux.HandleFunc(apiCreateStep.Path, stepHandler.handleCreateStep)
	mux.HandleFunc(apiUpdateStep.Path, stepHandler.handleUpdateStep)
	mux.HandleFunc(apiListStep.Path, stepHandler.handleListSteps)

	mux.HandleFunc(apiCreateCase.Path, caseHandler.handleCreateCase)
	mux.HandleFunc(apiUpdateCase.Path, caseHandler.handleUpdateCase)
	mux.HandleFunc(apiListCases.Path, caseHandler.handleListCases)

	app.Server.Handler = mux
}

func (app *App) handlePing(w http.ResponseWriter, r *http.Request) {
	err := nut.WriteResponse(w, http.StatusOK, nut.Map{"message": "pong"})
	if err != nil {
		log.Println(err)
	}
}

func (app *App) handleCleanHub(w http.ResponseWriter, r *http.Request) {
	err := nut.ValidateRoute(apiCleanHub, r.Method, r.URL.Path)
	if err != nil {
		err = nut.WriteResponse(w, http.StatusNotFound, nut.Map{"message": "invalid route", "error": err.Error()})
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = app.Store.TruncateSteps()
	if err != nil {
		err = nut.WriteResponse(w, http.StatusInternalServerError, nut.Map{"message": "error truncate steps", "error": err.Error()})
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = app.Store.TruncateCases()
	if err != nil {
		err = nut.WriteResponse(w, http.StatusInternalServerError, nut.Map{"message": "error truncate test cases", "error": err.Error()})
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = nut.WriteResponse(w, http.StatusOK, nut.Map{"message": "clean hub successfully"})
	if err != nil {
		log.Println(err)
	}
}

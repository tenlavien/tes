package nut

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"log"
	"net/http"
	"time"
)

var DefaultHTTPClient = http.Client{
	Timeout:   60 * time.Second,
	Transport: &http.Transport{},
}

type Parameter struct {
	Tag string // validator v10 tags
}

func ParseParams(r *http.Request, parseFrom ParamType) (Map, error) {
	buf := bytes.NewBuffer(make([]byte, 0))

	var reader io.Reader
	params := Map{}

	switch parseFrom {
	case ParamBody:
		reader = io.TeeReader(r.Body, buf)
		decoder := json.NewDecoder(reader)
		err := decoder.Decode(&params)
		if err != nil {
			if err != io.EOF {
				log.Println("[error] error read body")
			}
			return nil, err
		}
	case ParamQuery:
		for key, val := range r.URL.Query() {
			if len(val) > 0 {
				params[key] = val[0]
			}
		}
	default:
		return nil, fmt.Errorf("unsupported parsing params from %s", parseFrom)
	}

	return params, nil
}

func WriteResponse(w http.ResponseWriter, statusCode int, resBody Map) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	res, err := json.Marshal(resBody)
	if err != nil {
		return err
	}
	_, err = w.Write(res)
	return err
}

type API struct {
	Method     string
	Path       string
	ParamFrom  ParamType
	Parameters map[string]Parameter
	RequireAuth bool
}

func ValidateParams(a API, inputParams Map, requiredParams []string) error {
	validate := validator.New()
	for _, p := range requiredParams {
		err := validate.Var(inputParams[p], "required")
		if err != nil {
			return err
		}
	}

	for key, val := range inputParams {

		// any regex
		validatorTag := a.Parameters[key].Tag
		if validatorTag == "" {
			continue
		}

		if err := validate.Var(val, validatorTag); err != nil {
			return err
		}
	}

	return nil
}

func ValidateRoute(presetAPI API, requestedMethod string, requestedPath string) error {
	if presetAPI.Method != requestedMethod  || presetAPI.Path != requestedPath {
		return fmt.Errorf("invalid method or request route not found")
	}
	log.Println("[http] handling request", requestedMethod, requestedPath)
	return nil
}

func ParseResponseBody(body io.Reader) (Map, error) {
	var m Map
	if body == nil {
		return m, nil
	}

	decoder := json.NewDecoder(body)
	err := decoder.Decode(&m)
	if err != nil {
		if err == io.EOF {
			return m, nil
		}

		log.Println("[error] error parse response body")
		return nil, err
	}
	return m, nil
}
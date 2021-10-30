package tehub

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/tenlavien/spec/nut"
	"net/http"
)

type Client struct {
	Root       string
	StepCreate nut.API
	StepUpdate nut.API
	CaseCreate nut.API
	CaseUpdate nut.API
}

func NewClient(root string) *Client {
	return &Client{
		Root:       root,
		StepCreate: nut.API{Method: http.MethodPost, Path: "/steps/create"},
		StepUpdate: nut.API{Method: http.MethodPatch, Path: "/steps/update"},
		CaseCreate: nut.API{Method: http.MethodPost, Path: "/cases/create"},
		CaseUpdate: nut.API{Method: http.MethodPatch, Path: "/cases/update"},
	}
}

func (hub *Client) RequestCreateStep(step *DBStep) (int64, *http.Response, error) {
	params, err := json.Marshal(step)
	if err != nil {
		return 0, nil, err
	}
	body := bytes.NewReader(params)
	req, err := http.NewRequestWithContext(context.Background(), hub.StepCreate.Method, hub.Root+hub.StepCreate.Path, body)
	if err != nil {
		return 0, nil, err
	}
	client := nut.DefaultHTTPClient
	res, err := client.Do(req)
	if err != nil {
		return 0, res, err
	}
	if res.StatusCode != http.StatusOK {
		return 0, res, fmt.Errorf("request error %s", res.Status)
	}
	m, err := nut.ParseResponseBody(res.Body)
	if err != nil {
		return 0, res, err
	}
	id, _ := m.GetInt64("id")
	return id, res, nil
}

func (hub *Client) RequestUpdateStep(step *DBStep) (*http.Response, error) {
	params, err := json.Marshal(step)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(params)
	req, err := http.NewRequestWithContext(context.Background(), hub.StepUpdate.Method, hub.Root+hub.StepUpdate.Path, body)
	if err != nil {
		return nil, err
	}
	client := nut.DefaultHTTPClient
	res, err := client.Do(req)
	if err != nil {
		return res, err
	}
	if res.StatusCode != http.StatusOK {
		return res, fmt.Errorf("request error %s", res.Status)
	}
	return res, nil
}

func (hub *Client) RequestCreateTestCase(tc *DBCase) (int64, *http.Response, error) {
	params, err := json.Marshal(tc)
	if err != nil {
		return 0, nil, err
	}
	body := bytes.NewReader(params)
	req, err := http.NewRequestWithContext(context.Background(), hub.CaseCreate.Method, hub.Root+hub.CaseCreate.Path, body)
	if err != nil {
		return 0, nil, err
	}
	client := nut.DefaultHTTPClient
	res, err := client.Do(req)
	if err != nil {
		return 0, res, err
	}
	if res.StatusCode != http.StatusOK {
		return 0, res, fmt.Errorf("request error %s", res.Status)
	}
	m, err := nut.ParseResponseBody(res.Body)
	if err != nil {
		return 0, res, err
	}
	id, _ := m.GetInt64("id")
	return id, res, nil
}

func (hub *Client) RequestUpdateTestCase(tc *DBCase) (*http.Response, error) {
	params, err := json.Marshal(tc)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(params)
	req, err := http.NewRequestWithContext(context.Background(), hub.CaseUpdate.Method, hub.Root+hub.CaseUpdate.Path, body)
	if err != nil {
		return nil, err
	}
	client := nut.DefaultHTTPClient
	res, err := client.Do(req)
	if err != nil {
		return res, err
	}
	if res.StatusCode != http.StatusOK {
		return res, fmt.Errorf("request error %s", res.Status)
	}
	return res, nil
}

package teman
//
//import (
//	"context"
//	"github.com/tenlavien/spec/nut"
//	"github.com/tenlavien/spec/tehub"
//	"net/http"
//)
//
//type shStep struct {
//	method  string
//	url     string
//	headers map[string]string
//	params  interface{}
//	out     *http.Response
//
//	context context.Context
//	client  http.Client
//	dbStep  *tehub.DBStep
//	hub     *HubClient
//}
//
//func NewStepSH(hub *HubClient, stepCode, description string) *httpStep {
//	return &httpStep{
//		context: context.Background(),
//		client:  defaultHTTPClient,
//		hub:     hub,
//		dbStep: &tehub.DBStep{
//			StepCode:        stepCode,
//			Description: description,
//			Status:      nut.Pending,
//		},
//	}
//}
//
//func (s *httpStep) Test(t *testing.T) {
//	params, err := json.Marshal(s.params)
//	assert.NoError(t, err)
//
//	body := bytes.NewReader(params)
//	req, err := http.NewRequestWithContext(s.context, s.method, s.url, body)
//	assert.NoError(t, err)
//	for key, val := range s.headers {
//		req.Header.Add(key, val)
//	}
//
//	in, err := s.StepIn(req)
//	assert.NoError(t, err)
//
//	res, err := s.client.Do(req)
//	assert.NoError(t, err)
//
//	out, err := s.StepOut(res)
//	assert.NoError(t, err)
//
//	//s.Update()
//	s.out = res
//}
//
//func (s *httpStep) WithURL(url string) *httpStep {
//	s.url = url
//	return s
//}
//
//func (s *httpStep) WithMethod(method string) *httpStep {
//	s.method = method
//	return s
//}
//
//func (s *httpStep) WithHeaders(headers map[string]string) *httpStep {
//	s.headers = headers
//	return s
//}
//
//func (s *httpStep) WithParams(params interface{}) *httpStep {
//	s.params = params
//	return s
//}
//
//func (s *httpStep) WithContext(ctx context.Context) *httpStep {
//	s.context = ctx
//	return s
//}
//
//func (s *httpStep) WithHTTPClient(client http.Client) *httpStep {
//	s.client = client
//	return s
//}
//
//func (s *httpStep) StepIn(req *http.Request) (string, error) {
//	cmd, err := http2curl.GetCurlCommand(req)
//	if err != nil {
//		return "", err
//	}
//	return cmd.String(), nil
//}
//
//func (s *httpStep) StepOut(res *http.Response) (string, error) {
//	return "", nil
//}
//
//func (s *httpStep) Out() interface{} {
//	return s.out
//}
//
//func (s *httpStep) Seed() (Step, error) {
//	id, _, err := s.hub.RequestCreateStep(s.dbStep)
//	if err != nil {
//		return nil, err
//	}
//	s.dbStep.ID = id
//	return s, nil
//}
//
//func (s *httpStep) Update() error {
//	return nil
//}

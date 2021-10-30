package teman

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"github.com/tenlavien/spec/nut"
	"github.com/tenlavien/spec/tehub"
	"io"
	"moul.io/http2curl"
	"net/http"
	"net/url"
	"testing"
	"time"
)

type httpStep struct {
	TStep

	inReq     *http.Request
	inPayload interface{}
	outRes    *http.Response

	context context.Context
	client  http.Client
}

func NewStepHTTP(hub *tehub.Client, stepCode, description string) *httpStep {
	return &httpStep{
		TStep:     NewStep(hub, stepCode, description),
		inReq:     &http.Request{},
		inPayload: nut.Map{},
		outRes:    &http.Response{},
		context:   context.Background(),
		client:    nut.DefaultHTTPClient,
	}
}

func (s *httpStep) Out() interface{} {
	return s.outRes
}

func (s *httpStep) DOFailed() bool {
	return s.outRes.StatusCode != http.StatusOK
}

func (s *httpStep) DOPassed() bool {
	return s.outRes.StatusCode == http.StatusOK
}

func (s *httpStep) Test(t *testing.T) string {
	err := s.compileRequest()
	require.NoError(t, err)

	in, err := s.StepIn()
	require.NoError(t, err)
	s.DBStep().StepIn = in
	s.DBStep().Status = nut.Running

	err = s.Update(s.DBStep())
	require.NoError(t, err)

	startTime := time.Now()

	res, err := s.client.Do(s.inReq)
	if err != nil {
		s.DBStep().StepOut = err.Error()
		s.DBStep().Status = nut.Failed
		err := s.Update(s.DBStep())
		require.NoError(t, err)
	}
	s.outRes = res
	s.DBStep().ElapsedMilliSeconds = time.Since(startTime).Milliseconds()

	if err != nil {
		return string(s.DBStep().Status)
	}

	if s.DOPassed() {
		s.DBStep().Status = nut.Passed
	}
	if s.DOFailed() {
		s.DBStep().Status = nut.Failed
	}

	out, err := s.StepOut()
	require.NoError(t, err)
	s.DBStep().StepOut = out

	err = s.Update(s.DBStep())
	require.NoError(t, err)

	return string(s.DBStep().Status)
}

type reqReader struct {
	io.Reader
}

func (reqReader) Close() error {
	return nil
}

func (s *httpStep) compileRequest() error {
	params, err := json.Marshal(s.inPayload)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(s.context, s.inReq.Method, s.inReq.URL.RawPath, reqReader{Reader: bytes.NewReader(params)})
	if err != nil {
		return err
	}
	s.inReq = req
	return nil
}

func (s *httpStep) WithRunID(runID string) *httpStep {
	s.DBStep().RunID = runID
	return s
}

func (s *httpStep) WithURL(requestedURL string) *httpStep {
	s.inReq.URL = &url.URL{
		RawPath: requestedURL,
	}
	return s
}

func (s *httpStep) WithHTTPClient(client http.Client) *httpStep {
	s.client = client
	return s
}

func (s *httpStep) WithMethod(method string) *httpStep {
	s.inReq.Method = method
	return s
}

func (s *httpStep) WithHeaders(headers map[string]string) *httpStep {
	for key, value := range headers {
		s.inReq.Header.Add(key, value)
	}
	return s
}

func (s *httpStep) WithPayLoad(payload interface{}) *httpStep {
	s.inPayload = payload
	return s
}

func (s *httpStep) WithContext(ctx context.Context) *httpStep {
	s.context = ctx
	return s
}

func (s *httpStep) StepIn() (string, error) {
	cmd, err := http2curl.GetCurlCommand(s.inReq)
	if err != nil {
		return "", err
	}
	return cmd.String(), nil
}

func (s *httpStep) StepOut() (string, error) {
	m, err := nut.ParseResponseBody(s.outRes.Body)
	if err != nil {
		return "", err
	}
	resBody, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return s.outRes.Status + " " + string(resBody), nil
}

package teman

import (
	"github.com/stretchr/testify/require"
	"github.com/tenlavien/spec/nut"
	"github.com/tenlavien/spec/tehub"
	"net/http"
	"testing"
)

func TestSteps(t *testing.T) {
	hub := tehub.NewClient("http://0.0.0.0:8801")
	require.NotNil(t, hub)

	step1 := NewStepHTTP(hub, "step_1", "do step 1").
		WithMethod(http.MethodGet).
		WithURL("http://testing-telco-5.tsengineering.io/score_api/ping")

	err := step1.Seed()
	require.NoError(t, err)

	status := step1.Test(t)
	require.EqualValues(t, "passed", status)

	res := step1.Out().(*http.Response)
	require.NotNil(t, res)

	m, err := nut.ParseResponseBody(res.Body)
	require.NoError(t, err)
	m.HasKeyValue("verdict", "success")
}


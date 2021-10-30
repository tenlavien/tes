package teman

import (
	"github.com/stretchr/testify/require"
	"github.com/tenlavien/spec/tehub"
	"net/http"
	"testing"
)

func TestTCase(t *testing.T) {
	hub := tehub.NewClient("http://0.0.0.0:8801")
	require.NotNil(t, hub)

	step1 := NewStepHTTP(hub, "step1", "do step 1").
		WithMethod(http.MethodGet).
		WithURL("http://testing-telco-5.tsengineering.io/score_api/ping")

	step2 := NewStepHTTP(hub, "step2", "do step 2").
		WithMethod(http.MethodGet).
		WithURL("http://testing-telco-5.tsengineering.io/score_api/do")

	tc := NewTCase(hub, "tc", "test something").
		AddSteps(step1, step2)

	err := tc.Seed()
	require.NoError(t, err)

	tc.Test(t)
}
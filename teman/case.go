package teman

import (
	"github.com/stretchr/testify/require"
	"github.com/tenlavien/spec/nut"
	"github.com/tenlavien/spec/tehub"
	"testing"
	"time"
)

type TCase interface {
	AddSteps(step ...TStep) TCase
	Parallel() TCase
	Seed() error
	Update(tc *tehub.DBCase) error
	DBCase() *tehub.DBCase
	TestIn() (string, error)
	TestOut() (string, error)
	DOPassed() bool
	DOFailed() bool
	Test(t *testing.T) string
}

type tCase struct {
	isParallel bool
	dbCase     *tehub.DBCase
	steps      []TStep
	hub        *tehub.Client
}

func NewTCase(hub *tehub.Client, caseCode string, caseDescription string) TCase {
	return &tCase{
		hub:        hub,
		isParallel: false,
		steps:      make([]TStep, 0),
		dbCase: &tehub.DBCase{
			CaseCode:    caseCode,
			Description: caseDescription,
			Status:      nut.Pending,
		},
	}
}

func (tc *tCase) AddSteps(step ...TStep) TCase {
	tc.steps = append(tc.steps, step...)
	return tc
}

func (tc *tCase) Parallel() TCase {
	tc.isParallel = true
	return tc
}

func (tc *tCase) Test(t *testing.T) string {
	in, err := tc.TestIn()
	require.NoError(t, err)
	tc.dbCase.CaseIn = in
	tc.dbCase.Status = nut.Running

	err = tc.Update(tc.dbCase)
	require.NoError(t, err)

	startTime := time.Now()

	if tc.isParallel {
		tc.testInParallel(t)
	} else {
		tc.testInOrder(t)
	}

	tc.dbCase.ElapsedMilliSeconds = time.Since(startTime).Milliseconds()

	if tc.DOPassed() {
		tc.dbCase.Status = nut.Passed
	}

	if tc.DOFailed() {
		tc.dbCase.Status = nut.Failed
	}

	out, err := tc.TestOut()
	require.NoError(t, err)
	tc.dbCase.CaseOut = out

	err = tc.Update(tc.dbCase)
	require.NoError(t, err)

	return string(tc.dbCase.Status)
}

func (tc *tCase) testInOrder(t *testing.T) {
	for _, step := range tc.steps {
		step.Test(t)
	}
}

func (tc *tCase) testInParallel(t *testing.T) {

}

func (tc *tCase) Seed() error {
	tcID, _, err := tc.hub.RequestCreateTestCase(tc.dbCase)
	if err != nil {
		return err
	}
	tc.dbCase.ID = tcID

	for _, step := range tc.steps {
		step.DBStep().CaseID = tcID
		err := step.Seed()
		if err != nil {
			return err
		}
	}
	return nil
}

func (tc *tCase) DBCase() *tehub.DBCase {
	return tc.dbCase
}

func (tc *tCase) TestIn() (string, error) {
	tcIn := nut.Map{}
	for _, step := range tc.steps {
		tcIn[step.DBStep().StepCode] = nut.Map{
			"description": step.DBStep().Description,
		}
	}
	in, err := tcIn.ToString()
	if err != nil {
		return "", err
	}
	return in, nil
}

func (tc *tCase) TestOut() (string, error) {
	tcOut := nut.Map{}
	for _, step := range tc.steps {
		tcOut[step.DBStep().StepCode] = nut.Map{
			"status": step.DBStep().Status,
		}
	}
	out, err := tcOut.ToString()
	if err != nil {
		return "", err
	}
	return out, nil
}

func (tc *tCase) Update(updatedTC *tehub.DBCase) error {
	_, err := tc.hub.RequestUpdateTestCase(updatedTC)
	return err
}

func (tc *tCase) DOPassed() bool {
	passedCount := 0
	for _, step := range tc.steps {
		if step.DOPassed() {
			passedCount = passedCount+1
		}
	}

	return len(tc.steps) > 0 && len(tc.steps) == passedCount
}

func (tc *tCase) DOFailed() bool {
	for _, step := range tc.steps {
		if step.DOFailed() {
			// one single step failed
			return true
		}
	}
	return false
}

package teman

import (
	"github.com/tenlavien/spec/nut"
	"github.com/tenlavien/spec/tehub"
	"testing"
)

type TStep interface {
	Test(t *testing.T) string
	Seed() error
	Out() interface{}
	DOFailed() bool
	DOPassed() bool
	DBStep() *tehub.DBStep
	Update(step *tehub.DBStep) error
	StepIn() (string, error)
	StepOut() (string, error)
}

type tStep struct {
	dbStep    *tehub.DBStep
	hubClient *tehub.Client
}

func NewStep(hub *tehub.Client, code string, description string) TStep {
	return &tStep{
		hubClient: hub,
		dbStep: &tehub.DBStep{
			StepCode:    code,
			Description: description,
			Status:      nut.Pending,
		},
	}
}

func (s *tStep) DBStep() *tehub.DBStep {
	return s.dbStep
}

func (s *tStep) DOPassed() bool {
	return false
}

func (s *tStep) DOFailed() bool {
	return false
}

func (s *tStep) Out() interface{} {
	return nil
}

func (s *tStep) Test(t *testing.T) string {
	t.Skip("no test to run")
	return string(nut.Skipped)
}

func (s *tStep) Seed() error {
	in, err := s.StepIn()
	if err != nil {
		return err
	}
	s.dbStep.StepIn = in

	id, _, err := s.hubClient.RequestCreateStep(s.dbStep)
	if err != nil {
		return err
	}
	s.dbStep.ID = id
	return nil
}

func (s *tStep) Update(step *tehub.DBStep) error {
	_, err := s.hubClient.RequestUpdateStep(step)
	return err
}

func (s *tStep) StepIn() (string, error) {
	return "", nil
}

func (s *tStep) StepOut() (string, error) {
	return "", nil
}

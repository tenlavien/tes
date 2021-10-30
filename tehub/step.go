package tehub

import (
	"github.com/tenlavien/spec/nut"
	"log"
	"time"
)

type DBStep struct {
	ID                  int64      `json:"id,omitempty" gorm:"column:id"`
	StepCode            string     `json:"step_code,omitempty" gorm:"column:step_code"`
	CaseID              int64      `json:"case_id,omitempty" gorm:"column:case_id"`
	RunID               string     `json:"run_id"`
	Description         string     `json:"description,omitempty" gorm:"column:description"`
	Status              nut.Status `json:"status,omitempty" gorm:"column:status"`
	StepIn              string     `json:"step_in,omitempty" gorm:"column:step_in"`
	StepOut             string     `json:"step_out,omitempty" gorm:"column:step_out"`
	ElapsedMilliSeconds int64      `json:"elapsed_milliseconds" gorm:"column:elapsed_milliseconds"`
	CreatedAt           time.Time  `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt           time.Time  `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (s *Store) CreateStep(step *DBStep) error {
	return s.DB.Table("steps").Create(step).Error
}

func (s *Store) UpdateStep(updatedStep *DBStep) error {
	currentStep := DBStep{
		ID: updatedStep.ID,
	}
	err := s.DB.Table("steps").First(&currentStep).Error
	if err != nil {
		log.Println("[error] record not found")
		return err
	}
	if updatedStep.Status != "" {
		currentStep.Status = updatedStep.Status
	}
	if updatedStep.StepIn != "" {
		currentStep.StepIn = updatedStep.StepIn
	}
	if updatedStep.StepOut != "" {
		currentStep.StepOut = updatedStep.StepOut
	}
	if updatedStep.ElapsedMilliSeconds > 0 {
		currentStep.ElapsedMilliSeconds = updatedStep.ElapsedMilliSeconds
	}
	if updatedStep.Description != "" {
		currentStep.Description = updatedStep.Description
	}
	if !updatedStep.UpdatedAt.IsZero() {
		currentStep.UpdatedAt = updatedStep.UpdatedAt
	}

	err = s.DB.Table("steps").Save(&currentStep).Error
	if err != nil {
		log.Println("[error] error saving update to db")
		return err
	}
	return nil
}

func (s *Store) ListSteps(step *DBStep, pageID, perPage int64) ([]DBStep, error) {
	var foundSteps []DBStep
	err := s.DB.Table("steps").Where(step).Scopes(Paginate(int(pageID), int(perPage))).Find(&foundSteps).Error
	return foundSteps, err
}

func (s *Store) TruncateSteps() error {
	return s.DB.Exec("truncate steps").Error
}

package tehub

import (
	"github.com/tenlavien/spec/nut"
	"log"
	"time"
)

type DBCase struct {
	ID                  int64      `json:"id" gorm:"column:id"`
	CaseCode            string     `json:"case_code" gorm:"column:case_code"`
	RunID               string     `json:"run_id"`
	Description         string     `json:"description" gorm:"column:description"`
	Status              nut.Status `json:"status" gorm:"column:status"`
	CaseIn              string     `json:"case_in" gorm:"column:case_in"`
	CaseOut             string     `json:"case_out" gorm:"column:case_out"`
	ElapsedMilliSeconds int64      `json:"elapsed_milliseconds" gorm:"column:elapsed_milliseconds"`
	CreatedAt           time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt           time.Time  `json:"updated_at" gorm:"column:updated_at"`
}

func (s *Store) CreateCase(tc *DBCase) error {
	return s.DB.Table("cases").Create(&tc).Error
}

func (s *Store) UpdateCase(updatedTC *DBCase) error {
	currentTC := DBCase{
		ID: updatedTC.ID,
	}
	err := s.DB.Table("cases").First(&currentTC).Error
	if err != nil {
		log.Println("[error] record not found")
		return err
	}
	if updatedTC.Status != "" {
		currentTC.Status = updatedTC.Status
	}
	if updatedTC.CaseIn != "" {
		currentTC.CaseIn = updatedTC.CaseIn
	}
	if updatedTC.CaseOut != "" {
		currentTC.CaseOut = updatedTC.CaseOut
	}
	if updatedTC.Description != "" {
		currentTC.Description = updatedTC.Description
	}
	if !updatedTC.UpdatedAt.IsZero() {
		currentTC.UpdatedAt = updatedTC.UpdatedAt
	}

	err = s.DB.Table("cases").Save(&currentTC).Error
	if err != nil {
		log.Println("[error] error saving update to db")
		return err
	}
	return nil
}

func (s *Store) ListCases(tc *DBCase, pageID, perPage int64)  ([]DBCase, error) {
	var foundCases []DBCase
	err := s.DB.Table("cases").Where(tc).Scopes(Paginate(int(pageID), int(perPage))).Find(&foundCases).Error
	return foundCases, err
}

func (s *Store) TruncateCases() error {
	return s.DB.Exec("truncate cases").Error
}

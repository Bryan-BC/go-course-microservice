package models

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/Bryan-BC/go-course-microservice/pkg/pb"
)

type Days []*pb.CourseIndex

type Course struct {
	Id          string `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Schedule    Days   `json:"schedule" gorm:"column:schedule;type:record"`
}

func (days *Days) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), days)
}

func (days Days) Value() (driver.Value, error) {
	return json.Marshal(days)
}

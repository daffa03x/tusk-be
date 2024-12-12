package models

import "time"

type Task struct {
	Id           int       `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	UserId       int       `gorm:"type:int" json:"user_id"`
	Title        string    `gorm:"type:varchar(255)" json:"title"`
	Description  string    `gorm:"type:text" json:"description"`
	Status       string    `gorm:"type:varchar(10)" json:"status"`
	Reason       string    `gorm:"type:text" json:"reason"`
	Revision     int8      `gorm:"type:int;default:0" json:"revision"`
	DueDate      time.Time `json:"due_date"`
	SubmitDate   time.Time `json:"submit_date"`
	RejectedDate time.Time `json:"rejected_date"`
	ApprovedDate time.Time `json:"approved_date"`
	Attachment   string    `gorm:"type:varchar(255)" json:"attachment"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	User         User      `gorm:"foreignKey:UserId" json:"user,omitempty"` // belongs to
}

package model

type Register struct {
	TaskName    string `json:"task_name" db:"task_name"`
	TaskOwner   string `json:"owner_name" db:"owner_name"`
	Description string `json:"description" db:"description"`
	Status      string `json:"status" db:"status"`
}

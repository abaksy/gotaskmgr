/*
Copyright © 2023 Aronya Baksy abaksy@gmail.com
*/
package db

import "time"

type Task struct {
	ID               uint64    `json:"id" storm:"increment"`
	Name             string    `json:"taskName"`
	CreationDate     time.Time `json:"creationDate"`
	CompletionDate   time.Time `json:"completionDate"`
	CompletionStatus int       `json:"completionStatus"`
}

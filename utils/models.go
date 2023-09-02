/*
Copyright © 2023 Aronya Baksy abaksy@gmail.com
*/
package cmd

import "time"

type Task struct {
	ID           uint64
	Name         string
	CreationDate time.Time
}

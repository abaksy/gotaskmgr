/*
Copyright © 2023 Aronya Baksy abaksy@gmail.com
*/
package db

import (
	"fmt"
	"os"
	"path"
	"time"

	constants "github.com/abaksy/gotaskmgr/const"
	"github.com/asdine/storm/v3"
)

func GetDBFilePath() string {
	userHome, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return path.Join(userHome, constants.DATABASE_DIR, constants.DATABASE_FILE)
}

// Initialize the BoltDB database and create a bucket under it
func InitDB() error {

	// Setup dir for DB file
	userHome, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	dbFileDir := path.Join(userHome, constants.DATABASE_DIR)
	if _, err := os.Stat(dbFileDir); err != nil {
		err := os.Mkdir(dbFileDir, 0700)
		if err != nil {
			return err
		}
	}
	dbFilePath := path.Join(dbFileDir, constants.DATABASE_FILE)

	// Create BoltDB database
	db, err := storm.Open(dbFilePath)
	if err != nil {
		return fmt.Errorf("error while opening DB: %v", err)
	}
	defer db.Close()
	return err
}

// Add a Task to the BoltDB bucket
func AddTask(task Task) error {
	dbFilePath := GetDBFilePath()
	db, err := storm.Open(dbFilePath)
	if err != nil {
		return fmt.Errorf("error while opening DB: %v", err)
	}
	defer db.Close()

	return db.Save(&task)
}

// List tasks present in the BoltDB bucket based on the filter (task status)
func ListTasks(filter int) error {
	dbFilePath := GetDBFilePath()
	db, err := storm.Open(dbFilePath)
	if err != nil {
		return fmt.Errorf("error while opening DB: %v", err)
	}
	defer db.Close()

	var tasks []Task
	err = db.All(&tasks)
	if err != nil {
		return err
	}

	for _, task := range tasks {
		if task.CompletionStatus == filter {
			fmt.Printf("%v. %v\n", task.ID, task.Name)
		}
	}
	return nil
}

// Mark a task as completed in the BoltDB bucket
func DoTask(taskID uint64) error {
	dbFilePath := GetDBFilePath()
	db, err := storm.Open(dbFilePath)
	if err != nil {
		return fmt.Errorf("error while opening DB: %v", err)
	}
	defer db.Close()

	return db.Update(&Task{ID: taskID, CompletionStatus: constants.COMPLETE, CompletionDate: time.Now()})
}

// Remove a task from the BoltDB Bucket
func RemoveTask(taskID uint64) error {
	dbFilePath := GetDBFilePath()
	db, err := storm.Open(dbFilePath)
	if err != nil {
		return fmt.Errorf("error while opening DB: %v", err)
	}
	defer db.Close()

	var task Task
	err = db.One("ID", taskID, &task)
	if err != nil {
		return err
	}
	return db.DeleteStruct(&task)
}

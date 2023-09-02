/*
Copyright © 2023 Aronya Baksy abaksy@gmail.com
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/boltdb/bolt"

	constants "github.com/abaksy/gotaskmgr/const"
)

func AddTask(task Task) error {
	db, err := bolt.Open(constants.DATABASE_FILE, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(constants.BUCKET_NAME))
		if err != nil {
			return fmt.Errorf("error creating bucket: %s", err)
		}
		id, err := b.NextSequence()
		if err != nil {
			return fmt.Errorf("error generating unique ID: %s", err)
		}
		task.ID = id

		if buf, err := json.Marshal(task); err != nil {
			return fmt.Errorf("error marshalling json! %s", err)
		} else if err := b.Put([]byte(strconv.FormatUint(task.ID, 10)), buf); err != nil {
			return err
		}
		return nil
	})

	return err
}

func ListTasks() error {
	db, err := bolt.Open(constants.DATABASE_FILE, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var allTasks []Task
	err = db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(constants.BUCKET_NAME))

		if b == nil {
			return fmt.Errorf("no tasks to list")
		}

		b.ForEach(func(k, v []byte) error {
			task := Task{}
			err := json.Unmarshal(v, &task)
			if err != nil {
				return err
			}
			allTasks = append(allTasks, task)
			return nil
		})
		return nil
	})

	if err != nil {
		return err
	}

	for _, t := range allTasks {
		fmt.Printf("%v. %v\n", t.ID, t.Name)
	}
	return nil
}

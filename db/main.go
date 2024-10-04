package db

import (
	"encoding/json"
	"errors"
	"io/fs"
	"log"
	"os"
)

type Account struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Number    int64  `json:"number"`
	Balance   int64  `json:"balance"`
}

type Accounts []Account

type Data struct {
	Accounts Accounts `json:"accounts"`
}

type DataBase struct {
	File string
	Data Data `json:"data"`
}

func (d *DataBase) Read() error {
	content, err := os.ReadFile(d.File)

	if err != nil {
		return err
	}

	data := &Data{}
	if err := json.Unmarshal(content, data); err != nil {
		d.Data = *data
		return nil
	}

	log.Fatalf("Failed to read %s due to error %s", d.File, err)
	return err
}

func (d *DataBase) Write() error {
	jsonDB, _ := json.Marshal(d.Data)

	err := os.WriteFile(d.File, jsonDB, 0777)

	if err != nil {
		log.Printf("Failed to write to file %s due to %s", d.File, err)
	}
	return nil
}

func initDataBase(filePath string) *DataBase {
	db := &DataBase{
		File: filePath,
	}

	content := db.Read()
	if !errors.Is(content, fs.ErrNotExist) {
		log.Fatalf("Failed to read file %s due to %s", db.File, content)
	}

	os.Create(filePath)
	db.Write()

	return db
}

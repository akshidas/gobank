package accounts

import (
	"encoding/json"
	"errors"
	"io/fs"
	"log"
	"os"
)

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

	if err := json.Unmarshal(content, d.Data); err != nil {
		log.Fatalf("Failed to read %s due to error :s", d.File, err)
	}

	return nil
}

func (d *DataBase) Write() error {
	jsonDB, _ := json.Marshal(d.Data)

	err := os.WriteFile(d.File, jsonDB, 0777)

	if err != nil {
		log.Println("Failed to write to file %s due to %s", d.File, err)
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

func getAll() Accounts {
	return initAccounts()
}

func getById(id int) (Account, int) {
	accounts := initAccounts()
	for i, v := range accounts {
		if v.ID == id {
			return v, i
		}
	}
	return Account{}, -1
}

func Add(account Account) {
	accounts := initAccounts()
	accounts = append(accounts, account)
	writeAccount(accounts)
}

func Update(id int, updateData Account) error {
	accounts := initAccounts()
	for i, v := range accounts {
		if v.ID == id {
			accounts[i] = updateData
			writeAccount(accounts)
			return nil
		}
	}
	return nil
}

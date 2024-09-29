package accounts

import (
	"encoding/json"
	"errors"
	"io/fs"
	"log"
	"os"
)

type Accounts []Account

type AccountsFile struct {
	Data Accounts `json:"data"`
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

func writeAccount(account Accounts) {
	file := &AccountsFile{
		Data: account,
	}
	jsonFile, _ := json.Marshal(file)
	err := os.WriteFile(".account.json", jsonFile, 0777)
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func initAccounts() Accounts {
	content, err := os.ReadFile(".account.json")
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			os.Create(".account.json")
			t := Accounts{}
			writeAccount(t)
			return t
		}
		log.Fatalf("Error while reading a file %v", err)
	}

	a := Accounts{}
	file := &AccountsFile{
		Data: a,
	}

	err = json.Unmarshal(content, file)

	if err != nil {
		log.Fatalf("Error while unmarshal the content  %v", err)
	}

	return file.Data

}

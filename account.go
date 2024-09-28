package main

type Account struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Number    int64  `json:"number"`
	Balance   int64  `json:"balance"`
}

type DataBase struct {
	accounts []Account
}

func (d *DataBase) getAccounts() []Account {
	return d.accounts
}

func (d *DataBase) getAccountById(id int) *Account {
	for _, v := range d.accounts {
		if v.ID == id {
			return &v
		}
	}
	return nil
}

func (d *DataBase) AddAccount(account *Account) {
	d.accounts = append(d.accounts, *account)
}

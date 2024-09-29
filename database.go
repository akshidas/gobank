package main

type DataBase struct {
	accounts []Account
}

func (d *DataBase) getAll() []Account {
	return d.accounts
}

func (d *DataBase) getById(id int) (*Account, int) {
	for i, v := range d.accounts {
		if v.ID == id {
			return &v, i
		}
	}
	return nil, -1
}

func (d *DataBase) Add(account *Account) {
	d.accounts = append(d.accounts, *account)
}

func (d *DataBase) Update(id int, updateData Account) {

	account, pos := d.getById(id)

	if account != nil {
		d.accounts[pos] = updateData
	}

}

package accounts

type DataBase struct {
	Accounts []Account
}

func (d *DataBase) getAll() []Account {
	return d.Accounts
}

func (d *DataBase) getById(id int) (*Account, int) {
	for i, v := range d.Accounts {
		if v.ID == id {
			return &v, i
		}
	}
	return nil, -1
}

func (d *DataBase) Add(account *Account) {
	d.Accounts = append(d.Accounts, *account)
}

func (d *DataBase) Update(id int, updateData Account) {

	account, pos := d.getById(id)

	if account != nil {
		d.Accounts[pos] = updateData
	}

}

package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositorySub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Kil", "Juja", "12038", "2000-01-01", "1"},
		{"1002", "Lik333333", "Juja", "12038", "2000-01-01", "1"},
	}

	return CustomerRepositoryStub{customers}
}

package services

import (
	"bytes"
	"encoding/csv"
	"io"
	"io/ioutil"
)

// CustomerService is responsible for CRUD operations related to customer records.
type CustomerService struct {
	DBFilePath string
}

// Customer Type
type Customer struct {
	Name    string `json:"name"`
	UUID    string `json:"uuid"`
	Email   string `json:"email"`
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
}

// NewCustomerService will instantiate and return a pointer to a customer service type
func NewCustomerService(db string) *CustomerService {
	return &CustomerService{DBFilePath: db}
}

// FindAll mocks what would be a database query.  In it's simple form, all records are returned
// and mapped to a customer type and stuffed into a slice.
func (s *CustomerService) FindAll() ([]Customer, error) {

	var customers []Customer

	b, err := ioutil.ReadFile(s.DBFilePath)
	if err != nil {
		return customers, err
	}

	buff := bytes.NewBuffer(b)
	r := csv.NewReader(buff)

	// Throw away the header row
	r.Read()

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		customer := Customer{
			UUID:    record[0],
			Name:    record[1],
			Email:   record[2],
			Address: record[3],
			City:    record[4],
			State:   record[5],
			Zip:     record[6],
		}

		customers = append(customers, customer)
	}

	return customers, nil
}

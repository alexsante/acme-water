package services

import (
	"testing"
)

func TestFindAll(t *testing.T) {

	cs := NewCustomerService("../db/customers.csv")

	t.Log("Given the need to retrieve all customers")
	{
		customers, err := cs.FindAll()
		if len(customers) == 0 {
			t.Fatal("Expected FindAll to return a slice with at least 5 customer types")
		}
		if err != nil {
			t.Fatal(err)
		}
	}

}

func TestDBError(t *testing.T) {

	cs := NewCustomerService("../db/badfile.csv")

	t.Log("Given the need to handle a database error")
	{
		_, err := cs.FindAll()
		if err == nil {
			t.Fatal("Expected NewCustomerService to return an error due to an invalid file path")
		}
	}

}

package services

import (
	"strconv"
	"strings"
	"testing"
)

func TestDispatch(t *testing.T) {

	ns := NewNotificationService()

	t.Log("Given the need to notify a customer of a new invoice")
	{
		c := Customer{
			Name:    "John Doe",
			Address: "123 Park Lane",
			City:    "New York",
			State:   "NY",
			Zip:     "10001",
		}
		inv := Invoice{Customer: c, AmountDue: 100.00}
		inv.ParsedAmountDue = strconv.FormatFloat(float64(inv.AmountDue), 'f', 2, 32)
		body, err := ns.Dispatch("nobody@test.com", "nobody@test.com", "Testing", inv)

		if err != nil {
			t.Fatal("Expected err to be nil but got ", err)
		}

		if strings.Contains(body, c.Name) == false {
			t.Fatal("Expected the notification to contain the customers name")
		}

		if strings.Contains(body, c.Address) == false {
			t.Fatal("Expected the notification to contain the customers address")
		}

		if strings.Contains(body, c.City) == false {
			t.Fatal("Expected the notification to contain the customers city")
		}

		if strings.Contains(body, c.State) == false {
			t.Fatal("Expected the notification to contain the customers state")
		}

		if strings.Contains(body, c.Zip) == false {
			t.Fatal("Expected the notification to contain the customers zip")
		}

		if strings.Contains(body, inv.ParsedAmountDue) == false {
			t.Fatal("Expected the notification to contain the invoice total ", inv.ParsedAmountDue)
		}

	}

}

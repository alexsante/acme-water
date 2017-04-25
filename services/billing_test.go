package services

import (
	"testing"
)

func TestExecute(t *testing.T) {

	cs := NewCustomerService("../db/customers.csv")
	ms := NewMeterService("http://localhost")
	ns := NewNotificationService()
	bs := NewBillingService(cs, ms, ns)

	t.Log("Given the need to bill all customers")
	{

		summary, err := bs.Execute()
		if err != nil {
			t.Fatal("Expected err to be nil but got ", err)
		}
		if summary.InvoicesSent != 5 {
			t.Fatal("Expected 5 invoices to be sent out but got ", summary.InvoicesSent)
		}
		if summary.TotalBilled != 129.90 {
			t.Fatal("Expected the total billed amount to be 129.90 but got ", summary.TotalBilled)
		}

	}

}

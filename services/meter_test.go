package services

import (
	"testing"
	"time"
)

func TestReconcile(t *testing.T) {

	ms := NewMeterService("http://localhost")
	t.Log("Given the need to reconcile an invoice for a single customer")
	{

		c := Customer{}
		now := time.Now()
		due, err := ms.Reconcile(c, now.Month(), now.Year())

		if due != 25.98 {
			t.Fatal("Expected the amount due to be 25.98")
		}

		if err != nil {
			t.Fatal("Expected a nil error but got ", err)
		}

		if ms.retries() != 3 {
			t.Fatal("Expected a total of 3 retries but got ", ms.retries())
		}

	}

}

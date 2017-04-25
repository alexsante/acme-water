package services

import (
	"strconv"
	"time"
)

// BillingService is responsible for coordinating the execution of a billing cycle.
type BillingService struct {
	cs *CustomerService
	ms *MeterService
	ns *NotificationService
}

// Invoice generated in a billing cycle.
type Invoice struct {
	Customer        Customer
	Month           time.Month
	Year            int
	AmountDue       float32
	ParsedAmountDue string
}

// BillingSummary is used to store reporting metrics for a billing cycle.
type BillingSummary struct {
	InvoicesSent int
	TotalBilled  float32
}

// NewBillingService will instantiate and return a pointer to a billing service type.
func NewBillingService(cs *CustomerService, ms *MeterService, ns *NotificationService) *BillingService {
	return &BillingService{
		cs: cs,
		ms: ms,
		ns: ns,
	}
}

// Execute will process all customer bills for the current month and year.  A notification will be
// dispatched to each customer with their amount due.
func (s *BillingService) Execute() (BillingSummary, error) {

	var invoices []Invoice
	var summary BillingSummary

	// Fetch all customers in the database
	customers, err := s.cs.FindAll()
	if err != nil {
		return summary, err
	}

	// For each customer, ask the metering service for their amount due
	for _, c := range customers {

		inv := Invoice{
			Customer: c,
			Month:    time.Now().Month(),
			Year:     time.Now().Year(),
		}

		// Fetch amount due from the metering service
		due, err := s.ms.Reconcile(c, inv.Month, inv.Year)
		if err != nil {
			return summary, err
		}

		// Amount due fetched from the metering service
		inv.AmountDue = due
		inv.ParsedAmountDue = "$" + strconv.FormatFloat(float64(inv.AmountDue), 'f', 2, 32)

		// Append to the invoice slice because it's going to be used by the notification service
		invoices = append(invoices, inv)

	}

	// For each invoice, notify the customer
	for _, i := range invoices {
		s.ns.Dispatch(i.Customer.Email, "noreply@acmewater.com", "Your invoice is ready!", i)
		// Update billing summary
		summary.TotalBilled += i.AmountDue
		summary.InvoicesSent++
	}

	return summary, nil
}

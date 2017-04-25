package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var Retries int

// MeterService serves as the adapter to the metering system.
type MeterService struct {
	Endpoint string
	// Maximum number of times the billing service will attempt to communicate with the metering
	// service when an error is returned
	maxRetries int
	// Amount of time (ms) to wait between retries
	retrySpan time.Duration
}

// MeterResponse contains the structure expected back from the metering system.
type MeterResponse struct {
	AmountDue float32 `json:"amount_due"`
}

// NewMeterService will create an instance of a MeterService when passed an endpoint
func NewMeterService(e string) *MeterService {
	s := &MeterService{
		Endpoint:   e,
		maxRetries: 3,
		retrySpan:  2000,
	}

	return s
}

// Reconcile queries the metering api for a customer's balance as of the month and year passed in.
func (s *MeterService) Reconcile(c Customer, m time.Month, y int) (float32, error) {

	resp := `{"amount_due": 25.98}`
	var mr MeterResponse

	// Ignore (e.g _,_ ) the response since this is a fake endpoint
	_, _ = http.Get(fmt.Sprintf("%v/due/%v/%v/%v", s.Endpoint, c.UUID, m, y))

	// Mock API failures.  This bit of code will attempt to execute 2 times before finally working on the 3rd attempt.
	if Retries < s.maxRetries {
		fmt.Println("Metric Service Failure - Trying again")
		Retries++
		time.Sleep(s.retrySpan * time.Millisecond)
		s.Reconcile(c, m, y)
	}

	err := json.Unmarshal([]byte(resp), &mr)
	if err != nil {
		return 0.00, err
	}

	return mr.AmountDue, nil
}

func (s *MeterService) retries() int {
	return Retries
}

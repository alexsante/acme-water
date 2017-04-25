/**
Acme WaterTM is looking to take their monthly billing out of spreadsheets and into an automated
system that email delivers billing statements to its customers.

Acme Water's metering team has recently exposed a REST web service that provides an "amount due"
in exchange for a combination of customer UUID and month of the year. The business department has
presented you with a CSV file with all existing customer data. Using the requirements defined below,
create your model and define your business logic that will enable Acme WaterTM to email all existing
customers their statements based on the "amount due" provided by the web service.
**/

package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/alexsante/acme-water/services"
)

func main() {

	// Initialize required services.  Billing service must go last because it depends
	// on all other services.
	fmt.Println(" -- Initializing Services -- ")
	ms := services.NewMeterService("http://localhost")
	cs := services.NewCustomerService("db/customers.csv")
	ns := services.NewNotificationService()
	bs := services.NewBillingService(cs, ms, ns)
	fmt.Println(" -- Service Initialization Complete -- ")

	summary, err := bs.Execute()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Billing cycle has been completed!\n  Total invoiced amount: $%v\n  Invoices Dispatched: %v", strconv.FormatFloat(float64(summary.TotalBilled), 'f', 2, 32), summary.InvoicesSent)

}

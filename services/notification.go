package services

import (
	"bytes"
	"fmt"
	"html/template"
)

// NotificationService is reponsible for dispatching notifications.  This mock service only logs out what would be
// email notifications.  SMS, SQS, SNS can all be added.
type NotificationService struct{}

// NewNotificationService instantiates a new notification service type
func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

// Dispatch a notification. Only email notifications are currently supported.
func (s *NotificationService) Dispatch(to, from, subject string, payload interface{}) (string, error) {

	tmpl, err := template.New("notification").Parse(ReceiptTmpl)
	if err != nil {
		fmt.Println("Unable to dispatch a notification because the template could not be parsed. - ", err)
	}

	buff := new(bytes.Buffer)
	err = tmpl.Execute(buff, payload)
	if err != nil {
		fmt.Println("Unable to dispatch a notification.  Payload Parsing Error. - ", err)
		return "", err
	}

	fmt.Println(" -- Dispatching Notification -- ")
	fmt.Println(buff.String())
	fmt.Println(" -- Notification Dispatched -- ")

	return buff.String(), nil
}

var ReceiptTmpl = `Dear {{.Customer.Name}},
Thank you for using Acme WaterTM for your address at {{.Customer.Address}} {{.Customer.City}}, {{.Customer.State}} {{.Customer.Zip}}. Your amount due for the month of {{.Month}} is {{.ParsedAmountDue}}.
Warm Regards, Acme WaterTM`

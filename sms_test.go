package messagebird_api

import (
	"log"
	"testing"
)

func TestSms(t *testing.T) {

	api := NewApi("6ahazAMtOpOmsctO7NhMe8vmA")

	sms := NewSms()
	sms.SetOriginator("JAKOUBEK")
	sms.AddRecipient("00491785384659")
	sms.SetMessageBody("Test-SMS via API")

	_, err := api.SendSms(sms)
	if err != nil {
		log.Fatal(err)
	}

}

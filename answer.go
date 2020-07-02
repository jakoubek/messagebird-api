package messagebird_api

import "time"

type ServiceAnswer struct {
	ID          string      `json:"id"`
	Href        string      `json:"href"`
	Direction   string      `json:"direction"`
	Type        string      `json:"type"`
	Originator  string      `json:"originator"`
	Body        string      `json:"body"`
	Reference   interface{} `json:"reference"`
	Validity    interface{} `json:"validity"`
	Gateway     int         `json:"gateway"`
	TypeDetails struct {
	} `json:"typeDetails"`
	Datacoding        string      `json:"datacoding"`
	Mclass            int         `json:"mclass"`
	ScheduledDatetime interface{} `json:"scheduledDatetime"`
	CreatedDatetime   time.Time   `json:"createdDatetime"`
	Recipients        struct {
		TotalCount               int `json:"totalCount"`
		TotalSentCount           int `json:"totalSentCount"`
		TotalDeliveredCount      int `json:"totalDeliveredCount"`
		TotalDeliveryFailedCount int `json:"totalDeliveryFailedCount"`
		Items                    []struct {
			Recipient        int64     `json:"recipient"`
			Status           string    `json:"status"`
			StatusDatetime   time.Time `json:"statusDatetime"`
			MessagePartCount int       `json:"messagePartCount"`
		} `json:"items"`
	} `json:"recipients"`
}


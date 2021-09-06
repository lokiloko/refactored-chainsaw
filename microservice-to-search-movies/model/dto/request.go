package dto

type (
	WebServiceRequestParams struct {
		Url     string
		Method  string
		Payload interface{}
	}
)
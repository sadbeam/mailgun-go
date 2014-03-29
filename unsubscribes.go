package mailgun

import (
	"github.com/mbanzon/simplehttp"
	"strconv"
)

type Unsubscription struct {
	CreatedAt string `json:"created_at"`
	Tag       string `json:"tag"`
	ID        string `json:"id"`
	Address   string `json:"address"`
}

// GetUnsubscribes retrieves a list of unsubscriptions issued by recipients of mail from your domain.
// Zero is a valid list length.
func (mg *mailgunImpl) GetUnsubscribes(limit, skip int) (int, []Unsubscription, error) {
	r := simplehttp.NewHTTPRequest(generateApiUrl(mg, unsubscribesEndpoint))
	if limit != DefaultLimit {
		r.AddParameter("limit", strconv.Itoa(limit))
	}
	if skip != DefaultSkip {
		r.AddParameter("skip", strconv.Itoa(skip))
	}
	r.SetBasicAuth(basicAuthUser, mg.ApiKey())
	var envelope struct {
		TotalCount int              `json:"total_count"`
		Items      []Unsubscription `json:"items"`
	}
	err := r.GetResponseFromJSON(&envelope)
	return envelope.TotalCount, envelope.Items, err
}

// GetUnsubscribesByAddress retrieves a list of unsubscriptions by recipient address.
// Zero is a valid list length.
func (mg *mailgunImpl) GetUnsubscribesByAddress(a string) (int, []Unsubscription, error) {
	r := simplehttp.NewHTTPRequest(generateApiUrlWithTarget(mg, unsubscribesEndpoint, a))
	r.SetBasicAuth(basicAuthUser, mg.ApiKey())
	var envelope struct {
		TotalCount int              `json:"total_count"`
		Items      []Unsubscription `json:"items"`
	}
	err := r.GetResponseFromJSON(&envelope)
	return envelope.TotalCount, envelope.Items, err
}

// Unsubscribe adds an e-mail address to the domain's unsubscription table.
func (mg *mailgunImpl) Unsubscribe(a, t string) error {
	r := simplehttp.NewHTTPRequest(generateApiUrl(mg, unsubscribesEndpoint))
	r.SetBasicAuth(basicAuthUser, mg.ApiKey())
	p := simplehttp.NewUrlEncodedPayload()
	p.AddValue("address", a)
	p.AddValue("tag", t)
	_, err := r.MakePostRequest(p)
	return err
}

// RemoveUnsubscribe removes the e-mail address given from the domain's unsubscription table.
// If passing in an ID (discoverable from, e.g., GetUnsubscribes()), the e-mail address associated
// with the given ID will be removed.
func (mg *mailgunImpl) RemoveUnsubscribe(a string) error {
	r := simplehttp.NewHTTPRequest(generateApiUrlWithTarget(mg, unsubscribesEndpoint, a))
	r.SetBasicAuth(basicAuthUser, mg.ApiKey())
	_, err := r.MakeDeleteRequest()
	return err
}

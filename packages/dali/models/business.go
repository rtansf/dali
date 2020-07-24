package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Business struct {
	Token  string `json:"token"`
	Incorporation Incorporation `json:"incorporation"`
	BusinessNameLegal string `json:"business_name_legal"`
	BusinessNameDba string `json:"business_name_dba"`
	OfficeLocation Address `json:"office_location"`
	ProprietorOrOfficer Personinfo `json:"proprietor_or_officer"`
	BeneficialOwner1 Personinfo `json:"beneficial_owner1"`
	BeneficialOwner2 Personinfo `json:"beneficial_owner2"`
	BeneficialOwner3 Personinfo `json:"beneficial_owner3"`
	BeneficialOwner4 Personinfo `json:"beneficial_owner4"`
	AttesterName string `json:"attester_name"`
	AttesterTitle string `json:"attester_title"`
	AttestationDate string `json:"attestation_date"`
	AttestationConsent bool `json:"attestation_consent"`
}

type Incorporation struct {
	IncorporationType string `json:"incorporation_type"`
}

type Personinfo struct {
        FirstName string   `json:"first_name"`
	MiddleName string  `json:"middle_name"`
	LastName string    `json:"last_name"`
	Home Address       `json:"home"`
	Dob string 	       `json:"dob"`
	Ssn string    	   `json:"ssn"`
}

type Address struct {
	Address1   string `json:"address1"`
	Address2   string `json:"address2"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}

func GetBusiness(token string) Business {

	req, _ := http.NewRequest("GET", "http://localhost:4280/v3/businesses/" + token, nil)
	req.SetBasicAuth("api_consumer", "marqeta")
	timeout := time.Duration(5 * time.Second)
	client := http.Client {
		Timeout : timeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(responseData))

	var b Business
	json.Unmarshal(responseData, &b)

	return b
}
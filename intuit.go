package functions

import (
	"XprtLive/dao"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type QuickAdd struct {
	GivenName string `json:"GivenName"`
	SSN string `json:"SSN"`
	PrimaryAddr map[string]string `json:"PrimaryAddr"`
	PrimaryPhone map[string]string `json:"PrimaryPhone"`
	FamilyName string `json:"FamilyName"`
}

func AddEmployeeIntuit(name, ssn, city, pin, line, phone, familyName string) (string, error) {
	client := http.Client{}
	emp := QuickAdd{name, ssn, map[string]string{"CountrySubDivisionCode":"IN", "City":city, "PostalCode":pin, "Line1":line}, map[string]string{"FreeFormNumber":phone}, familyName}
	body, _ := json.Marshal(&emp)
	req, err := http.NewRequest("POST", dao.DIRECTLY_INTUIT+"/v3/company/"+dao.INTUIT_ID+"/employee", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error in creating new call", err)
	}
	req.Header.Set("content-type", "application/json")
	fmt.Println("REQQ ", req)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error in getting new call", err)
		return "Error in Retrieving your booking info", errors.New("RETIREIVING ERROR")
	}
	body, _ = json.Marshal(&res.Body)
	req, err = http.NewRequest("POST", dao.DIRECTLY_LIVE_CUSTOMER_BACKEND+"updateEmploye", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error in creating new call", err)
	}
	req.Header.Set("content-type", "application/json")
	fmt.Println("REQQ ", req)
	res, err = client.Do(req)
	if err != nil {
		fmt.Println("Error in getting new call", err)
		return "Error in Retrieving your booking info", errors.New("RETIREIVING ERROR")
	}
	return "Succes", nil
}

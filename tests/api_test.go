package main

import (
	"fmt"
	"github.com/cucumber/godog"
	m "goFruits/bundles"
	"io/ioutil"
	"net/http"
	"strings"
)

var host = "http://localhost:8080/api/v1"

var res *http.Response

func aRequestWithPayloadIsSentToTheEndpoint(method, endpoint string, payload string) error {
	var reader = strings.NewReader(payload)
	var request, err = http.NewRequest(method, host+endpoint,  reader)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		return fmt.Errorf("could not create request %s", err.Error())
	}
	res, err = http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("could not send request %s", err.Error())
	}
	return nil
}

func aRequestIsSentToTheEndpoint(method, endpoint string) error {
	var reader = strings.NewReader("")
	var request, err = http.NewRequest(method, host+endpoint,  reader)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		return fmt.Errorf("could not create request %s", err.Error())
	}
	res, err = http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("could not send request %s", err.Error())
	}
	return nil
}

func theHTTPresponseCodeShouldBe(expectedCode int) error {
	if expectedCode != res.StatusCode {
		return fmt.Errorf("status code not as expected! Expected '%d', got '%d'", expectedCode, res.StatusCode)
	}
	return nil
}

func theResponseContentShouldBe(expectedContent string) error {
	body, _ := ioutil.ReadAll(res.Body)
	if expectedContent != strings.TrimSpace(string(body)) {
		return fmt.Errorf("status code not as expected! Expected '%s', got '%s'", expectedContent, string(body))
	}
	return nil
}

func removeAllFruits () error {
	m.WriteCSV("../data.csv", nil)
	return nil
}


func FeatureContext(s *godog.Suite) {
	s.Step(`^I don\'t have fruits$`, removeAllFruits)
	s.Step(`^I add a fruit with "([^"]*)" request to the endpoint "([^"]*)" with the following payload '(.*)'$`, aRequestWithPayloadIsSentToTheEndpoint)
	s.Step(`^the HTTP-response code should be "(\d+)"$`, theHTTPresponseCodeShouldBe)
	s.Step(`^the response content should be '(.*)'$`, theResponseContentShouldBe)
	s.Step(`^I retrieve fruit info with "([^"]*)" request to the endpoint "([^"]*)"$`, aRequestIsSentToTheEndpoint)
	s.Step(`^I update a fruit with "([^"]*)" request to the endpoint "([^"]*)" with the following payload '(.*)'$`, aRequestWithPayloadIsSentToTheEndpoint)
	s.Step(`^I delete a fruit with "([^"]*)" request to the endpoint "([^"]*)"$`, aRequestIsSentToTheEndpoint)

}

package main

import (
	"jsontostruct/utils"
	"jsontostruct/validation"
	"testing"
)

func TestValid(t *testing.T) {

	hosts := utils.FetchHosts()
	expectedHardwareDetails := utils.ExtractExpectedHardwareDetails()

	validHosts := validation.Valid(hosts, expectedHardwareDetails)

	if len(validHosts) > 0 {
		t.Logf("Host Match Found")
	} else {
		t.Errorf("Host Match Not Found")
	}

}

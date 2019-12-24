package main

import (
	"fmt"
	"jsontostruct/utils"
	"jsontostruct/validation"
)

func main() {

	hosts := utils.FetchHosts()
	expectedHardwareDetails := utils.ExtractExpectedHardwareDetails()

	fmt.Println(hosts)
	fmt.Println(expectedHardwareDetails)

	validHostList := validation.Valid(hosts, expectedHardwareDetails)
	fmt.Println(validHostList)
}

package main

import (
	"fmt"
	"jsontostruct/utils"
	"jsontostruct/validation"
	"time"
)

func main() {

	startTime := time.Now()
	hosts := utils.FetchHosts()
	fmt.Println("Host Array Size", len(hosts.Host))
	expectedHardwareDetails := utils.ExtractExpectedHardwareDetails()
	validHostList := validation.Valid(hosts, expectedHardwareDetails)
	fmt.Println(validHostList)
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

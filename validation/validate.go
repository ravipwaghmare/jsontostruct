package validation

import (
	"fmt"
	"jsontostruct/ironic"
	"sync"
)

var jobs = make(chan ironic.Job, 50)
var results = make(chan ironic.Result, 50)

// Valid returns true or false according to the test condition
func Valid(hostList ironic.Data, expectedHardwareConfig ironic.ExpectedHardwareConfiguration) []ironic.Host {
	go allocate(hostList, expectedHardwareConfig)
	done := make(chan []ironic.Host)
	go result(done)
	createWorkerPool(hostList)
	hosts := <-done
	// Your Logic here
	return hosts
}

func checkValidation(host ironic.Host, expected ironic.ExpectedHardwareConfiguration) bool {
	return false
}

func allocate(hostList ironic.Data, expectedHardwareConfig ironic.ExpectedHardwareConfiguration) {
	for i := 0; i < len(hostList.Host); i++ {
		job := ironic.Job{i, hostList.Host[i], expectedHardwareConfig}
		jobs <- job
	}

	close(jobs)
}

func result(done chan []ironic.Host) {
	var validHost []ironic.Host
	for result := range results {
		validHost = append(validHost, result.Job.Host)
	}
	done <- validHost
}

func createWorkerPool(hostList ironic.Data) {
	var wg sync.WaitGroup
	for i := 0; i < len(hostList.Host); i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := ironic.Result{job, isValid(job)}
		results <- output
	}
	wg.Done()
}

func isValid(job ironic.Job) bool {

	if job.ExpectedHardwareConfiguration.ExpectedCPU.Count < job.Host.Spec.HardwareDetails.CPU.Count {
		fmt.Println("CPU Valid, Now Check Memory")
		if job.ExpectedHardwareConfiguration.ExpectedDisk.SizeBytesGB < job.Host.Spec.HardwareDetails.Storage.SizeBytes*1024*1024 {
			fmt.Println("Memory Valid, Now Check Nics")
			if job.ExpectedHardwareConfiguration.ExpectedNICS.NumberOfNICS < job.Host.Spec.HardwareDetails.Nics.NoOfNics {
				fmt.Println("Number of Nics are avialale, Now Check RAM")
				if job.ExpectedHardwareConfiguration.ExpectedRAM < job.Host.Spec.HardwareDetails.Memory.PhysicalMb {
					fmt.Println("RAM is available, Now Check Vendor Name")
					if job.ExpectedHardwareConfiguration.ExpectedSystemVendor.Name == job.Host.Spec.HardwareDetails.SystemVendor.Manufacturer {
						fmt.Println("Host Found")
						return true
					}
				}
			}
		}
	}

	return false
}

package validation

import (
	"fmt"
	"jsontostruct/ironic"
)

// Valid returns true or false according to the test condition
func Valid(hostList ironic.Data, expectedHardwareConfig ironic.Profile) []ironic.Host {
	var validHost []ironic.Host
	// Your Logic here
	for _, host := range hostList.Host {

		for _, profile := range expectedHardwareConfig.ExpectedHardwareConfiguration {
			if profile.ExpectedCPU.Count < host.Spec.HardwareDetails.CPU.Count {
				fmt.Println("CPU Valid, Now Check Memory")
				if profile.ExpectedDisk.SizeBytesGB < host.Spec.HardwareDetails.Storage.SizeBytes*1024*1024 {
					fmt.Println("Memory Valid, Now Check Nics")
					if profile.ExpectedNICS.NumberOfNICS < host.Spec.HardwareDetails.Nics.NoOfNics {
						fmt.Println("Number of Nics are avialale, Now Check RAM")
						if profile.ExpectedRAM < host.Spec.HardwareDetails.Memory.PhysicalMb {
							fmt.Println("RAM is available, Now Check Vendor Name")
							if profile.ExpectedSystemVendor.Name == host.Spec.HardwareDetails.SystemVendor.Manufacturer {
								fmt.Println("Host Found")
								validHost = append(validHost, host)
							}
						}
					}
				}
			}
		}

	}
	return validHost
}

package validation

import (
	"jsontostruct/ironic"
)

// Valid returns true or false according to the test condition
func Valid(hostList ironic.Data, expectedHardwareConfig ironic.Profile) map[interface{}][]ironic.ExpectedHardwareConfiguration {

	validHost := make(map[interface{}][]ironic.ExpectedHardwareConfiguration)

	for _, host := range hostList.Host {
		for _, profile := range expectedHardwareConfig.ExpectedHardwareConfiguration {
			if profile.ExpectedCPU.Count < host.Spec.HardwareDetails.CPU.Count {
				if profile.ExpectedDisk.SizeBytesGB < host.Spec.HardwareDetails.Storage.SizeBytes*1024*1024 {
					if profile.ExpectedNICS.NumberOfNICS < host.Spec.HardwareDetails.Nics.NoOfNics {
						if profile.ExpectedRAM < host.Spec.HardwareDetails.Memory.PhysicalMb {
							if profile.ExpectedSystemVendor.Name == host.Spec.HardwareDetails.SystemVendor.Manufacturer {
								newHost, ok := validHost[host.Spec.HardwareDetails]
								if ok {
									validHost[host.Spec.HardwareDetails] = append(newHost, profile)
								} else {
									var validProfile []ironic.ExpectedHardwareConfiguration
									validHost[host.Spec.HardwareDetails] = append(validProfile, profile)
								}

							}
						}
					}
				}
			}
		}

	}
	return validHost
}

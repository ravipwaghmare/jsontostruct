package ironic

//Data ironic introspected data
type Data struct {
	Host []struct {
		Spec struct {
			HardwareDetails struct {
				Boot struct {
					CurrentBootMode string `json:"current_boot_mode"`
					PxeInterface    string `json:"pxe_interface"`
				} `json:"boot"`
				SystemVendor struct {
					SerialNumber string `json:"serial_number"`
					ProductName  string `json:"product_name"`
					Manufacturer string `json:"manufacturer"`
				} `json:"system_vendor"`
				Memory struct {
					PhysicalMb int   `json:"physical_mb"`
					Total      int64 `json:"total"`
				} `json:"memory"`
				CPU struct {
					Count        int      `json:"count"`
					Frequency    string   `json:"frequency"`
					Flags        []string `json:"flags"`
					ModelName    string   `json:"model_name"`
					Architecture string   `json:"architecture"`
				} `json:"cpu"`
			} `json:"HardwareDetails"`
		} `json:"Spec"`
	} `json:"host"`
}

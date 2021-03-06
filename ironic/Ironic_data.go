package ironic

// Data ironic introspected data
type Data struct {
	Host []Host `json:"host"`
}

// Host Spec of the host
type Host struct {
	Spec Spec `json:"Spec"`
}

// Spec hardware details for the spec
type Spec struct {
	HardwareDetails HardwareDetails `json:"HardwareDetails"`
}

// HardwareDetails config details
type HardwareDetails struct {
	Boot         Boot         `json:"boot"`
	SystemVendor SystemVendor `json:"system_vendor"`
	Memory       Memory       `json:"memory"`
	CPU          CPU          `json:"cpu"`
	Storage      Storage      `json:"storage"`
	Nics         Nics         `json:"nics"`
}

// Boot boot details
type Boot struct {
	CurrentBootMode string `json:"current_boot_mode"`
	PxeInterface    string `json:"pxe_interface"`
}

// SystemVendor vendor details
type SystemVendor struct {
	SerialNumber string `json:"serial_number"`
	ProductName  string `json:"product_name"`
	Manufacturer string `json:"manufacturer"`
}

// CPU cpu details
type CPU struct {
	Count        int      `json:"count"`
	Frequency    string   `json:"frequency"`
	Flags        []string `json:"flags"`
	ModelName    string   `json:"model_name"`
	Architecture string   `json:"architecture"`
}

// Memory memory details
type Memory struct {
	PhysicalMb int `json:"physical_mb"`
	Total      int `json:"total"`
}

// Storage details
type Storage struct {
	SizeBytes     int `json:"size_bytes"`
	NumberOfDisks int `json:"noOfDisks"`
}

// Nics count
type Nics struct {
	NoOfNics int `json:"noOfNics"`
}

//Profile Expected hardware profile to match with the host
type Profile struct {
	ExpectedHardwareConfiguration []ExpectedHardwareConfiguration `json:"expectedValidationConfiguration"`
}

// ExpectedHardwareConfiguration details to match with the host
type ExpectedHardwareConfiguration struct {
	ExpectedCPU          ExpectedCPU          `json:"expectedCPU"`
	ExpectedDisk         ExpectedDisk         `json:"expectedDisk"`
	ExpectedNICS         ExpectedNICS         `json:"expectedNICS"`
	ExpectedRAM          int                  `json:"expectedRAM"`
	ExpectedSystemVendor ExpectedSystemVendor `json:"expectedSystemVendor"`
}

// ExpectedCPU expected cpu count
type ExpectedCPU struct {
	Count int `json:"count"`
}

// ExpectedDisk size and number of disk needed
type ExpectedDisk struct {
	SizeBytesGB   int `json:"sizeBytesGB"`
	NumberOfDisks int `json:"numberOfDisks"`
}

// ExpectedNICS count of nics cards
type ExpectedNICS struct {
	NumberOfNICS int `json:"numberOfNICS"`
}

// ExpectedSystemVendor Vendor details
type ExpectedSystemVendor struct {
	Name string `json:"name"`
}

//Job jobs to be executed
type Job struct {
	ID                            int
	Host                          Host
	ExpectedHardwareConfiguration ExpectedHardwareConfiguration
}

//Result store result of job
type Result struct {
	Job     Job
	IsValid bool
}

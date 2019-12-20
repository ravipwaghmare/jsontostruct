package ironic

// Data is extracted details from Ironic
type Data struct {
	Users []struct {
		Name   string `json:"name"`
		Type   string `json:"type"`
		Age    int    `json:"age"`
		Social struct {
			Facebook string `json:"facebook"`
			Twitter  string `json:"twitter"`
		} `json:"social"`
	} `json:"users"`
}

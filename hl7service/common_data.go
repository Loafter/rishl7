package rishl7

type Hl7ConD struct {
	SerIp string
	Port  int `json:"Port,string"`
}

type Hl7Request struct {
	Hl7c Hl7ConD
	Data string
}

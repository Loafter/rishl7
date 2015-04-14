package main

type PatientData struct {
	PtName string
	PtDatB string
	PtMrn  string
}
type Hl7ConD struct {
	SerIp string
	Port  int `json:"Port,string"`
}

type ADT01Req struct {
	Hl7c Hl7ConD
	Pd   PatientData
}

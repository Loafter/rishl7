package main

import "log"
import "rishl7/hl7service"

func main() {
	hl7j := rishl7.HL7Js{}
	log.Println("gui located add http://localhost:9979/index.html")
	hl7j.Start(9979)

}




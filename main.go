package main

import "log"

func main() {
	hl7j := HL7Js{}
	hl7j.Start(9979)
	log.Println("gui located add http://localhost:9979/index.html")
}

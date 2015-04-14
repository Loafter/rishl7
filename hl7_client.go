package main

import (
	"io/ioutil"
	"net"
)
import "strconv"
import "encoding/base64"
import "log"

const adt01Smp = "TVNIfF5+XCZ8RVBJQ0FEVHxESHxMQUJBRFR8REh8MjAxMzAxMDExMjI2fHxBRFReQTAxfEhMN01TRzAwMDAxfFB8Mi4zfA0KRVZOfEEwMXwyMDEzMDEwMTEyMjN8fA0KUElEfHx8TVJOMTIzNDVeNV5NMTF8fEFQUExFU0VFRF5KT0hOXkFeSUlJfHwxOTcxMDEwMXxNfHxDfDEgQ0FUQUxZWkUgU1RSRUVUXl5NQURJU09OXldJXjUzMDA1LTEwMjB8R0x8KDQxNCkzNzktMTIxMnwoNDE0KTI3MS0zNDM0fHxTfHxNUk4xMjM0NTAwMV4yXk0xMHwxMjM0NTY3ODl8OTg3NjU0Xk5DfA0KTksxfDF8QVBQTEVTRUVEXkJBUkJBUkFeSnxXSUZFfHx8fHx8TkteTkVYVCBPRiBLSU4NClBWMXwxfEl8MjAwMF4yMDEyXjAxfHx8fDAwNDc3N15HT09EXlNJRE5FWV5KLnx8fFNVUnx8fHxBRE18QTB8DQo="

type HL7Client struct {
}

func (dc *HL7Client) ADTA01(pd PatientData, hl7cd Hl7ConD) ([]byte, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", hl7cd.SerIp+":"+strconv.Itoa(hl7cd.Port))
	if err != nil {
		return nil, err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	defer func() {
		if conn != nil {
			conn.Close()
		}

	}()
	if err != nil {
		return nil, err
	}
	data, _ := base64.StdEncoding.DecodeString(adt01Smp)
	ending := []byte{0x0d, 0x1c, 0x0d}
	start := []byte{0x0b}
	data = append(data[:len(data)-3], ending...)
	data = append(start[:], data[:]...)
	log.Println(data)

	if _, err = conn.Write(data); err != nil {
		return nil, err
	}

	if err := conn.CloseWrite(); err != nil {
		return nil, err
	}

	redat, err := ioutil.ReadAll(conn)
	if err != nil {
		return nil, err
	}

	log.Println(redat)
	return redat, nil
}

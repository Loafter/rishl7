package rishl7

import "net"
import "io/ioutil"
import "strconv"
import "log"

const adt01Smp = "TVNIfF5+XCZ8QUtHVU5ISVN8UkFEWU9MT0rEsHxHRV9SSVN8fDIwMTUwNDI4MTA0MjQ0fHxPUk1eTzAxfDF8UHwyLjN8fHx8fHx8DQpFVk58TzAxfDIwMTUwNDI4DQpQSUR8fHxMMDQzNzk0MV5eXkhCWVN8fEtVUlRVTFVTXlRBTElQfHwxOTM3MDQyMHxNfHx8Xl5efHx8fHx8fHxLVVJVTV9LT0RVfHx8fHx8fHx8fHwNClBWMXx8T3x8fHx8fHx8fHx8fHx8fHx8QTE1MDM0NDY0MHxLVVJVTV9LT0RVXkdUQVJJSHx8fHx8fHx8fHx8fHx8fHx8fHx8fHx8fHx8fHx8fHxWfA0KT1JDfE5XfDI2NzI0NDk3fHx8U0N8fDIwMTUwNDI4MTA0MjQ0fHwyMDE1MDQyODEwNDI0NHxeXnx8Xl58fHx8fHx8DQpPQlJ8MXwyNjcyNDQ5N3x8ODAwMDFeQWtjacSfZXIgZ3JhZmlzaSBQLkEuICh0ZWsgecO2bikgLSBBcmthIMOWbl58fHwyMDE1MDQyODEwNDI0NHx8fHx8fHx8fHx8fHx8fHx8Q1J8fHxeXl5eXnx8fHx8fHx8fHx8fHx8fHx8Xl58IA0K"

type HL7Client struct {
}

func (dc *HL7Client) SendHL7Data(data []byte, hl7cd Hl7ConD) ([]byte, error) {
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

	ending := []byte{0x0d, 0x1c, 0x0d}
	start := []byte{0x0b}
	data = append(data[:len(data)-3], ending...)
	data = append(start[:], data[:]...)
	log.Println(string(data[:]))

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

	log.Println(string(redat[:]))
	return redat, nil
}

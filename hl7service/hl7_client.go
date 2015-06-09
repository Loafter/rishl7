package rishl7

import "net"
import "io/ioutil"
import "strconv"
import "log"
import "bytes"

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

	start := []byte{0x0b}
	if _, err = conn.Write(start); err != nil {
		return nil, err
	}

	stbts := bytes.Split(data, []byte{'\n'})
	for _, i := range stbts {
		i=append(i, 0x0d)
		if _, err = conn.Write(i); err != nil {
			return nil, err
		}
		log.Println(string(i[:]))

	}
	ending := []byte{0x1c, 0x0d}
	if _, err = conn.Write(ending); err != nil {
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

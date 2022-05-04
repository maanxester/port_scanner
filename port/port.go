package port

import (
	"net"
	"strconv"
	"time"
)

type ScanResult struct {
	Port  string
	State string
}

func ScanPort(protocol, hostname string, port int) ScanResult {
	result := ScanResult{Port: protocol + "/" + strconv.Itoa(port)}
	address := hostname + ":" + strconv.Itoa(port)

	// Por conta de ser um port scanner, será necessário
	// resolver várias portas no IP alvo, por isso,
	// é necessário a função DialTimeout para ter um tempo de resolução
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		result.State = "Closed"
		return result
	}
	// Instrução utilizada para finalizar conexões, limpar recursos etc, pois,
	// qualquer função precedida ao defer, não é invocado até o final da função que se encontra o defer.
	defer conn.Close()

	result.State = "Open"
	return result
}

func InitialScan(hostname string) []ScanResult {
	var results []ScanResult

	for i := 0; i <= 1024; i++ {
		results = append(results, ScanPort("tcp", hostname, i))
		results = append(results, ScanPort("udp", hostname, i))
	}

	return results
}

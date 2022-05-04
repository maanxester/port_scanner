package port

import (
	"net"
	"strconv"
	"time"
)

func ScanPort(protocol, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)

	// Por conta de ser um port scanner, será necessário
	// resolver várias portas no IP alvo, por isso,
	// é necessário a função DialTimeout para ter um tempo de resolução
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		return false
	}
	// Instrução utilizada para finalizar conexões, limpar recursos etc, pois,
	// qualquer função precedida ao defer, não é invocado até o final da função que se encontra o defer.
	defer conn.Close()

	return true
}

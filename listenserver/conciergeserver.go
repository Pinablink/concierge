package listenserver

import (
	"fmt"
	"net"

	"github.com/Pinablink/mTools/mmessage"
	"github.com/Pinablink/mTools/mtime"
)

//
type ConciergeServer struct {
	protocolo string
	ip        string
}

//
func NewConciergeServer(refProtocolo string, refIp string) *ConciergeServer {
	return &ConciergeServer{protocolo: refProtocolo,
		ip: refIp}
}

//
func (ref *ConciergeServer) InitListen() net.Listener {
	fmt.Println(mmessage.FormatMSG(mtime.DateTimeCurrentDefault(), "Solicitação de escuta no Servidor > ", ref.ip))
	list, err := net.Listen(ref.protocolo, ref.ip)

	if err != nil {
		fmt.Println(mmessage.FormatMSG(mtime.DateTimeCurrentDefault(), "Problema ao abrir escuta no servidor > ", ref.ip))
		panic(err)
	}

	fmt.Println(mmessage.FormatMSG(mtime.DateTimeCurrentDefault(), "Solicitação atendida escutando em: ", ref.ip))
	return list
}

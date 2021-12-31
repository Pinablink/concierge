package main

import (
	"concierge/grpcserver"
	"concierge/implcontract"
	"concierge/listenserver"
	"fmt"
	"net"

	"github.com/Pinablink/lingue"
	"github.com/Pinablink/lingue/oslabel"
	"github.com/Pinablink/mTools/mmessage"
	"github.com/Pinablink/mTools/mtime"
)

func main() {

	var refConciergeServer *listenserver.ConciergeServer
	var refImplContract *implcontract.ConciergeImplContract
	var refGrpc *grpcserver.ConciergeGrpcServer

	lingue.NewLingue().ExecCommand(oslabel.CLEAR_CMD)

	fmt.Println("")
	fmt.Println("  << Concierge >>  ")
	fmt.Println("")
	fmt.Println(mmessage.FormatMSG(mtime.DateTimeCurrentDefault(), "Sistema Inicializado", ""))

	refConciergeServer = listenserver.NewConciergeServer("tcp", "xxxxxxxxxxxxxxx")

	var listen net.Listener = refConciergeServer.InitListen()

	refImplContract = implcontract.NewConciergeImplContract()
	refGrpc = grpcserver.NewConciergeServer(listen, refImplContract)
	refGrpc.RegisterAndServe()

}

package grpcserver

import (
	"fmt"
	"net"

	"github.com/Pinablink/mTools/mmessage"
	"github.com/Pinablink/mTools/mtime"
	"github.com/Pinablink/proto/conciergef"
	"google.golang.org/grpc"
)

//
type ConciergeGrpcServer struct {
	iListen  net.Listener
	contract conciergef.ConciergefServer
}

//
func NewConciergeServer(refListen net.Listener, refContract conciergef.ConciergefServer) *ConciergeGrpcServer {
	return &ConciergeGrpcServer{iListen: refListen,
		contract: refContract}
}

//
func (ref *ConciergeGrpcServer) RegisterAndServe() {
	grpcServer := grpc.NewServer()
	conciergef.RegisterConciergefServer(grpcServer, ref.contract)

	error := grpcServer.Serve(ref.iListen)

	if error != nil {
		fmt.Println(mmessage.FormatMSG(mtime.DateTimeCurrentDefault(), "Erro na inicialização do server grpc", ""))
		panic(error)
	}
}

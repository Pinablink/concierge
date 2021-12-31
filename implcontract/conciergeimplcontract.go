package implcontract

import (
	"context"
	"fmt"
	"g2ssms/send"
	"os"

	"github.com/Pinablink/mTools/mmessage"
	"github.com/Pinablink/mTools/mtime"
	"github.com/Pinablink/proto/conciergef"
)

var refResultSys *conciergef.Result

//
type ConciergeImplContract struct {
	conciergef.UnimplementedConciergefServer
}

//
func NewConciergeImplContract() *ConciergeImplContract {
	refResultSys = &conciergef.Result{}
	return &ConciergeImplContract{}
}

func fSendSMS(msg string) {
	smsMsg := "%s %s %s %s"

	obSendSMS := &send.SSendSMS{}
	obSendSMS.UrlService = os.Getenv("SSMS_URL_SERVICE")
	obSendSMS.Acao = send.SendSms
	obSendSMS.Login = os.Getenv("SSMS_LOGIN_SERVICE")
	obSendSMS.Token = os.Getenv("SSMS_TOKEN_SERVICE")
	obSendSMS.Numero = os.Getenv("SSMS_NUM_DEST_SERVICE")
	obSendSMS.Msg = send.SMsg{
		Msg: fmt.Sprintf(smsMsg, msg, mtime.DateTimeCurrentDefault(), "OK", "OK"),
	}

	idResponse, err := obSendSMS.SendSMS()

	if err != nil {
		fmt.Println(mmessage.FormatMSG(mtime.DateTimeCurrentDefault(), "Problemas no envio de SMS", ""))
	}

	fmt.Println(mmessage.FormatMSG(mtime.DateTimeCurrentDefault(), "SMS enviado...", idResponse))

}

//
func (*ConciergeImplContract) CallTest(ctx context.Context, req *conciergef.Hello) (*conciergef.Hi, error) {

	fmt.Println(mmessage.FormatMSG(mtime.DateTimeCurrentDefault(), "Cumprimento recebido ", ""))
	fSendSMS("Conciergef - CallTest Processado: ")

	return &conciergef.Hi{
		Status: req.Status,
	}, nil

}

//
func (*ConciergeImplContract) CallSearch(ctx context.Context, req *conciergef.Search) (*conciergef.Result, error) {
	fSendSMS("Conciergef - CallSearch Processado: ")
	return refResultSys, nil
}

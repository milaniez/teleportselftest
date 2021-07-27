package clientjob

import (
	"github.com/milaniez/teleportselftest/util/mtls"
	"google.golang.org/grpc"
)

type Handle struct {
	conn *grpc.ClientConn
}

func HandleNew(clientCert, ClientKey, caDir, addr string) (*Handle, error) {
	creds, err := mtls.GetTLSCreds(clientCert, ClientKey, caDir, false)
	if err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}
	return &Handle{
		conn: conn,
	}, nil
}

func HandleDel(h *Handle) error {
	if err := h.conn.Close(); err != nil {
		return err
	}
	return nil
}
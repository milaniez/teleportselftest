package clientjob

import (
	"github.com/milaniez/teleportselftest/util/mtls"
	"google.golang.org/grpc"
)

type ClientJobHandle struct {
	conn *grpc.ClientConn
}

func ClientJobHandleNew(clientCert, ClientKey, caDir, addr string) (*ClientJobHandle, error) {
	creds, err := mtls.GetTLSCreds(clientCert, ClientKey, caDir, false)
	if err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}
	return &ClientJobHandle{
		conn: conn,
	}, nil
}

func ClientJobHandleDel(h *ClientJobHandle) error {
	if err := h.conn.Close(); err != nil {
		return err
	}
	return nil
}
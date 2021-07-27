package clientjob

import (
	"context"
	"github.com/milaniez/teleportselftest/protojob"
	"github.com/milaniez/teleportselftest/util/mtls"
	"google.golang.org/grpc"
	"io"
)

type Handle struct {
	conn *grpc.ClientConn
	client protojob.WorkerClient
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
		client: protojob.NewWorkerClient(conn),
	}, nil
}

func HandleDel(h *Handle) error {
	return h.conn.Close()
}

func (h *Handle) Start(cmd *protojob.Cmd) (*protojob.JobID, error) {
	return h.client.Start(context.Background(), cmd)
}

func (h *Handle) GetStatus(id *protojob.JobID) (*protojob.Status, error) {
	return h.client.GetStatus(context.Background(), id)
}

func (h *Handle) GetJobIDs() (*protojob.JobIDList, error) {
	return h.client.GetJobIDs(context.Background(), &(protojob.Empty{}))
}

func (h *Handle) GetResult(id *protojob.JobID) (*protojob.Result, error) {
	return h.client.GetResult(context.Background(), id)
}

func (h *Handle) StreamOutput(id *protojob.JobID) (chan *protojob.Output, chan error) {
	outputChan := make(chan *protojob.Output)
	errChan := make(chan error)
	go func() {
		defer close(errChan)
		defer close(outputChan)
		stream, err := h.client.StreamOutput(context.Background(), id)
		if err != nil {
			errChan <- err
			return
		}
		for {
			output, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				errChan <- err
				return
			}
			outputChan <- output
		}
	}()
	return outputChan, errChan
}
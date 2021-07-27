package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/milaniez/teleportselftest/protojob"
	"github.com/milaniez/teleportselftest/util/mtls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
	"log"
	"math/rand"
	"net"
	"time"
)

type WorkerServer struct {
	protojob.UnimplementedWorkerServer
}

func GetCertSerialFromCtx(ctx context.Context) (string, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return "", errors.New("issue getting client peer")
	}
	tlsInfo := p.AuthInfo.(credentials.TLSInfo)
	return fmt.Sprintf("%v", tlsInfo.State.VerifiedChains[0][0].SerialNumber.Text(16)), nil
}

func (w WorkerServer) Start(ctx context.Context, cmd *protojob.Cmd) (*protojob.JobID, error) {
	// TODO replace with the real implementation
	serial, err := GetCertSerialFromCtx(ctx)
	if err != nil {
		serial = "unknown"
	}
	fmt.Printf("(%v), Got Start request: name = %v, args = %v\n", serial, cmd.Name, cmd.Args)
	return &protojob.JobID{Id: rand.Uint64()}, nil
}

func (w WorkerServer) GetStatus(ctx context.Context, id *protojob.JobID) (*protojob.Status, error) {
	// TODO: replace with the real implementation
	serial, err := GetCertSerialFromCtx(ctx)
	if err != nil {
		serial = "unknown"
	}
	fmt.Printf("(%v), Got GetStatus request: id = %v\n", serial, id.Id)
	return &protojob.Status{RunStatus: protojob.Status_RUN_STATUS_FINISHED, ExitStatus: protojob.Status_EXIT_STATUS_OK}, nil
}

func (w WorkerServer) GetJobIDs(ctx context.Context, _ *protojob.Empty) (*protojob.JobIDList, error) {
	// TODO: replace with the real implementation
	serial, err := GetCertSerialFromCtx(ctx)
	if err != nil {
		serial = "unknown"
	}
	fmt.Printf("(%v), Got GetJobIDs request\n", serial)
	jobIDCnt := rand.Intn(10) + 3
	ret := protojob.JobIDList{
		Id: make([]uint64, jobIDCnt, jobIDCnt),
	}
	for i := range ret.Id {
		ret.Id[i] = rand.Uint64()
	}
	return &ret, nil
}

func (w WorkerServer) GetResult(ctx context.Context, id *protojob.JobID) (*protojob.Result, error) {
	// TODO: replace with the real implementation
	serial, err := GetCertSerialFromCtx(ctx)
	if err != nil {
		serial = "unknown"
	}
	fmt.Printf("(%v), Got GetResult request: id = %v\n", serial, id.Id)
	return &protojob.Result{Finished: true, Output: "Done!!"}, nil
}

func (w WorkerServer) StreamOutput(id *protojob.JobID, stream protojob.Worker_StreamOutputServer) error {
	// TODO: replace with the real implementation
	serial, err := GetCertSerialFromCtx(stream.Context())
	if err != nil {
		serial = "unknown"
	}
	fmt.Printf("(%v), Got StreamOutput request: id = %v\n", serial, id.Id)
	outputCnt := rand.Intn(20)
	for i := 0; i < outputCnt; i++  {
		if err := stream.Send(&protojob.Output{Output: "Something is done!"}); err == nil {
			return errors.New("issue in sending to stream")
		}
	}
	return nil
}

func WorkerServerNew() WorkerServer {
	return WorkerServer{}
}

func main() {
	var (
		certFile = flag.String(
			"cert",
			"/home/mehdi/mzzcerts/localcerts/server.crt",
			"Server certificate public key file",
		)
		keyFile  = flag.String(
			"key",
			"/home/mehdi/mzzcerts/localkeys/server.key",
			"Server certificate secret key file",
		)
		caDir    = flag.String(
			"cadir",
			"/home/mehdi/mzzcerts/cacerts/",
			"Certificate Authority directory",
		)
		addr     = flag.String(
			"address",
			"localhost:8443",
			"Server address",
		)
	)
	flag.Parse()
	log.SetFlags(log.Llongfile | log.LstdFlags)
	rand.Seed(time.Now().UTC().UnixNano())
	listener, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	creds, err := mtls.GetTLSCreds(*certFile, *keyFile, *caDir, true)
	if err != nil {
		log.Fatalf("issue getting creds: %v", err)
	}
	serverOpt := grpc.Creds(creds)
	grpcServer := grpc.NewServer(serverOpt)
	protojob.RegisterWorkerServer(grpcServer, WorkerServerNew())
	reflection.Register(grpcServer)
	log.Fatal(grpcServer.Serve(listener))
}

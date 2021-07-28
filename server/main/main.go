package main

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"github.com/milaniez/teleportselftest/protojob"
	"github.com/milaniez/teleportselftest/util/mytls"
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

func GetCertFPFromCtx(ctx context.Context) *[]string {
	ret := make([]string, 0, 0)
	p, ok := peer.FromContext(ctx)
	if !ok {
		return &ret
	}
	tlsInfo := p.AuthInfo.(credentials.TLSInfo)
	ret = make([]string, 0, len(tlsInfo.State.VerifiedChains))
	for _, chain := range tlsInfo.State.VerifiedChains {
		sha512Val := sha512.Sum512(chain[0].Raw)
		ret = append(ret, hex.EncodeToString(sha512Val[:]))
	}
	return &ret
}

func (w WorkerServer) Start(ctx context.Context, cmd *protojob.Cmd) (*protojob.JobID, error) {
	// TODO replace with the real implementation
	fp := GetCertFPFromCtx(ctx)
	fmt.Printf("(%v), Got Start request: name = %v, args = %v\n", *fp, cmd.Name, cmd.Args)
	return &protojob.JobID{Id: rand.Uint64()}, nil
}

func (w WorkerServer) GetStatus(ctx context.Context, id *protojob.JobID) (*protojob.Status, error) {
	// TODO: replace with the real implementation
	fp := GetCertFPFromCtx(ctx)
	fmt.Printf("(%v), Got GetStatus request: id = %v\n", *fp, id.Id)
	return &protojob.Status{RunStatus: protojob.Status_RUN_STATUS_FINISHED, ExitStatus: protojob.Status_EXIT_STATUS_OK}, nil
}

func (w WorkerServer) GetJobIDs(ctx context.Context, _ *protojob.Empty) (*protojob.JobIDList, error) {
	// TODO: replace with the real implementation
	fp := GetCertFPFromCtx(ctx)
	fmt.Printf("(%v), Got GetJobIDs request\n", *fp)
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
	fp := GetCertFPFromCtx(ctx)
	fmt.Printf("(%v), Got GetResult request: id = %v\n", *fp, id.Id)
	return &protojob.Result{Finished: true, Output: "Done!!"}, nil
}

func (w WorkerServer) StreamOutput(id *protojob.JobID, stream protojob.Worker_StreamOutputServer) error {
	// TODO: replace with the real implementation
	fp := GetCertFPFromCtx(stream.Context())
	fmt.Printf("(%v), Got StreamOutput request: id = %v\n", *fp, id.Id)
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

func Intercept(
	ctx context.Context,
	req interface{},
	_ *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	fmt.Println("called")
	if p, ok := peer.FromContext(ctx); ok {
		if tlsInfo, ok := p.AuthInfo.(credentials.TLSInfo); ok {
			for _, item := range tlsInfo.State.PeerCertificates {
				fmt.Println("request certificate subject:", item.Subject)
			}
		}
	}
	return handler(ctx, req)
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
	creds, err := mytls.GetTLSCreds(*certFile, *keyFile, *caDir, true)
	if err != nil {
		log.Fatalf("issue getting creds: %v", err)
	}
	serverOpt := grpc.Creds(creds)
	grpcServer := grpc.NewServer(serverOpt, grpc.UnaryInterceptor(Intercept))
	protojob.RegisterWorkerServer(grpcServer, WorkerServerNew())
	reflection.Register(grpcServer)
	log.Fatal(grpcServer.Serve(listener))
}

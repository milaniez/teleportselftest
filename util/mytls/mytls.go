package mytls

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"strings"
)

func GetTLSCreds(certFile, keyFile, caDir string, isServer bool) (credentials.TransportCredentials, error) {
	if !strings.HasSuffix(caDir, "/") {
		caDir += "/"
	}
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}
	caCertPool := x509.NewCertPool()
	caCertPoolMemCnt := 0
	caDirFiles, err := ioutil.ReadDir(caDir)
	if err != nil {
		return nil, err
	}
	for _, entry := range caDirFiles {
		if entry.IsDir() || !(strings.HasSuffix(entry.Name(), ".pem") ||
			strings.HasSuffix(entry.Name(), ".crt")) {
			continue
		}
		caFilePath := caDir + entry.Name()
		caFile, err := ioutil.ReadFile(caFilePath)
		if err != nil {
			return nil, err
		}
		if !caCertPool.AppendCertsFromPEM(caFile) {
			log.Printf("Warning: Certificate %v not added.\n", caFilePath)
		} else {
			caCertPoolMemCnt++
		}
	}
	if caCertPoolMemCnt == 0 {
		return nil, errors.New("no CA certificate")
	}
	tlsConfig := &tls.Config{
		ClientCAs:  caCertPool,
		VerifyPeerCertificate: func(_ [][]byte, verifiedChains [][]*x509.Certificate) error {
			for _, chain := range verifiedChains {
				if len(chain) > 2 || len(chain) == 0 {
					continue
				}
				return nil
			}
			return errors.New("invalid certificate chains")
		},
		Certificates: []tls.Certificate{cert},
	}
	if isServer {
		tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
		tlsConfig.ClientCAs = caCertPool
	} else {
		tlsConfig.RootCAs = caCertPool
	}
	return credentials.NewTLS(tlsConfig), nil
}
/*
Copyright © 2023 Caner Emeç caner.emec@gmail.com

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package connection

import (
	"crypto/x509"
	"fmt"
	"os"
	"time"

	c "github.com/caner-emec/chaincode-stalker/internal/types"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	evaluateTimeout     = 5 * time.Second
	endorseTimeout      = 15 * time.Second
	submitTimeout       = 5 * time.Second
	commitStatusTimeout = 1 * time.Minute
)

func NewConnection(config *c.Config) (*client.Gateway, *grpc.ClientConn, error) {
	// The gRPC client connection should be shared by all Gateway connections to this endpoint
	clientConnection := newGrpcConnection(config.Connection)

	id := newIdentity(config.Identity)
	sign := newSign(config.Identity)

	// Create a Gateway connection for a specific client identity
	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(evaluateTimeout),
		client.WithEndorseTimeout(endorseTimeout),
		client.WithSubmitTimeout(submitTimeout),
		client.WithCommitStatusTimeout(commitStatusTimeout),
	)
	if err != nil {
		panic(err)
	}

	return gw, clientConnection, err
}

// newGrpcConnection creates a gRPC connection to the Gateway server.
func newGrpcConnection(conn c.ConnectionConfig) *grpc.ClientConn {
	certificatePEM, err := os.ReadFile(conn.TLSCert)
	if err != nil {
		panic(err)
	}

	certificate, err := identity.CertificateFromPEM(certificatePEM)
	if err != nil {
		fmt.Println(conn.TLSCert)
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, conn.GatewayPeer)

	connection, err := grpc.Dial(conn.PeerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection
}

// newIdentity creates a client identity for this Gateway connection using an X.509 certificate.
func newIdentity(conf c.IdentityConfig) *identity.X509Identity {
	certificatePEM, err := os.ReadFile(conf.Cert)
	if err != nil {
		panic(err)
	}

	certificate, err := identity.CertificateFromPEM(certificatePEM)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(conf.MspID, certificate)
	if err != nil {
		panic(err)
	}

	return id
}

// newSign creates a function that generates a digital signature from a message digest using a private key.
func newSign(conf c.IdentityConfig) identity.Sign {
	priv, err := os.ReadFile(conf.PrivateKey)
	if err != nil {
		panic(err)
	}

	privateKey, err := identity.PrivateKeyFromPEM(priv)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
}

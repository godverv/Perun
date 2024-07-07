package connections_cache

import (
	"io"
	"net"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/Red-Sock/Perun/internal/domain"
)

func getVelezKey(velezConn domain.VelezConnection) ([]byte, error) {
	signer, err := ssh.ParsePrivateKey(velezConn.Ssh.Key)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing private key")
	}

	clientConfig := &ssh.ClientConfig{
		User: velezConn.Ssh.Username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	tcpConn, err := ssh.Dial("tcp", velezConn.Node.Addr, clientConfig)
	if err != nil {
		return nil, errors.Wrap(err, "error connecting to "+velezConn.Node.Addr)
	}

	sshClient, err := sftp.NewClient(tcpConn)
	if err != nil {
		return nil, errors.Wrap(err, "error creating sftp client")
	}

	f, err := sshClient.Open(velezConn.Node.CustomVelezKeyPath)
	if err != nil {
		return nil, errors.Wrap(err, "error opening velez key")
	}

	defer f.Close()

	velezKey, err := io.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "error reading velez key")
	}

	return velezKey, nil
}

func getGrpcConnection(connStr string) (*grpc.ClientConn, error) {
	dial, err := grpc.NewClient(
		connStr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, errors.Wrap(err, "error dialing")
	}

	return dial, nil
}

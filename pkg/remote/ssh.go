package remote

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

type Ssh struct {
	config *ssh.ClientConfig
}

func NewSsh(user string) (*Ssh, error) {

	dirname, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("userHomeDirError: %v", err)
	}
	// A public key may be used to authenticate against the remote
	// server by using an unencrypted PEM-encoded private key file.
	//
	// If you have an encrypted private key, the crypto/x509 package
	// can be used to decrypt it.
	keyFile := dirname + "/.ssh/id_rsa"
	key, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, fmt.Errorf("readFileError: private key %v", err)
	}
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, fmt.Errorf("parsePrivateKeyError: %v", err)
	}

	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			// Use the PublicKeys method for remote authentication.
			ssh.PublicKeys(signer),
		},
		Timeout: 5 * time.Second,
	}

	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	ssh := &Ssh{
		config: sshConfig,
	}

	return ssh, nil
}

func (s *Ssh) ConnectSSH(host string) (*ssh.Client, error) {
	client, err := ssh.Dial("tcp", host, s.config)
	if err != nil {
		return nil, fmt.Errorf("sshDialError: unable to connect to %s, %v",host,err)
	}

	return client, nil
}

func SshCommand(client *ssh.Client, command string) (*bytes.Buffer, *bytes.Buffer, error) {

	session, err := client.NewSession()
	if err != nil {
		return nil, nil, fmt.Errorf("clientNewSessionError %v", session)
	}
	defer session.Close()

	// StdinPipe for commands
	// stdin, err := s.session.StdinPipe()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Uncomment to store output in variable
	var b, c bytes.Buffer

	session.Stdout = &b
	session.Stderr = &c

	// Enable system stdout
	// Comment these if you uncomment to store in variable
	// s.session.Stdout = os.Stdout
	// s.session.Stderr = os.Stderr

	// Start remote shell
	// err:= s.session.Shell()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err = session.Start(command)
	if err != nil {
		return nil, nil, fmt.Errorf("sessionStartError %v", err)
	}
	err = session.Wait()
	if err != nil {
		return nil, nil, fmt.Errorf("sessionWaitError %v, %s", err, c.String())
	}

	return &b, &c, nil
}

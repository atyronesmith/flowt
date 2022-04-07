package remote

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

type Ssh struct {
	config *ssh.ClientConfig
}

func NewSsh() *Ssh {

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	// A public key may be used to authenticate against the remote
	// server by using an unencrypted PEM-encoded private key file.
	//
	// If you have an encrypted private key, the crypto/x509 package
	// can be used to decrypt it.
	keyFile := dirname + "/.ssh/undercloud.pkey"
	fmt.Printf("private key: %s\n", keyFile)
	key, err := ioutil.ReadFile(keyFile)
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	sshConfig := &ssh.ClientConfig{
		User: "heat-admin",
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

	return ssh
}

func (s *Ssh) ConnectSSH(host string) (*ssh.Client, error) {
	fmt.Printf("Connecting to: %s\n", host)
	client, err := ssh.Dial("tcp", host, s.config)
	if err != nil {
		fmt.Printf("Unable to connect to %s, %v\n", host, err)
		return nil, err
	}

	fmt.Printf("Connected to: %s\n", host)

	return client, nil
}

func SshCommand(client *ssh.Client, commands []string) (*bytes.Buffer, error) {

	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	// StdinPipe for commands
	// stdin, err := s.session.StdinPipe()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Uncomment to store output in variable
	var b bytes.Buffer

	session.Stdout = &b
	session.Stderr = &b

	// Enable system stdout
	// Comment these if you uncomment to store in variable
	// s.session.Stdout = os.Stdout
	// s.session.Stderr = os.Stderr

	// Start remote shell
	// err:= s.session.Shell()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	for _, cmd := range commands {
		err := session.Start(cmd)
		if err != nil {
			fmt.Printf("error %v\n", err)
			log.Fatal(err)
		}
		fmt.Printf("cmd: %s\n", cmd)
		err = session.Wait()
		if err != nil {
			log.Fatal(err)
			fmt.Printf("error %v\n", err)
		}
	}
	fmt.Printf("done len %d\n", len(b.Bytes()))

	//	time.Sleep(5 * time.Second)
	//Wait for sess to finish
	// err := s.session.Wait()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return &b, nil
}

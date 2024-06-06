package main

import (
	"os"

	"golang.org/x/crypto/ssh"
)

func connectToSftpServer(host, user, pemPath string) (*ssh.Client, error) {

	// Read the private key file
	pvtkeyBytes, err := os.ReadFile(pemPath)
	if err != nil {
		panic(err)
	}

	privateKey, err := ssh.ParsePrivateKey(pvtkeyBytes)
	if err != nil {
		panic(err)
	}

	config := &ssh.ClientConfig{
		User:            user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(privateKey),
		},
	}

	conn, err := ssh.Dial("tcp", host, config)
	if err != nil {
		panic(err)
	}
	return conn, nil

}

// func connectToServer(user, pemPath string) {

// 	pvtKeyBytes, err := os.ReadFile(pemPath)
// 	if err != nil {
// 		log.Fatalf("Could not read file: ERR: %v", err)
// 	} else {
// 		fmt.Println("Read Successful!!")
// 	}

// 	pvtkey, err := ssh.ParsePrivateKey(pvtKeyBytes)
// 	if err != nil {
// 		log.Fatalf("Parse Private Key Error: ERR: %v", err)
// 	} else {
// 		fmt.Println("Parsed successfully")
// 	}

// 	config := &ssh.ClientConfig{
// 		User: user,
// 		Auth: []ssh.AuthMethod{
// 			ssh.PublicKeys(pvtkey),
// 		},
// 	}

// 	conn, err := ssh.Dial("tcp", "serverIP:22", config)
// 	if err != nil {
// 		log.Fatalf("Error, could not dial: ERR: %v", err)
// 	}
// 	defer conn.Close()

// 	sftpClient, err := sftp.NewClient(conn)
// 	if err != nil {
// 		log.Fatal("Failed to create SFTP client: ", err)
// 	}
// 	defer sftpClient.Close()
// }

package main

import (
	"fmt"
	"log"

	"github.com/pkg/sftp"
)

func main() {

	host := "IP:22"
	user := "username"
	remotePath := "./try.txt"
	localPath := "D:/SFTPConnectionFolder/newfile.txt"
	pemFilePath := "D:/testing.pem"

	// 	// connectToServer(user, pemFilePath)

	conn, err := connectToSftpServer(host, user, pemFilePath)
	if err != nil {
		log.Fatal("Failed to connect: ", err)
	} else {
		fmt.Println("success")
	}

	sftpClient, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatal("Failed to create SFTP client: ", err)
	}
	defer sftpClient.Close()

	files, err := listFiles(sftpClient)
	if err != nil {
		log.Fatal("Failed to list files: ", err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}

	err = downloadFile(sftpClient, remotePath, localPath)
	if err != nil {
		log.Fatal("Failed to download file: ", err)
	}
}

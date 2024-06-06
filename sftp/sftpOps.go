package main

import (
	"io"
	"os"

	"github.com/pkg/sftp"
)

func listFiles(sftpClient *sftp.Client) ([]os.FileInfo, error) {
	files, err := sftpClient.ReadDir(".")
	if err != nil {
		return nil, err
	}
	return files, nil
}

func downloadFile(sftpClient *sftp.Client, remoteFile string, localFile string) error {
	remote, err := sftpClient.Open(remoteFile)
	if err != nil {
		return err
	}
	defer remote.Close()

	local, err := os.Create(localFile)
	if err != nil {
		return err
	}
	defer local.Close()

	_, err = io.Copy(local, remote)
	return err
}

package components

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"os"
)

type SshInterface interface {
	connect() (string, error)
	RunCommand(cmd string)
	DownloadFile(remotePath string, localPath string)
	Close()
}

type sshConnection struct {
	conn *ssh.Client
	sshHost string
	sshUser string
	sshPassword string
	sshPort string
}

func (c *sshConnection) connect() (string, error) {
	// ssh connect
	config := &ssh.ClientConfig {
		User: c.sshUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(c.sshPassword),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	var err error
	c.conn, err = ssh.Dial("tcp", c.sshHost + ":" + c.sshPort, config)
	if err != nil {
		panic(err)
		//return "", errors.New("cant connect to SSH")
	}
	//defer c.conn.Close()

	return "ok", nil
}

func (c *sshConnection) RunCommand(cmd string)  {
	sess, err := c.conn.NewSession()
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	sessStdOut, err := sess.StdoutPipe()
	if err != nil {
		panic(err)
	}

	go io.Copy(os.Stdout, sessStdOut)
	sessStderr, err := sess.StderrPipe()
	if err != nil {
		panic(err)
	}
	go io.Copy(os.Stderr, sessStderr)
	err = sess.Run(cmd)
	if err != nil {
		panic(err)
	}
}

func (c *sshConnection) DownloadFile(remotePath string, localPath string) {
	client, err := sftp.NewClient(c.conn)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	srcFile, err := client.Open(remotePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer srcFile.Close()

	dstFile, err := os.Create(localPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	//log.Fatal(err)

	//// walk a directory
	//w := client.Walk("/tmp")
	//for w.Step() {
	//	if w.Err() != nil {
	//		continue
	//	}
	//	log.Println(w.Path())
	//}

	//// check it's there
	//fi, err := client.Lstat("/tmp/dump.sql")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(fi)

	//fmt.Println("start download")
	//client, err := scp.NewClientBySSH(c.conn)
	//if err != nil {
	//	fmt.Println("Error creating new SSH session from existing connection", err)
	//}
	//
	//_, err2 := os.Create(localPath)
	//f, err3 := os.OpenFile(localPath, os.O_WRONLY, 0666)
	//fmt.Println(err2)
	//fmt.Println(err3)
	//defer client.Close()
	//defer f.Close()
	//
	//err = client.CopyFile(f, remotePath, "0655")
	//
	//if err != nil {
	//	fmt.Println("Error while copying file ", err)
	//}
}

func (c *sshConnection) Close() {
	fmt.Println("Close connection")
	c.conn.Close()
}

// ssh connection
func GetSshConnect(sshHost string, sshUser string, sshPassword string, sshPort string) SshInterface {
	res := &sshConnection{
		sshHost: sshHost,
		sshUser: sshUser,
		sshPassword: sshPassword,
		sshPort: sshPort,
	}
	_, _ = res.connect()
	return res
}

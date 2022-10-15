package Ubuntu

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

func Startinstall() {

	// var cmd *exec.Cmd
	// var stdout []byte
	// var err error

	// apt update
	// apt install mysql-server
	// mysql_secure_installation
	// nano /etc/mysql/mysql.conf.d/mysqld.cnf
	// mysql
	// GRANT ALL PRIVILEGES ON *.* TO 'room'@'*' IDENTIFIED BY 'therif' WITH GRANT OPTION;
	// FLUSH PRIVILEGES;
	// exit

	fmt.Println("=============")
	log.Println("-- Check System Update --")
	//AsyncCmdBashSudo("apt-get update")

	fmt.Println("======== HTTP MAU PAKAI APA ? ========")
	fmt.Println("1. Nginx")
	fmt.Println("2. Apache2")
	fmt.Println("Isi Nomor Pilihan : ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	jawaban := input.Text()
	//fmt.Print(jawaban)
	if jawaban == "1" {
		fmt.Println("=============")
		log.Println("-- Install Nginx Server --")
		AsyncCmdBashSudo("apt-get install nginx -y")
	}
	if jawaban == "2" {
		fmt.Println("=============")
		log.Println("-- Install Apache2 Server --")
		AsyncCmdBashSudo("apt-get install apache2 -y")
	}

	fmt.Println("=============")
	log.Println("-- Install PHP 8.1 --")
	AsyncCmdBashSudo("apt-get install php8.1 -y")

	fmt.Println("=============")
	log.Println("-- Install MySQL Server --")
	AsyncCmdBashSudo("apt-get install mysql-server -y")

	fmt.Println("=============")
	log.Println("-- Starting MySQL Server --")
	AsyncCmdBashSudo("sudo /etc/init.d/mysql restart")

	fmt.Println("=============")
	log.Println("-- Setup MySQL Secure Installation --")
	AsyncCmdBashSudo("mysql_secure_installation")

	fmt.Println("=============")
	log.Println("-- Setup MySQL Secure Installation --")
	//AsyncCmdBashSudo("mysql -uroot -p -e \"GRANT ALL PRIVILEGES ON *.* TO 'room'@'*' IDENTIFIED BY 'therif' WITH GRANT OPTION;FLUSH PRIVILEGES;\"")
	//mysql -uroot -p -e "GRANT ALL PRIVILEGES ON *.* TO 'room'@'*' IDENTIFIED BY 'therif' WITH GRANT OPTION;FLUSH PRIVILEGES;"
	fmt.Println("-- Completed !!! --")
	//CmdBash("sudo apt-get upgrade -y")

}

// func tcp_con_handle(con exec.Cmd) {
// 	log.Println("tcp_con_handle")
// 	chan_to_stdout := stream_copy(con, os.Stdout)

// 	chan_to_remote := stream_copy(os.Stdin, con)
// 	select {
// 	case <-chan_to_stdout:
// 		log.Println("Remote connection is closed")
// 	case <-chan_to_remote:
// 		log.Println("Local program is terminated")
// 	}
// }

// Performs copy operation between streams: os and tcp streams
func stream_copy(src io.Reader, dst io.Writer) <-chan int {
	log.Println("stream_copy")
	buf := make([]byte, 1024)
	sync_channel := make(chan int)
	go func() {
		defer func() {
			if con, ok := dst.(net.Conn); ok {
				con.Close()
				log.Printf("Connection from %v is closed\n", con.RemoteAddr())
			}
			sync_channel <- 0 // Notify that processing is finished
		}()
		for {

			var nBytes int
			var err error
			nBytes, err = src.Read(buf)
			if err != nil {
				if err != io.EOF {
					log.Printf("Read error: %s\n", err)
				}
				break
			}

			// _, err = dst.Write(buf[0:nBytes])
			// if err != nil {
			// 	log.Fatalf("Write error: %s\n", err)
			// }

			log.Println(buf[0:nBytes])

		}
	}()
	return sync_channel
}

func CmdBashSudo(cmdnya string) {
	if len(strings.TrimSpace(cmdnya)) > 0 {
		cmd := exec.Command("bash", "-c", "sudo "+cmdnya)

		cmd.Stdin = os.Stdin
		//cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		stdout, err := cmd.Output()
		if err != nil {
			//fmt.Println("Err", err)
			fmt.Fprintln(os.Stderr, err)
			return
		}
		fmt.Println(string(stdout))
	}
}

func AsyncCmdBashSudo(cmdnya string) {
	cmd := exec.Command("bash", "-c", "sudo "+cmdnya)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmdReader, _ := cmd.StdoutPipe()

	scanner := bufio.NewScanner(cmdReader)
	done := make(chan bool)
	go func() {
		for scanner.Scan() {
			fmt.Print(scanner.Text())
		}
		done <- true
	}()
	cmd.Start()
	<-done
	_ = cmd.Wait()

}

func CmdBash(cmdnya string) {
	if len(strings.TrimSpace(cmdnya)) > 0 {
		cmd := exec.Command("bash", "-c", cmdnya)

		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println("Err", err)
			return
		}
		fmt.Println(string(stdout))
	}
}

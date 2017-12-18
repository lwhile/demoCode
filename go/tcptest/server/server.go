package server

import (
	"log"
	"net"
	"time"
)

// // StartServer start a server acvititu
// func StartServer() {
// 	ltn, err := net.Listen("tcp", ":8888")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	defer ltn.Close()
// 	log.Println("Listen ok")

// 	var c int
// 	for {
// 		time.Sleep(10 * time.Second)
// 		if _, err := ltn.Accept(); err != nil {
// 			log.Println("Accept error:", err)
// 			break
// 		}
// 		c++
// 		log.Printf("#%d: accept a new connection\n", c)
// 	}
// }

// ReadServer read data from a tcp connect
func ReadServer() {
	ltn, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	defer ltn.Close()
	conn, err := ltn.Accept()
	defer conn.Close()
	for {
		var buf = make([]byte, 10)
		log.Println("start to read 	from conn")
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Read %d bytes,content is %s\n", n, string(buf))
	}
}

// ReadServerFull read data from tcp read buffer
func ReadServerFull() {
	ltn, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	defer ltn.Close()

	conn, err := ltn.Accept()
	if err != nil {
		log.Fatal("accept err:", err)
	}

	for {
		//read from the connection
		time.Sleep(5 * time.Second)

		var buf = make([]byte, 60000)
		log.Println("start to read from conn")
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("conn read %d byte, error: %s\n", n, err)
			if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
				continue
			}
		}
		log.Printf("read %d byte,content is %s\n", n, string(buf[:n]))
	}
}

// CloseServer test close tcp connect
func CloseServer() {
	ltn, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("ltn err:", err)
	}
	defer ltn.Close()
	conn, err := ltn.Accept()
	if err != nil {
		log.Fatal("accpet err:", err)
	}
	log.Println("Accept a new conn")
	defer conn.Close()

	var buf = make([]byte, 10)
	log.Println("start to read from conn")
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("conn read error:", err)
	} else {
		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	}

	n, err = conn.Write(buf)
	if err != nil {
		log.Println("conn write error:", err)
	} else {
		log.Printf("write %d bytes, content is %s\n", n, string(buf[:n]))
	}
}

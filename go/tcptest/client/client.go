package client

import (
	"log"
	"net"
	"time"
)

// func establishConn(i int) net.Conn {
// 	conn, err := net.Dial("tcp", ":8888")
// 	if err != nil {
// 		log.Printf("#%d: dial error:%s", i, err)
// 	}
// 	log.Printf("#%d connect to server ok", i)
// 	return conn
// }

// // StartClient start client activity
// func StartClient() {
// 	var cs []net.Conn
// 	for i := 1; i < 1000; i++ {
// 		conn := establishConn(i)
// 		if conn != nil {
// 			cs = append(cs, conn)
// 		}
// 	}
// }

// WriteClient wirte data to server by a tcp connect
func WriteClient(data []byte) {
	log.Println("Begin dial...")
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	log.Println("Dial success.")

	time.Sleep(time.Second * 1)

	conn.Write(data)
}

// WriteClientFull writer data to tcp write buffer
func WriteClientFull(data []byte) {
	log.Println("Begin dial...")
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Println("dial error: ", err)
		return
	}

	conn.Close()
	log.Println("Dail ok.")

	var buf = make([]byte, 32)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal("Read error: ", err)
	} else {
		log.Printf("read %d bytes\n", n)
	}
}

// CloseClient test close tcp connect
func CloseClient() {
	log.Println("Begin dial...")

	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Println(err)
	}
	// 这里提前关闭tcp连接
	conn.Close()
	log.Println("close ok")

	var buf = make([]byte, 32)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("read error", err)
	} else {
		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	}

	n, err = conn.Write(buf)
	if err != nil {
		log.Println("write error:", err)
	} else {
		log.Printf("write %d bytes, content is %s\n", n, string(buf[:n]))
	}

	time.Sleep(1000 * time.Second)
}

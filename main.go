package main

import (
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error listening: %s", err)
	}
	log.Printf("Listening on %s", l.Addr())
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("Error accepting: %s", err)
		}
		log.Printf("Accepted connection from %s", conn.RemoteAddr())
		us, err := net.Dial("tcp", "127.0.0.1:4444")
		if err != nil {
			log.Fatalf("Error dialing: %s", err)
		}
		// go io.Copy(conn, us)
		// io.Copy(us, conn)
		go func(conn net.Conn, us net.Conn) {
			var b [1024]byte
			n, err := conn.Read(b[:])
			if err != nil {
				log.Fatalf("Error reading: %s", err)
			}
			log.Printf("Read %d bytes", n)
			n, err = us.Write(b[:n])
			if err != nil {
				log.Fatalf("Error writing: %s", err)
			}
			log.Printf("Wrote %d bytes", n)
		}(conn, us)
		go func(conn net.Conn, us net.Conn) {
			var b [1024]byte
			n, err := us.Read(b[:])
			if err != nil {
				log.Fatalf("Error reading: %s", err)
			}
			log.Printf("Read %d bytes", n)
			n, err = conn.Write(b[:n])
			if err != nil {
				log.Fatalf("Error writing: %s", err)
			}
			log.Printf("Wrote %d bytes", n)
		}(conn, us)
	}

}

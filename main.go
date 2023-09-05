package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
_	"time"
)

func main() {
	li, err := net.Listen("tcp",":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
	respond(conn)
}
func request(conn net.Conn) {

	var lnCounter int

	// if err := conn.SetDeadline(time.Now().Add(10 * time.Second)); err != nil {
	// 	log.Println("CONN TIMEOUT")
	// }

	//READ FROM TCP CONNECTION
	scanner := bufio.NewScanner(conn)
	
	lnCounter = 0
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if lnCounter == 0 {
			m := strings.Fields(ln)
			fmt.Println("---METHOD---URL---", m[0], m[1])

		}
		if ln == "" { break }
		lnCounter++

	}
}

func respond(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World!</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
// err := conn.SetDeadline(time.Now().Add(10 * time.Second))
// if err != nil {
// 	log.Println("CONN TIMEOUT")
// }
// scanner := bufio.NewScanner(conn)
// for scanner.Scan() {
// 	ln := scanner.Text()
// 	fmt.Println(ln)
// 	fmt.Fprintf(conn, "I heard you say: %s\n", ln)
// }
// defer conn.Close()

// fmt.Println("Code got here")
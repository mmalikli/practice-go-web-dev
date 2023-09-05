package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net"
	_ "net/url"
	"strings"
	_ "time"
)

func main() {
  
	tpl := template.Must(template.ParseFiles("tpl.gohtml"))

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
		go handle(conn, tpl)
	}
}

func handle(conn net.Conn, tpl *template.Template) {
	defer conn.Close()
	var url string
	request(conn, &url)
	respond(conn, &url, tpl)
}
func request(conn net.Conn, url *string) {
	var lnCounter int

	//READ FROM TCP CONNECTION
	scanner := bufio.NewScanner(conn)
	
	lnCounter = 0
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if lnCounter == 0 {
			*url = strings.Fields(ln)[1]
		}
		if ln == "" { break }
		lnCounter++

	}
}

func respond(conn net.Conn, url *string, tpl *template.Template) { 
	//body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World!</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	//fmt.Fprintf(conn, "Content-Length: %d\r\n")
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	tpl.Execute(conn,*url)
	//fmt.Printf("\n---%s---\n",*url)
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
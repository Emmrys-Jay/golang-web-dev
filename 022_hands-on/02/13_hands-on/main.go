package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	lstner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer lstner.Close()

	for {
		conn, err := lstner.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
			//remember to put continue so the current iteration is skipped
			//therefore the line where a string is being wriiten to conn wont run
		}
		go serve(conn)

	}
}

func serve(conn net.Conn) {
	defer conn.Close()

	response(conn)

	request(conn)
}

func response(conn net.Conn) {
	body := `<!DOCTYPE html><html><head><meta charset="UTF-8"><title>Response</title></head><body>You connected succesfully</body></html>`

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)

}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			s := strings.Fields(ln)
			fmt.Printf("***METHOD: %s\n", s[0])
			fmt.Printf("***URI: %s\n", s[1])
		}
		if ln == "" {
			break
		}
		i++
	}
}

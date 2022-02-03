package main

import (
	"bufio"
	"fmt"
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
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	retURL := request(conn)

	response(conn, retURL)
}

func request(conn net.Conn) string {
	i := 0
	scanner := bufio.NewScanner(conn)
	var urlStr string
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			m := strings.Fields(ln)[0]
			urlStr = strings.Fields(ln)[1]
			fmt.Println("***METHOD", m)
			fmt.Println("***URL", urlStr)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
	return urlStr
}

func response(conn net.Conn, urlStr string) {

	uri := fmt.Sprintf("<div>Uniform resource locator (URL): %s</div>", urlStr)
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong>` + uri + `</body></html`
	//io.WriteString(conn)

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}

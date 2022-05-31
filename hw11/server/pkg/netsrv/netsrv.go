package netsrv

import (
	"bufio"
	"fmt"
	"go-search/hw11/server/pkg/crawler"
	"log"
	"net"
	"strings"
)

type server struct {
	listener net.Listener
	pages    []crawler.Document
}

type Interface interface {
	Start([]crawler.Document)
}

func New() *server {
	s := server{}

	listener, err := net.Listen("tcp4", ":8000")
	if err != nil {
		log.Println(err)
		return &s
	}

	s.listener = listener
	return &s
}

func (s *server) Start(data []crawler.Document) {
	s.pages = data
	log.Println("Ready for clients")

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		go s.handler(conn)
	}
}

func (s *server) handler(conn net.Conn) {
	defer conn.Close()

	for {
		res, err := readUserInput(conn)
		if err != nil {
			return
		}

		pages := search(res, s.pages)
		sendsResponsesToUser(conn, []string{"Results for: ", res})
		sendsResponsesToUser(conn, pages)
	}
}

func readUserInput(conn net.Conn) (req string, err error) {
	r := bufio.NewReader(conn)
	msg, _, err := r.ReadLine()
	if err != nil {
		return req, err
	}
	req = strings.ToLower(string(msg))
	return req, err
}

func sendsResponsesToUser(conn net.Conn, text []string) {
	for _, str := range text {
		_, err := conn.Write([]byte(str))
		if err != nil {
			return
		}
	}
	_, err := conn.Write([]byte("\n"))
	if err != nil {
		return
	}
}

func search(req string, pages []crawler.Document) (res []string) {
	for _, p := range pages {
		if strings.Contains(strings.ToLower(p.Title), req) {
			res = append(res, fmt.Sprintf("Document: '%s' (%s)\n", p.Title, p.URL))
		}
	}

	if len(res) == 0 {
		res = append(res, "Nothing found")
	}
	return res
}

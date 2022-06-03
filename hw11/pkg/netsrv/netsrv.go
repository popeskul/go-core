package netsrv

import (
	"bufio"
	"fmt"
	"go-search/hw11/pkg/crawler"
	"log"
	"net"
	"strings"
	"time"
)

type server struct {
	listener net.Listener
	pages    []crawler.Document
}

type Interface interface {
	New([]crawler.Document)
	Start()
}

func New(data []crawler.Document) *server {
	s := server{
		pages: data,
	}

	listener, err := net.Listen("tcp4", "localhost:8000")
	if err != nil {
		log.Println(err)
		return &s
	}

	s.listener = listener
	return &s
}

func (s *server) Start() error {
	log.Println("Ready for clients")

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return err
		}

		err = conn.SetDeadline(time.Now().Add(time.Minute * 5))
		if err != nil {
			return err
		}

		if err != nil {
			return err
		}
		go s.handler(conn)
	}
}

func (s *server) handler(conn net.Conn) {
	defer conn.Close()

	for {
		res, err := readUserInput(conn)
		if err != nil {
			log.Println(err)
			return
		}

		pages := search(res, s.pages)
		err = sendsResponsesToUser(conn, []string{"Results for: ", res})
		if err != nil {
			log.Println(err)
			return
		}
		err = sendsResponsesToUser(conn, pages)
		if err != nil {
			log.Println(err)
			return
		}
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

func sendsResponsesToUser(conn net.Conn, text []string) error {
	for _, str := range text {
		_, err := conn.Write([]byte(str))
		if err != nil {
			return err
		}
	}
	_, err := conn.Write([]byte("\n"))
	if err != nil {
		return err
	}

	return nil
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

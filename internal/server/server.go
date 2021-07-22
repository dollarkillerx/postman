package server

import (
	"github.com/dollarkillerx/erguotou"

	"log"
	"sync"
	"time"
)

type server struct {
	app *erguotou.Engine

	wxMu    sync.RWMutex
	wxToken string
}

func NewServer() *server {
	return &server{}
}

func (s *server) Run(addr string) error {
	s.app = erguotou.New()

	s.app.Use(erguotou.Logger)
	s.router()

	go s.updateWechatToken()
	time.Sleep(time.Millisecond * 300)

	log.Println("Postman Run: ", addr)
	if err := s.app.Run(erguotou.SetHost(addr)); err != nil {
		return err
	}

	return nil
}

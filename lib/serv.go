package lib

import (
        "errors"
        "net/http"

        "golang.org/x/crypto/openpgp"
)

func Serve(ip []byte, port uint) error {
        return nil
}

type ctrl struct {
        pkt *Proto
        reply chan<- *Proto
}

type message struct {
        from *openpgp.Entity
        to *openpgp.Entity
        data []byte
}

type server struct {
        whitelist []string

        ents []*openpgp.Entity

        msgs []*message
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func (s *server) loop(stream <-chan ctrl) error {
        for input := range stream {
                out, err := s.handle(input.pkt)

                if err != nil {
                        return err
                }

                input.reply<- out
        }

        return nil
}

func (s *server) handle(input *Proto) (*Proto, error) {
        switch input.Cmd {
        case Download:
                return s.send(input)
        case Upload:
                return s.recv(input)
        default:
                return nil, errors.New("Unknown  Proto command")
        }
}

func (s *server) send(input *Proto) (*Proto, error) {
        return nil, nil
}

func (s *server) recv(input *Proto) (*Proto, error) {
        return nil, nil
}

package lib

import (
        "fmt"
        "net/http"

        "github.com/pkg/errors"
        "github.com/gorilla/mux"
)

func Serve(ip []byte, port uint, auth Auth) error {
        s := &server{}
        s.stream = make(chan ctrl, 1)
        s.auth = auth

        r := mux.NewRouter()

        r.HandleFunc("/{to}", s.out).Methods("GET")
        r.HandleFunc("/{to}/{from}", s.in).Methods("POST")

        errch := s.control()

        // TODO IP bytes serialize correctly to string in all cases?
        at := fmt.Sprintf("%v:%v", ip, port)
        http.ListenAndServe(at, r)

        s.stream<- ctrl {done: true,}
        err := <-errch

        if err != nil {
                return errors.Wrap(err, "Ctrl loop failed")
        }

        return nil
}

type ctrl struct {
        done bool
        pkt *Proto
        reply chan<- *Proto
}

type server struct {
        auth Auth

        msgs map[PeerAddr]Message

        stream chan ctrl
}

func (s *server) in(w http.ResponseWriter, r *http.Request) {
}

func (s *server) out(w http.ResponseWriter, r *http.Request) {
}

func (s *server) control() <-chan error {
        errch := make(chan error)

        go func() {
                err := s.loop()
                errch<- err
        }()

        return errch
}

func (s *server) loop() error {
        for input := range s.stream {
                out, err := s.handle(input.pkt)

                // Signal close
                if out == nil {
                        return nil
                }

                if err != nil {
                        return err
                }

                input.reply<- out
        }

        return nil
}

func (s *server) handle(input *Proto) (*Proto, error) {
        switch input.Cmd {
        case Recv:
                return s.send(input)
        case Send:
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

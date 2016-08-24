package lib

import (
        "fmt"
        "net/http"

        "github.com/pkg/errors"
        "github.com/gorilla/mux"
)

func Serve(ip []byte, port uint, auth Auth) error {
        ctrlstream := make(chan ctrl, 1)
        m := &model{}

        s := &server{}
        s.stream = ctrlstream

        m.auth = auth
        m.stream = ctrlstream

        r := mux.NewRouter()

        r.HandleFunc("/{to}", s.recv).Methods("GET")
        r.HandleFunc("/{to}/{from}", s.send).Methods("POST")

        errch := m.spawn()

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
        stream chan<- ctrl
}

func (s *server) send(w http.ResponseWriter, r *http.Request) {
}

func (s *server) recv(w http.ResponseWriter, r *http.Request) {
}

type model struct {
        auth Auth
        msgs map[PeerAddr]Message
        stream <-chan ctrl
}

func (m *model) spawn() <-chan error {
        errch := make(chan error)

        go func() {
                err := m.loop()
                errch<- err
        }()

        return errch
}

func (m *model) loop() error {
        for input := range m.stream {
                out, err := m.handle(input.pkt)

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

func (m *model) handle(input *Proto) (*Proto, error) {
        switch input.Cmd {
        case Recv:
                return m.send(input)
        case Send:
                return m.recv(input)
        default:
                return nil, fmt.Errorf("Not handling Proto command: %v", input.Cmd)
        }
}

func (m *model) send(input *Proto) (*Proto, error) {
        return nil, nil
}

func (m *model) recv(input *Proto) (*Proto, error) {
        return nil, nil
}

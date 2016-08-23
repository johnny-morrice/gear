package lib

import (
        "net/http"
        "io"

        "golang.org/x/crypto/openpgp"
)

type remote struct {
        addr string
        port string
        client *http.Client
}

func (rem *remote) send(r io.Reader, me, them *openpgp.Entity) error {
        return nil
}

func (rem *remote) recv(w io.Writer, me *openpgp.Entity) error {
        return nil
}

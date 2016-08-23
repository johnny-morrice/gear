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

func (rem *remote) upload(r io.Reader, me, them *openpgp.Entity) error {
        return nil
}

func (rem *remote) download(w io.Writer, me *openpgp.Entity) error {
        return nil
}

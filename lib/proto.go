package lib

import (
        "io"

        "golang.org/x/crypto/openpgp"
)

type Cmd string

const (
        Download = Cmd("download")
        Upload = Cmd("upload")
        Data = Cmd("data")
        Err = Cmd("error")
)

type Proto struct {
        Cmd Cmd
        From *openpgp.Entity
        To *openpgp.Entity
        Data io.Reader
}

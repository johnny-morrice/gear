package lib

type Cmd string

const (
        Download = Cmd("download")
        Upload = Cmd("upload")
        Data = Cmd("data")
        Err = Cmd("error")
)

type PeerAddr string

type Message struct {
        From PeerAddr
        To PeerAddr
        Data []byte
}

type Proto struct {
        Message
        Cmd Cmd
}

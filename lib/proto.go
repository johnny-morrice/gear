package lib

type Cmd string

const (
        Send = Cmd("send")
        Recv = Cmd("recv")
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

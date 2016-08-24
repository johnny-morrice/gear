package lib

type Cmd string

const (
        Send = Cmd("send")
        Recv = Cmd("recv")
        Data = Cmd("data")
        Close = Cmd("close")
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

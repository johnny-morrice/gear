package lib

type Auth interface {
        Validate(from PeerAddr, data []byte) error
}

type pgpauth struct {

}

package lib

import (
	"fmt"
	"io"

	"github.com/pkg/errors"

	"golang.org/x/crypto/openpgp"
)

type CipherText io.Reader
type PlainText io.Reader

type Crypto interface {
	Decrypt(from, to PeerAddr, data CipherText) (PlainText, error)
	Encrypt(from PeerAddr, to []PeerAddr, data PlainText) (CipherText, error)
}

type pgpsec struct {
	emap entitymap
	keys openpgp.KeyRing

	prompter openpgp.PromptFunction
}

func (sec *pgpsec) Decrypt(from, to PeerAddr, data CipherText) (PlainText, error) {
	r := io.Reader(data)

	md, err := openpgp.ReadMessage(r, sec.keys, sec.prompter, nil)

	if err == nil {
		err = sec.checkSecurity(md, from, to)
	} else {
		err = errors.Wrap(err, "openpgp.ReadMessage failed")
	}

	if err != nil {
		return nil, errors.Wrap(err, "Decrypt failed")
	}

	return  PlainText(md.UnverifiedBody), nil
}

func (sec *pgpsec) Encrypt(from PeerAddr, to []PeerAddr, data PlainText) (CipherText, error) {
	ciphR, ciphW := io.Pipe()

	toE, terr := sec.emap.getall(to...)

	fromE, ferr := sec.emap.get(from)

	var err error

	if terr != nil {
		err = errors.Wrap(terr, "no receiver key found")
	}

	if ferr != nil {
		err = errors.Wrap(ferr, "no sender key found")
	}

	if err != nil {
		err = encrypt(ciphW, data, toE, fromE)
	}

	if err != nil {
		return nil, errors.Wrap(err, "Encrypt failed")
	}

	return CipherText(ciphR), nil
}

func encrypt(ciphW io.WriteCloser, plainR io.Reader, toE []*openpgp.Entity, fromE *openpgp.Entity) error {
	plainW, err := openpgp.Encrypt(ciphW, toE, fromE, nil, nil)

	if err == nil {
		_, err = io.Copy(plainW, plainR)

		plainW.Close()

		if err != nil {
			err = errors.Wrap(err, "copy failed")
		}
	} else {
		err = errors.Wrap(err, "openpgp.Encrypt failed")
	}

	return err
}

func (sec *pgpsec) checkSecurity(md *openpgp.MessageDetails, from, to PeerAddr) error {
	sender, skerr := sec.emap.get(from)
	receiver, rkerr := sec.emap.get(to)

	if !md.IsEncrypted {
		return errors.New("Message was not encrypted")
	}

	if !md.IsSigned {
		return errors.New("Message was not signed")
	}

	if skerr != nil {
		return errors.Wrap(skerr, "no sender key found")
	}

	if rkerr != nil {
		return errors.Wrap(rkerr, "no receiver key found")
	}

	// Check sanity of address consistency.
	if md.DecryptedWith.Entity != receiver {
		return errors.New("bad receiver")
	}

	if md.SignedBy.Entity != sender {
		return errors.New("bad sender")
	}

	return nil
}

type entitymap map[PeerAddr]*openpgp.Entity

func (emap entitymap) get(peer PeerAddr) (*openpgp.Entity, error) {
	e, ok := emap[peer]

	if !ok {
		return nil, fmt.Errorf("No entity for '%v'", peer)
	}

	return e, nil
}

func (emap entitymap) getall(peers... PeerAddr) ([]*openpgp.Entity, error) {

	ents := make([]*openpgp.Entity, len(peers))

	// Not returning early on error is awkward but prevents analysis
	// by timing.
	var err error
	for i, p := range peers {
		e, gerr := emap.get(p)

		ents[i] = e

		if gerr != nil {
			err = gerr
		}
	}

	if err != nil {
		return nil, err
	}

	return ents, nil
}

func (emap entitymap) put(peer PeerAddr, e *openpgp.Entity) {
	emap[peer] = e
}

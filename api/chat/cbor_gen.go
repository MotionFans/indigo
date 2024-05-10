// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package chat

import (
	"fmt"
	"io"
	"math"
	"sort"

	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

func (t *ActorDeclaration) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{162}); err != nil {
		return err
	}

	// t.LexiconTypeID (string) (string)
	if len("$type") > 1000000 {
		return xerrors.Errorf("Value in field \"$type\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("$type"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("$type")); err != nil {
		return err
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("chat.bsky.actor.declaration"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("chat.bsky.actor.declaration")); err != nil {
		return err
	}

	// t.AllowIncoming (string) (string)
	if len("allowIncoming") > 1000000 {
		return xerrors.Errorf("Value in field \"allowIncoming\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("allowIncoming"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("allowIncoming")); err != nil {
		return err
	}

	if len(t.AllowIncoming) > 1000000 {
		return xerrors.Errorf("Value in field t.AllowIncoming was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.AllowIncoming))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.AllowIncoming)); err != nil {
		return err
	}
	return nil
}

func (t *ActorDeclaration) UnmarshalCBOR(r io.Reader) (err error) {
	*t = ActorDeclaration{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("ActorDeclaration: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringWithMax(cr, 1000000)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.LexiconTypeID (string) (string)
		case "$type":

			{
				sval, err := cbg.ReadStringWithMax(cr, 1000000)
				if err != nil {
					return err
				}

				t.LexiconTypeID = string(sval)
			}
			// t.AllowIncoming (string) (string)
		case "allowIncoming":

			{
				sval, err := cbg.ReadStringWithMax(cr, 1000000)
				if err != nil {
					return err
				}

				t.AllowIncoming = string(sval)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}

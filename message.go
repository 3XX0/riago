package riago

import (
	"errors"

	"code.google.com/p/goprotobuf/proto"
)

var ErrInvalidResponseCode = errors.New("invalid response code")

// Encodes a request code and proto structure into a message byte buffer.
func encode(code uint8, req proto.Message) (buf []byte, err error) {
	if req != nil {
		buf, err = proto.Marshal(req)
		if err != nil {
			return
		}
	}
	size := uint32(len(buf) + 1)
	header := []byte{byte(size >> 24), byte(size >> 16), byte(size >> 8), byte(size), code}
	buf = append(header, buf...)
	return
}

// Decodes a message byte buffer into a proto response, error code or nil.
// Resulting object depends on response type.
func decode(buf []byte, resp proto.Message) error {
	if len(buf) < 1 {
		return ErrInvalidResponseCode
	}
	code := uint8(buf[0])
	buf = buf[1:]

	if code == MsgRpbErrorResp {
		resp = new(RpbErrorResp)
	}
	if resp == nil {
		return nil
	}
	err := proto.Unmarshal(buf, resp)

	e, ok := resp.(*RpbErrorResp)
	if ok && err == nil {
		err = errors.New(string(e.Errmsg))
	}
	return err
}

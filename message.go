package riago

import (
	"errors"

	"code.google.com/p/goprotobuf/proto"
)


var (
	ErrInvalidResponseBody = errors.New("invalid response body")
	ErrInvalidResponseCode = errors.New("invalid response code")
	ErrInvalidRequestCode  = errors.New("invalid request code")
)

// Encodes a request code and proto structure into a message byte buffer.
func encode(code byte, req proto.Message) (buf []byte, err error) {
	var reqbuf []byte

	if req != nil {
		reqbuf, err = proto.Marshal(req)
		if err != nil {
			return
		}
	}

	size := int32(len(reqbuf) + 1)
	buf = []byte{byte(size >> 24), byte(size >> 16), byte(size >> 8), byte(size), code}
	buf = append(buf, reqbuf...)

	return
}

// Decodes a message byte buffer into a proto response, error code or nil
// Resulting object depends on response type.
func decode(buf []byte, resp proto.Message) (err error) {
	var respbuf []byte

	if len(buf) < 1 {
		return ErrInvalidResponseCode
	}

	code := buf[0]

	if len(buf) > 1 {
		respbuf = buf[1:]
	} else {
		respbuf = make([]byte, 0)
	}

	if code < 0 || code > 60 {
		return ErrInvalidResponseCode
	}

	switch code {
	case MsgRpbErrorResp:
		errResp := &RpbErrorResp{}
		if err = proto.Unmarshal(respbuf, errResp); err == nil {
			err = errors.New(string(errResp.Errmsg))
		}
	case MsgRpbPingResp, MsgRpbSetClientIdResp, MsgRpbSetBucketResp, MsgRpbDelResp:
		resp = nil
	default:
		err = proto.Unmarshal(respbuf, resp)
	}

	return
}

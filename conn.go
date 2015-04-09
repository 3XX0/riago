package riago

import (
	"code.google.com/p/goprotobuf/proto"
	"github.com/3XX0/pooly"
	"io"
	"net"
	"time"
)

type Conn struct {
	conn         *net.TCPConn
	lastChecked  time.Time
	writeTimeout time.Duration
	readTimeout  time.Duration
}

func Riak(c *pooly.Conn) *Conn {
	return c.Interface().(*Conn)
}

// Encode and write a request to the Riak server.
func (c *Conn) request(code uint8, req proto.Message) error {
	buf, err := encode(code, req)
	if err != nil {
		return err
	}
	if c.writeTimeout > 0 {
		c.conn.SetWriteDeadline(time.Now().Add(c.writeTimeout))
	}
	_, err = c.conn.Write(buf)
	return err
}

// Read and decode a response from the Riak server.
func (c *Conn) response(resp proto.Message) error {
	if c.readTimeout > 0 {
		c.conn.SetReadDeadline(time.Now().Add(c.readTimeout))
	}

	buf := make([]byte, 4)
	if _, err := io.ReadFull(c.conn, buf); err != nil {
		return err
	}
	size := uint32(buf[0])<<24 | uint32(buf[1])<<16 | uint32(buf[2])<<8 | uint32(buf[3])

	buf = make([]byte, size)
	if _, err := io.ReadFull(c.conn, buf); err != nil {
		return err
	}

	return decode(buf, resp)
}

func (c *Conn) do(code uint8, req proto.Message, resp proto.Message) error {
	if err := c.request(code, req); err != nil {
		return err
	}
	return c.response(resp)
}

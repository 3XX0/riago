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

func RConn(c *pooly.Conn) *Conn {
	return c.Interface().(*Conn)
}

// Encode and write a request to the Riak server.
func (c *Conn) request(code byte, req proto.Message) error {
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

	sizebuf := make([]byte, 4)
	if _, err := io.ReadFull(c.conn, sizebuf); err != nil {
		return err
	}

	size := int(sizebuf[0])<<24 + int(sizebuf[1])<<16 + int(sizebuf[2])<<8 + int(sizebuf[3])
	buf := make([]byte, size)
	if _, err := io.ReadFull(c.conn, buf); err != nil {
		return err
	}

	return decode(buf, resp)
}

func (c *Conn) do(code byte, req proto.Message, resp proto.Message) error {
	if err := c.request(code, req); err != nil {
		return err
	}
	return c.response(resp)
}

func (c *Conn) Ping() error {
	return c.do(MsgRpbPingReq, nil, nil)
}

// Performs a Riak Server info request.
func (c *Conn) ServerInfo() (resp *RpbGetServerInfoResp, err error) {
	resp = new(RpbGetServerInfoResp)
	err = c.do(MsgRpbGetServerInfoReq, nil, resp)
	return
}

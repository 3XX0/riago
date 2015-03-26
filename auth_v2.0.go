// +build 2.0

package riago

// Performs a Riak Auth request.
func (c *Conn) Authenticate(req *RpbAuthReq) error {
	return c.do(MsgRpbAuthReq, req, nil)
}

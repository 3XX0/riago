package riago

// Performs a Riak Get request.
func (c *Conn) Get(req *RpbGetReq) (resp *RpbGetResp, err error) {
	resp = new(RpbGetResp)
	err = c.do(MsgRpbGetReq, req, resp)
	return
}

// Performs a Riak Put request.
func (c *Conn) Put(req *RpbPutReq) (resp *RpbPutResp, err error) {
	resp = new(RpbPutResp)
	err = c.do(MsgRpbPutReq, req, resp)
	return
}

// Performs a Riak Del request.
func (c *Conn) Del(req *RpbDelReq) error {
	return c.do(MsgRpbDelReq, req, nil)
}

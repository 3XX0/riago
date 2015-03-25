// +build 2.0

package riago

// Performs a Riak CRDT Fetch request.
func (c *Conn) DtFetch(req *DtFetchReq) (resp *DtFetchResp, err error) {
	resp = new(DtFetchResp)
	err = c.do(MsgDtFetchReq, req, resp)
	return
}

// Performs a Riak CRDT Update request.
func (c *Conn) DtUpdate(req *DtUpdateReq) (resp *DtUpdateResp, err error) {
	resp = new(DtUpdateResp)
	err = c.do(MsgDtUpdateReq, req, resp)
	return
}

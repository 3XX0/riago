package riago

// Perform a Riak Search Query request.
func (c *Conn) SearchQuery(req *RpbSearchQueryReq) (resp *RpbSearchQueryResp, err error) {
	resp = new(RpbSearchQueryResp)
	err = c.do(MsgRpbSearchQueryReq, req, resp)
	return
}

package riago

// Perform a Riak Search Query request.
func (c *Conn) SearchQuery(req *RpbSearchQueryReq) (resp *RpbSearchQueryResp, err error) {
	resp = new(RpbSearchQueryResp)
	err = c.do(MsgRpbSearchQueryReq, req, resp)
	return
}

// Perform a Riak Yokozuna Index Get request.
func (c *Conn) YokozunaIndexGet(req *RpbYokozunaIndexGetReq) (resp *RpbYokozunaIndexGetResp, err error) {
	resp = new(RpbYokozunaIndexGetResp)
	err = c.do(MsgRpbYokozunaIndexGetReq, req, resp)
	return
}

// Perform a Riak Yokozuna Index Put request.
func (c *Conn) YokozunaIndexPut(req *RpbYokozunaIndexPutReq) error {
	return c.do(MsgRpbYokozunaIndexPutReq, req, nil)
}

// Perform a Riak Yokozuna Index Delete request.
func (c *Conn) YokozunaIndexDelete(req *RpbYokozunaIndexDeleteReq) error {
	return c.do(MsgRpbYokozunaIndexDeleteReq, req, nil)
}

// Perform a Riak Yokozuna Index Get request.
func (c *Conn) YokozunaSchemaGet(req *RpbYokozunaSchemaGetReq) (resp *RpbYokozunaSchemaGetResp, err error) {
	resp = new(RpbYokozunaSchemaGetResp)
	err = c.do(MsgRpbYokozunaSchemaGetReq, req, resp)
	return
}

// Perform a Riak Yokozuna Schema Put request.
func (c *Conn) YokozunaSchemaPut(req *RpbYokozunaSchemaPutReq) error {
	return c.do(MsgRpbYokozunaSchemaPutReq, req, nil)
}

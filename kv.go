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

// Perform a Riak Get Bucket request.
func (c *Conn) GetBucket(req *RpbGetBucketReq) (resp *RpbGetBucketResp, err error) {
	resp = new(RpbGetBucketResp)
	err = c.do(MsgRpbGetBucketReq, req, resp)
	return
}

// Perform a Riak Set Bucket request.
func (c *Conn) SetBucket(req *RpbSetBucketReq) error {
	return c.do(MsgRpbSetBucketReq, req, nil)
}

// Perform a Riak List Buckets request. The protobufs say that it will return
// multiple responses but it in fact does not.
func (c *Conn) ListBuckets(req *RpbListBucketsReq) (resp *RpbListBucketsResp, err error) {
	resp = new(RpbListBucketsResp)
	err = c.do(MsgRpbListBucketsReq, req, resp)
	return
}

// Perform a Riak List Keys request. Returns multiple list keys responses.
func (c *Conn) ListKeys(req *RpbListKeysReq) ([]*RpbListKeysResp, error) {
	var resps []*RpbListKeysResp

	if err := c.request(MsgRpbListKeysReq, req); err != nil {
		return nil, err
	}
	for {
		resp := new(RpbListKeysResp)
		if err := c.response(resp); err != nil {
			return nil, err
		}
		resps = append(resps, resp)

		if resp.GetDone() {
			break
		}
	}
	return resps, nil
}

// Perform a Riak Index (2i) request. The protobufs say that it will return
// multiple responses but it in fact does not.
func (c *Conn) Index(req *RpbIndexReq) (resp *RpbIndexResp, err error) {
	resp = new(RpbIndexResp)
	err = c.do(MsgRpbIndexReq, req, resp)
	return
}

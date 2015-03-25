package riago

// Performs a Riak Get Bucket request.
func (c *Conn) GetBucket(req *RpbGetBucketReq) (resp *RpbGetBucketResp, err error) {
	resp = new(RpbGetBucketResp)
	err = c.do(MsgRpbGetBucketReq, req, resp)
	return
}

// Performs a Riak Set Bucket request.
func (c *Conn) SetBucket(req *RpbSetBucketReq) error {
	return c.do(MsgRpbSetBucketReq, req, nil)
}

// Performs a Riak List Buckets request.
// The protobufs say that it will return multiple responses but it in fact does not.
func (c *Conn) ListBuckets(req *RpbListBucketsReq) (resp *RpbListBucketsResp, err error) {
	resp = new(RpbListBucketsResp)
	err = c.do(MsgRpbListBucketsReq, req, resp)
	return
}

// Performs a Riak List Keys request.
// Returns multiple list keys responses.
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

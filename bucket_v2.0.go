// +build 2.0

package riago

// Performs a Riak Reset Bucket request.
func (c *Conn) ResetBucket(req *RpbResetBucketReq) error {
	return c.do(MsgRpbResetBucketReq, req, nil)
}

// Performs a Riak Get Bucket Type request.
func (c *Conn) GetBucketType(req *RpbGetBucketTypeReq) (resp *RpbGetBucketResp, err error) {
	resp = new(RpbGetBucketResp)
	err = c.do(MsgRpbGetBucketTypeReq, req, resp)
	return
}

// Performs a Riak Set Bucket Type request.
func (c *Conn) SetBucketType(req *RpbSetBucketTypeReq) error {
	return c.do(MsgRpbSetBucketTypeReq, req, nil)
}

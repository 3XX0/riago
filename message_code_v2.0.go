// +build 2.0

package riago

const (
	MsgRpbResetBucketReq         = 29
	MsgRpbResetBucketResp        = 30
	MsgRpbGetBucketTypeReq       = 31
	MsgRpbSetBucketTypeResp      = 32
	MsgRpbCSBucketReq            = 40 // XXX Riak CS only
	MsgRpbCSBucketResp           = 41 // XXX Riak CS only
	MsgRpbCounterUpdateReq       = 50
	MsgRpbCounterUpdateResp      = 51
	MsgRpbCounterGetReq          = 52
	MsgRpbCounterGetResp         = 53
	MsgRpbYokozunaIndexGetReq    = 54
	MsgRpbYokozunaIndexGetResp   = 55
	MsgRpbYokozunaIndexPutReq    = 56
	MsgRpbYokozunaIndexDeleteReq = 57
	MsgRpbYokozunaSchemaGetReq   = 58
	MsgRpbYokozunaSchemaGetResp  = 59
	MsgRpbYokozunaSchemaPutReq   = 60
	MsgDtFetchReq                = 80
	MsgDtFetchResp               = 81
	MsgDtUpdateReq               = 82
	MsgDtUpdateResp              = 83
	MsgRpbAuthReq                = 253
	MsgRpbAuthResp               = 254
	MsgRpbStartTls               = 255
)

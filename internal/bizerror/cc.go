package bizerror

import "errors"

const (
	ErrCodeKvAddBadParams         = 1011001
	ErrCodeKvAddClusterNotFound   = 1011002
	ErrCodeKvAddClusterFind       = 1011003
	ErrCodeClusterIdNotMatchAppId = 1011004
	ErrCodeKvAddKeyFind           = 1011005
	ErrCodeKvAddKeyExisted        = 1011006
)

var KvAddCodeMap = ErrCodeMap{
	ErrCodeKvAddBadParams:         errors.New("参数有误"),
	ErrCodeKvAddClusterNotFound:   errors.New("cluster 不存在"),
	ErrCodeKvAddClusterFind:       errors.New("cluster 查询失败"),
	ErrCodeClusterIdNotMatchAppId: errors.New("appId 有误"),
	ErrCodeKvAddKeyFind:           errors.New("key 校验失败"),
	ErrCodeKvAddKeyExisted:        errors.New("key 已经存在"),
}

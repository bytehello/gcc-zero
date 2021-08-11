package bizerror

import "errors"

var (
	ErrCodeClusterFind          = 1031001
	ErrCodeCLusterNameExisted   = 1031002
	ErrCodeClusterUpdate        = 1031003
	ErrCodeClusterAppNotExisted = 1031004
	ErrCodeClusterAdd           = 1031005
)
var ClusterMap = ErrCodeMap{
	ErrCodeClusterFind:          errors.New("查询失败"),
	ErrCodeCLusterNameExisted:   errors.New("名称已经存在"),
	ErrCodeClusterUpdate:        errors.New("更新失败"),
	ErrCodeClusterAppNotExisted: errors.New("app 不存在"),
	ErrCodeClusterAdd:           errors.New("添加失败"),
}

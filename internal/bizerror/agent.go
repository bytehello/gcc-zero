package bizerror

import "errors"

const (
	ErrCodeReleaseKvFail = 1051001
	ErrCodeReleaseKvStat = 1051002
)

var AgentMap = ErrCodeMap{
	ErrCodeReleaseKvFail: errors.New("发布失败"),
	ErrCodeReleaseKvStat: errors.New("获取配置信息失败"),
}

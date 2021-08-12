package bizerror

import (
	"github.com/bytehello/gkit/bizerror"
	"github.com/pkg/errors"
)

type BizCode int
type ErrCodeMap map[int]error

var (
	Silence   = bizerror.Silence
	IsSilence = bizerror.IsSilence

	err = map[BizCode]ErrCodeMap{
		// 请按顺序使用，最好不要跳级
		// 错误码和错误处理最好一对一，不要偷懒
		100: defaultErrCodeMap,
		101: KvAddCodeMap,
		102: AppCodeMap,
		103: ClusterMap,
		104: ClientMap,
	}
)

// New return BizError by code
// Example:
//		New(1010000)
func New(code int) error {
	err := errByCode(code)
	return bizerror.Wrap(code, err.Error(), err)
}

// Newf return BizError by code and extra message
// Example:
//		Newf(1010000, "extra message: %s", "other")
//		Newf(1010000, "extra message")
func Newf(code int, format string, args ...interface{}) error {
	err := errByCode(code)
	return bizerror.Wrap(code, err.Error(), errors.Wrapf(err, format, args...))
}

// Wrap return BizError by code and error
//
// Example:
//     Wrap(1001001, err)
func Wrap(code int, underlyingErr error) error {
	err := errByCode(code)

	return bizerror.Wrap(code, err.Error(), mergeErr(err, underlyingErr))
}

// Wrapf return BizError by code, error and extra message
//
// Example:
//     Wrapf(1001001, err, "关于 err 的更多信息")
//     Wrapf(1001001, err, "关于 err 的更多信息: %s", "及参数")
func Wrapf(code int, underlyingErr error, format string, args ...interface{}) error {
	err := errByCode(code)

	return bizerror.Wrapf(code, err.Error(), mergeErr(err, underlyingErr), format, args...)
}

// Customized return BizError by code and customized message
func Customized(code int, msg string) error {
	err := errByCode(code)
	errMsg := err.Error()
	if msg != "" {
		errMsg = msg
	}
	return bizerror.Wrap(code, errMsg, err) // Use err as underlying error
}

func errByCode(code int) error {
	bizCode := BizCode(code / 10000)
	codeMap, ok := err[bizCode]
	if !ok {
		return errors.Errorf("Unknown Error Code: %d", code)
	}
	err, ok := codeMap[code]
	if !ok {
		return errors.Errorf("Unknown Error Code: %d", code)
	}
	return err
}

// 将 error 上的属性转移到 underlying error 上
func mergeErr(err, underlyingErr error) (retErr error) {
	retErr = underlyingErr

	var decorators []func(error) error
	if IsSilence(err) {
		decorators = append(decorators, Silence)
	}

	for _, decorator := range decorators {
		retErr = decorator(retErr)
	}

	return retErr
}

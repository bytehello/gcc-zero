package errorhandler

import (
	"github.com/bytehello/gcc-zero/common/errorx"
	"github.com/tal-tech/go-zero/core/logx"
	"net/http"
)

func Handler(err error) (i int, i2 interface{}) {
	switch e := err.(type) {
	case *errorx.CodeError:
		return http.StatusOK, e.Data()
	default:
		logx.Info("StatusInternalServerError err:", err)
		return http.StatusInternalServerError, nil
	}
}

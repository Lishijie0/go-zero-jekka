package xresp

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"jekka-api-go/pkg/response/xerr"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// HttpResult http返回
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	// 成功返回
	if err == nil {
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
		return
	}
	// 错误返回
	errCode := xerr.ServerCommonError
	errMsg := xerr.MapErrMsg(xerr.ServerCommonError)
	fmt.Println(err)
	causeErr := errors.Cause(err)
	// 自定义错误
	if e, ok := causeErr.(*xerr.CodeError); ok {
		errCode = e.GetErrCode()
		errMsg = e.GetErrMsg()
	} else {
		// grpc错误
		if grpcStatus, ok := status.FromError(causeErr); ok {
			grpcCode := uint32(grpcStatus.Code())
			if xerr.IsCodeErr(grpcCode) { //区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
				errCode = grpcCode
				errMsg = grpcStatus.Message()
			}
		}
	}

	logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)

	httpx.WriteJson(w, http.StatusBadRequest, Error(errCode, errMsg))
}

// ParamErrorResult http 参数错误返回
func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	errMsg := fmt.Sprintf("%s ,%s", xerr.MapErrMsg(xerr.ReuqestParamError), err.Error())
	httpx.WriteJson(w, http.StatusBadRequest, Error(xerr.ReuqestParamError, errMsg))
}

package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	// 系统错误，请稍后再试
	message[ServerCommonError] = "System error, please try again later"
	// 参数错误
	message[ReuqestParamError] = "Parameter error"
	// token 过期
	message[TokenExpireError] = "Token expired"
	// 生成token失败
	message[TokenGenerateError] = "Generate token error"
	// 数据库繁忙,请稍后再试
	message[DbError] = "Database is busy, please try again later"
	// 更新数据影响行数为0
	message[DbUpdateAffectedZeroError] = "Update affected 0 rows"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	}
	return message[ServerCommonError]
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	}
	return false
}

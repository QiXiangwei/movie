package library

// ========================****用户相关****===================================
const ERR_NO_USER_PHONE_IS_NULL    = 1000
const ERR_NO_PHONE_NOT_MATCH       = 1001
const ERR_NO_USER_PASSWARD_IS_NULL = 1002
const ERR_NO_PHONE_EXISTED         = 1003
const ERR_NO_REGISTER_FAILED       = 1004
const ERR_NO_LOGIN_FAILED          = 1005

const ERR_STR_USER_PHONE_IS_NULL    = "手机号不能为空"
const ERR_STR_PHONE_NOT_MATCH       = "手机号格式不正确"
const ERR_STR_USER_PASSWORD_IS_NULL = "密码不能为空"
const ERR_STR_PHONE_EXISTED         = "该手机号已注册"
const ERR_STR_REGISTER_FAILED       = "注册失败，请稍后重试"
const ERR_STR_LOGIN_FAILED          = "登陆失败，手机号或密码错误"

// ========================****频道相关****===================================
const ERR_NO_CHANNEL_NAME_IS_NULL   = 2000
const ERR_NO_CHANNEL_CREATE_FAILED  = 2001
const ERR_NO_CHANNEL_ID_ILLEGAL     = 2002
const ERR_NO_CHANNEL_UNEXISTED      = 2003
const ERR_NO_CHANNEL_ONLINE_FAILED  = 2004
const ERR_NO_CHANNEL_OFFLINE_FAILED = 2005
const ERR_NO_CHANNEL_DELETE_FAILED  = 2006

const ERR_STR_CHANNEL_NAME_IS_NULL   = "频道名称不能为空"
const ERR_STR_CHANNEL_CREATE_FAILED  = "创建频道失败"
const ERR_STR_CHANNEL_ID_ILLEGAL     = "频道id不合法"
const ERR_STR_CHANNEL_UNEXISTED      = "频道不存在"
const ERR_STR_CHANNEL_ONLINE_FAILED  = "频道上线失败"
const ERR_STR_CHANNEL_OFFLINE_FAILED = "频道下线失败"
const ERR_STR_CHANNEL_DELETE_FAILED  = "频道删除失败"

var ERR_MAP = map[int]string{
	// ========================****用户相关****===================================
	ERR_NO_USER_PHONE_IS_NULL    : ERR_STR_USER_PHONE_IS_NULL,
	ERR_NO_PHONE_NOT_MATCH       : ERR_STR_PHONE_NOT_MATCH,
	ERR_NO_USER_PASSWARD_IS_NULL : ERR_STR_USER_PASSWORD_IS_NULL,
	ERR_NO_PHONE_EXISTED         : ERR_STR_PHONE_EXISTED,
	ERR_NO_REGISTER_FAILED       : ERR_STR_REGISTER_FAILED,
	ERR_NO_LOGIN_FAILED          : ERR_STR_LOGIN_FAILED,

	// ========================****频道相关****===================================
	ERR_NO_CHANNEL_NAME_IS_NULL   : ERR_STR_CHANNEL_NAME_IS_NULL,
	ERR_NO_CHANNEL_CREATE_FAILED  : ERR_STR_CHANNEL_CREATE_FAILED,
	ERR_NO_CHANNEL_ID_ILLEGAL     : ERR_STR_CHANNEL_ID_ILLEGAL,
	ERR_NO_CHANNEL_UNEXISTED      : ERR_STR_CHANNEL_UNEXISTED,
	ERR_NO_CHANNEL_ONLINE_FAILED  : ERR_STR_CHANNEL_ONLINE_FAILED,
	ERR_NO_CHANNEL_OFFLINE_FAILED : ERR_STR_CHANNEL_OFFLINE_FAILED,
	ERR_NO_CHANNEL_DELETE_FAILED  : ERR_STR_CHANNEL_DELETE_FAILED,
}

func ErrMap(errNo int) string {
	return ERR_MAP[errNo]
}

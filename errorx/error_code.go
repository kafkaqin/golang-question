package errorx

const (
	// 通用错误码
	CODE_UNDEFINED                           = 100001
	CODE_NOT_FOUND                           = 100002
	CODE_TIMEOUT                             = 100003
	CODE_INTERNALSERVER                      = 100004
	CODE_INVALID_PARAMS                      = 100005  // 请求参数错误
	CODE_BIZ_UNIQUE_CODE_ALREADY_EXISTS      = 100006  // 业务幂等code已存在
	CODE_SORT_NUM_ALREADY_EXISTS             = 100007  // 排序号已存在
	CODE_CFG_NOT_EXISTS                      = 100008  //常量配置错误
	CODE_DATA_LEN_ERROR                      = 100009  //数据格式错误
	CODE_DATA_CFG_TABLE_ERROR                = 1000011 //配置表错误
	CODE_CONTENT_CONTAINS_ILLEGAL_CHARACTERS = 1000012 //内容包含非法字符
	CODE_CONTENT_IS_EMPTY                    = 1000013 //内容为空
	CODE_CONTENT_LENGTH_OVER_LIMIT           = 1000014 //内容长度超出上限

	// 认证鉴权
	CODE_NO_SIGN_IN                     = 200001 // 未登录
	CODE_INCORRECT_USERNAME             = 200002 // 用户名或密码错误
	CODE_USER_ALREADY_EXISTS            = 200003 // 用户已存在
	CODE_TOKEN_IS_NOT_VALID             = 200004 //token 失效  已经弃用！！！
	CODE_TOKEN_GENERATE_FAILED          = 200005 //token 生成失败
	CODE_INCORRECT_PASSWORD_OR_USERNAME = 200006 //密码错误
	CODE_TOKEN_INVALID                  = 200007 //token 无效
	CODE_USER_IS_BANNED                 = 200008 // 用户已禁用
	CODE_TOKEN_EXPIRED                  = 200009
	CODE_DISPLAY_UID_EXISTS             = 2000010

	// 用户
	CODE_USER_ROLE_NOT_EXISTS                  = 201001 //角色不存在
	CODE_USER_ROLE_GENDER_ERROR                = 201002 //用户角色错误
	CODE_USER_NAME_IS_EMPTY                    = 201003 //角色名为空
	CODE_USER_NAME_CONTAINS_ILLEGAL_CHARACTERS = 201004 //角色名包含非法字符
	CODE_USER_NAME_REPEATED                    = 201005 //角色名跟原本的名称重复
	CODE_USER_NAME_TOO_LONG                    = 201006 //角色名太长
	CODE_USER_NAME_HAS_TAKEN                   = 201007 //角色名已被占用
	CODE_USER_ALREADY_INIT_ROLE_COMPLETE       = 201008 // 用户已完成初始化角色

	CODE_SECRET_TOO_SHORT = 1001
)

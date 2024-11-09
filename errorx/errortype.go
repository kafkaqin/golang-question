package errorx

const (
	// 通用错误码
	ERRTYPE_UNDEFINED                           ErrType = "undefined"
	ERRTYPE_NOT_FOUND                           ErrType = "not_found"
	ERRTYPE_TIMEOUT                             ErrType = "timeout"
	ERRTYPE_INTERNAL_SERVER                     ErrType = "internal_server_error"
	ERRTYPE_INVALID_PARAMS                      ErrType = "invalid params"                      // 请求参数错误
	ERRTYPE_BIZ_UNIQUE_CODE_ALREADY_EXISTS      ErrType = "unique_code_already_exists"          // 业务幂等code已存在
	ERRTYPE_SORT_NUM_ALREADY_EXISTS             ErrType = "sort_num_already_exists"             // 排序号已存在
	ERRTYPE_CFG_NOT_EXISTS                      ErrType = "cfg_not_exists"                      //常量配置错误
	ERRTYPE_DATA_LEN_ERROR                      ErrType = "data_len_error"                      //数据格式错误
	ERRTYPE_DATA_CFG_TABLE_ERROR                ErrType = "data_cfg_table_error"                //配置表错误
	ERRTYPE_CONTENT_CONTAINS_ILLEGAL_CHARACTERS ErrType = "content_contains_illegal_characters" //内容包含非法字符
	ERRTYPE_CONTENT_IS_EMPTY                    ErrType = "content_is_empty"                    //内容为空
	ERRTYPE_CONTENT_LENGTH_OVER_LIMIT           ErrType = "content_length_over_limit"           //内容长度超出上限

	// 认证鉴权
	ERRTYPE_NO_SIGN_IN                     ErrType = "no_sign_in"                     // 未登录
	ERRTYPE_INCORRECT_USERNAME             ErrType = "incorrect_username"             // 用户名或密码错误
	ERRTYPE_USER_ALREADY_EXISTS            ErrType = "user_already_exists"            // 用户已存在
	ERRTYPE_TOKEN_IS_NOT_VALID             ErrType = "token_is_not_valid"             //token 失效  已经弃用！！！
	ERRTYPE_TOKEN_GENERATE_FAILED          ErrType = "token_generate_failed"          //token 生成失败
	ERRTYPE_INCORRECT_PASSWORD_OR_USERNAME ErrType = "incorrect_password_or_username" //密码错误
	ERRTYPE_TOKEN_INVALID                  ErrType = "token_invalid"                  //token 无效
	ERRTYPE_USER_IS_BANNED                 ErrType = "user_is_banned"                 // 用户已禁用
	ERRTYPE_TOKEN_EXPIRED                  ErrType = "token_expired"
	ERRTYPE_DISPLAY_UID_EXISTS             ErrType = "display_uid_exists"

	// 用户
	ERRTYPE_USER_ROLE_NOT_EXISTS                  ErrType = "user role not exists"                  //角色不存在
	ERRTYPE_USER_ROLE_GENDER_ERROR                ErrType = "user_role_gender_error"                //用户角色错误
	ERRTYPE_USER_NAME_IS_EMPTY                    ErrType = "user_name_is_empty"                    //角色名为空
	ERRTYPE_USER_NAME_CONTAINS_ILLEGAL_CHARACTERS ErrType = "user_name_contains_illegal_characters" //角色名包含非法字符
	ERRTYPE_USER_NAME_REPEATED                    ErrType = "user_name_repeated"                    //角色名跟原本的名称重复
	ERRTYPE_USER_NAME_TOO_LONG                    ErrType = "user_name_too_long"                    //角色名太长
	ERRTYPE_USER_NAME_HAS_TAKEN                   ErrType = "user_name_has_taken"                   //角色名已被占用
	ERRTYPE_USER_ALREADY_INIT_ROLE_COMPLETE       ErrType = "user_already_init_role_complete"       // 用户已完成初始化角色

	ERRTYPE_SECRET_TOO_SHORT = "secret_too_short"
)

package errorx

var codeToErrTypeMap = map[int]ErrType{
	CODE_UNDEFINED:                           ERRTYPE_UNDEFINED,
	CODE_NOT_FOUND:                           ERRTYPE_NOT_FOUND,
	CODE_TIMEOUT:                             ERRTYPE_TIMEOUT,
	CODE_INTERNALSERVER:                      ERRTYPE_INTERNAL_SERVER,
	CODE_INVALID_PARAMS:                      ERRTYPE_INVALID_PARAMS,                 // 请求参数错误
	CODE_BIZ_UNIQUE_CODE_ALREADY_EXISTS:      ERRTYPE_BIZ_UNIQUE_CODE_ALREADY_EXISTS, // 业务幂等code已存在
	CODE_SORT_NUM_ALREADY_EXISTS:             ERRTYPE_SORT_NUM_ALREADY_EXISTS,        // 排序号已存在
	CODE_CFG_NOT_EXISTS:                      ERRTYPE_CFG_NOT_EXISTS,                 //常量配置错误
	CODE_DATA_LEN_ERROR:                      ERRTYPE_DATA_LEN_ERROR,                 //数据格式错误
	CODE_DATA_CFG_TABLE_ERROR:                ERRTYPE_DATA_CFG_TABLE_ERROR,           //配置表错误
	CODE_CONTENT_CONTAINS_ILLEGAL_CHARACTERS: ERRTYPE_DATA_CFG_TABLE_ERROR,           //内容包含非法字符
	CODE_CONTENT_IS_EMPTY:                    ERRTYPE_CONTENT_IS_EMPTY,               //内容为空
	CODE_CONTENT_LENGTH_OVER_LIMIT:           ERRTYPE_CONTENT_LENGTH_OVER_LIMIT,      //内容长度超出上限

	// 认证鉴权
	CODE_NO_SIGN_IN:                     ERRTYPE_NO_SIGN_IN,                     // 未登录
	CODE_INCORRECT_USERNAME:             ERRTYPE_INCORRECT_USERNAME,             // 用户名或密码错误
	CODE_USER_ALREADY_EXISTS:            ERRTYPE_USER_ALREADY_EXISTS,            // 用户已存在
	CODE_TOKEN_IS_NOT_VALID:             ERRTYPE_TOKEN_IS_NOT_VALID,             //token 失效  已经弃用！！！
	CODE_TOKEN_GENERATE_FAILED:          ERRTYPE_TOKEN_GENERATE_FAILED,          //token 生成失败
	CODE_INCORRECT_PASSWORD_OR_USERNAME: ERRTYPE_INCORRECT_PASSWORD_OR_USERNAME, //密码错误
	CODE_TOKEN_INVALID:                  ERRTYPE_TOKEN_INVALID,                  //token 无效
	CODE_USER_IS_BANNED:                 ERRTYPE_USER_IS_BANNED,                 // 用户已禁用
	CODE_TOKEN_EXPIRED:                  ERRTYPE_TOKEN_EXPIRED,
	CODE_DISPLAY_UID_EXISTS:             ERRTYPE_DISPLAY_UID_EXISTS,

	// 用户
	CODE_USER_ROLE_NOT_EXISTS:                  ERRTYPE_USER_ROLE_NOT_EXISTS,                  //角色不存在
	CODE_USER_ROLE_GENDER_ERROR:                ERRTYPE_USER_ROLE_GENDER_ERROR,                //用户角色错误
	CODE_USER_NAME_IS_EMPTY:                    ERRTYPE_USER_NAME_IS_EMPTY,                    //角色名为空
	CODE_USER_NAME_CONTAINS_ILLEGAL_CHARACTERS: ERRTYPE_USER_NAME_CONTAINS_ILLEGAL_CHARACTERS, //角色名包含非法字符
	CODE_USER_NAME_REPEATED:                    ERRTYPE_USER_NAME_REPEATED,                    //角色名跟原本的名称重复
	CODE_USER_NAME_TOO_LONG:                    ERRTYPE_USER_NAME_TOO_LONG,                    //角色名太长
	CODE_USER_NAME_HAS_TAKEN:                   ERRTYPE_USER_NAME_HAS_TAKEN,                   //角色名已被占用
	CODE_USER_ALREADY_INIT_ROLE_COMPLETE:       ERRTYPE_USER_ALREADY_INIT_ROLE_COMPLETE,       // 用户已完成初始化角色

	CODE_SECRET_TOO_SHORT: ERRTYPE_SECRET_TOO_SHORT,
}

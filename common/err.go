package common

import "errors"

var (
	ErrGuzhiInvalid = errors.New("咕值信息不合法")
	ErrInvalidParam = errors.New("参数不合法")
	ErrInternal     = errors.New("服务器内部错误")
	ErrPrivacy      = errors.New("用户开启了“完全隐私保护”，获取数据失败")
)

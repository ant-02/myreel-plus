package response

import (
	"github.com/ant-02/myreel-plus/kitex_gen/model"
	"github.com/ant-02/myreel-plus/pkg/errno"
)

func BuildBaseResp(err error) *model.BaseResp {
	if err == nil {
		return &model.BaseResp{
			Code: errno.SuccessCode,
			Msg:  errno.Success.ErrorMsg,
		}
	}
	Errno := errno.ConvertErr(err)
	return &model.BaseResp{
		Code: Errno.ErrorCode,
		Msg:  Errno.ErrorMsg,
	}
}

func BuildSuccessResp() *model.BaseResp {
	return BuildBaseResp(nil) // 直接调用原始函数，传入 nil 表示无错误
}

func BuildTypeList[T any, U any](items []U, buildFunc func(U) T) []T {
	if len(items) == 0 {
		return nil
	}

	list := make([]T, len(items))
	for i, item := range items {
		list[i] = buildFunc(item)
	}
	return list
}

func IsSuccess(code int64) bool {
	return code == errno.SuccessCode
}

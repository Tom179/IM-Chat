package ctxdata

import "context"

// 解析出来的jwt信息存放到context中
func GetUid(ctx context.Context) string {
	if u, ok := ctx.Value(Identify).(string); ok {
		return u
	}
	return ""
}

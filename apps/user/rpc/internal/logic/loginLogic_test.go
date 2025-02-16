package logic

import (
	"Im-chat/Chat/apps/user/rpc/user"
	"context"
	"fmt"
	"testing"
)

func TestRegisterLogic_Register(t *testing.T) {
	fmt.Println("开始测试test")
	type args struct {
		in *user.LoginReq
	}
	tests := []struct {
		name string
		args args
		//want    *user.RegisterResp
		wangPrint bool
		wantErr   bool
	}{
		{
			"1", args{in: &user.LoginReq{
				Phone:    "15340407586",
				Password: "Qy85891607",
			}}, true, false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLoginLogic(context.Background(), svcCtx)
			got, err := l.Login(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wangPrint {
				t.Log(tt.name, got)
			}
		})
	}
}

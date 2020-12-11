package wechat

import ( 
	"fmt"
	"github.com/chadiandianwenrou/go/http/client" 
	"github.com/chadiandianwenrou/go/types"
)

 
type User struct{
	Errcode 		int 		`json:"errcode"`
	Userid  		string 		`json:"userid"`
	Mobile  		string 		`json:"mobile"`
}

func GetUser(access_token,userid string) (u *User, err error) {
	Url:= fmt.Sprintf(wxGetUserUrl,access_token,userid)
	ctx,err:=client.Get(Url)
	if err != nil{
		return nil,err
	}
	WechatUser := &User{}

	err =types.FormatJSON(ctx,WechatUser)
	if err != nil{
		return nil,err
	}
	return WechatUser,nil
}
package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

const (
	RECODEOK         = "0"
	RECODEDBERR      = "4001"
	RECODENODATA     = "4002"
	RECODEDATAEXIST  = "4003"
	RECODEDATAERR    = "4004"
	RECODESESSIONERR = "4101"
	RECODELOGINERR   = "4102"
	RECODEPARAMERR   = "4103"
	RECODEUSERERR    = "4104"
	RECODEHASHERR    = "4105"
	RECODEPWDERR     = "4106"
	RECODEEXISTSERR  = "4201"
	RECODEIPCERR     = "4202"
	RECODETHIRDERR   = "4301"
	RECODEIOERR      = "4302"
	RECODESERVERERR  = "4500"
	RECODEUNKNOWERR  = "4501"
)

var recodeText = map[string]string{
	RECODEOK:         "成功",
	RECODEDBERR:      "数据库操作错误",
	RECODENODATA:     "无数据",
	RECODEDATAEXIST:  "数据已存在",
	RECODEDATAERR:    "数据错误",
	RECODESESSIONERR: "用户未登录",
	RECODELOGINERR:   "用户登录失败",
	RECODEPARAMERR:   "参数错误",
	RECODEUSERERR:    "用户不存在或密码错误",
	RECODEHASHERR:    "HASH错误",
	RECODEPWDERR:     "密码错误",
	RECODEEXISTSERR:  "重复上传错误",
	RECODEIPCERR:     "IPC错误",
	RECODETHIRDERR:   "与以太坊交互失败",
	RECODEIOERR:      "文件读写错误",
	RECODESERVERERR:  "内部错误",
	RECODEUNKNOWERR:  "未知错误",
}

func RecodeText(code string) string {
	str, ok := recodeText[code]
	if ok {
		return str
	}
	return recodeText[RECODEUNKNOWERR]
}

type Resp struct {
	Errno  string      `json:"errno"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

//resp数据响应
func ResponseData(c echo.Context, resp *Resp) {
	resp.ErrMsg = RecodeText(resp.Errno)
	c.JSON(http.StatusOK, resp)
}

//读取dir目录下文件名带address的文件
func GetFileName(address, dirname string) (string, error) {

	data, err := ioutil.ReadDir(dirname)
	if err != nil {
		fmt.Println("read dir err", err)
		return "", err
	}
	for _, v := range data {
		if strings.Index(v.Name(), address) > 0 {
			//代表找到文件
			return v.Name(), nil
		}
	}

	return "", nil
}

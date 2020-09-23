package extra

import (
	"fmt"
	"socketIo/dbs"
	"socketIo/utils"

	"github.com/labstack/echo"
)

//注册功能
func Register(c echo.Context) (err error) {
	//1. 响应数据结构初始化
	var resp utils.Resp
	resp.Errno = utils.RECODEOK
	defer utils.ResponseData(c, &resp)
	account := &dbs.Account{}
	sql := fmt.Sprintf("insert into account(account,pwd,name) values('%s','%s','%s')",
		account.Account,
		account.Pwd,
		account.Name,
	)

	fmt.Println(sql)

	_, err = dbs.Create(sql)
	if err != nil {
		fmt.Println("failed to get session")
		resp.Errno = utils.RECODESESSIONERR
		return err
	}
	return nil
}

// 功能arr 选择的列，table表名，id索引，返回一条数据
func FindOne(arr []string, table string, id int) map[string]string {
	sql := "select "
	for k, v := range arr {
		if k == 0 {
			sql += v
		} else {
			sql += "," + v
		}
	}
	sql += " from " + table + " where id = ?"
	return dbs.FindOne(sql, id, arr)
}

// 功能分页查询
func Page(arr []string, table string, id int) []map[string]string {
	sql := "select "
	for k, v := range arr {
		if k == 0 {
			sql += v
		} else {
			sql += "," + v
		}
	}
	sql += " from " + table + " where id = ?"
	return dbs.Page(sql, id, arr)
}

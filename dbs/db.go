package dbs

import (
	"database/sql"
	"fmt"
	"irisORM/configs"

	_ "github.com/go-sql-driver/mysql"
)

type (
	Account struct {
		Account string `json:"account"`
		Pwd     string `json:"pwd"`
		Name    string `json:"name"`
	}
)

//数据库连接的全局变量
var DBConn *sql.DB

//init函数是本包被其他文件引用时自动执行，并且整个工程只会执行一次
func init() {
	fmt.Println("call dbs.Init", configs.Config)
	DBConn = InitDB(configs.Config.Db.Connstr, configs.Config.Db.Driver)
}

//初始化数据库连接
func InitDB(connstr, Driver string) *sql.DB {
	db, err := sql.Open(Driver, connstr)
	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return db
}

//通用查询，返回map嵌套map
func DBQuery(sql string) ([]map[string]string, int, error) {
	fmt.Println("query is called:", sql)
	rows, err := DBConn.Query(sql)
	if err != nil {
		fmt.Println("query data err", err)
		return nil, 0, err
	}
	//得到列名数组
	cols, err := rows.Columns()
	//获取列的个数
	colCount := len(cols)
	values := make([]string, colCount)
	oneRows := make([]interface{}, colCount)
	for k, _ := range values {
		oneRows[k] = &values[k] //将查询结果的返回地址绑定，这样才能变参获取数据
	}
	//存储最终结果
	results := make([]map[string]string, 1)
	idx := 0
	//循环处理结果集
	for rows.Next() {
		rows.Scan(oneRows...)
		rowmap := make(map[string]string)
		for k, v := range values {
			rowmap[cols[k]] = v

		}
		if idx > 0 {
			results = append(results, rowmap)
		} else {
			results[0] = rowmap
		}
		idx++
		//fmt.Println(values)
	}
	//fmt.Println("---------------------------------------")
	fmt.Println("query..idx===", idx)
	return results, idx, nil

}

// dgsg
func Create(sql string) (int64, error) {
	res, err := DBConn.Exec(sql)
	if err != nil {
		fmt.Println("exec sql err,", err, "sql is ", sql)
		return -1, err
	}
	return res.LastInsertId()
}

// FindOne
func FindOne(sql string, id int, arr []string) map[string]string {
	rows := DBConn.QueryRow(sql, id)
	length := len(arr)
	values := make([]string, length)
	oneRows := make([]interface{}, length)
	for k, _ := range values {
		oneRows[k] = &values[k]
	}
	rows.Scan(oneRows...)
	rowmap := make(map[string]string)
	for k, v := range values {
		rowmap[arr[k]] = v

	}
	return rowmap
}

// Page
func Page(sql string, id int, arr []string) []map[string]string {
	rows, _ := DBConn.Query(sql, id)
	length := len(arr)
	values := make([]string, length)
	oneRows := make([]interface{}, length)
	for k, _ := range values {
		oneRows[k] = &values[k]
	}
	results := make([]map[string]string, 1)
	idx := 0
	for rows.Next() {
		rows.Scan(oneRows...)
		rowmap := make(map[string]string)
		for k, v := range values {
			rowmap[arr[k]] = v
		}
		if idx > 0 {
			results = append(results, rowmap)
		} else {
			results[0] = rowmap
		}
		idx++
	}
	return results
}

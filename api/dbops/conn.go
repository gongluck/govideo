/*
 * @Author: gongluck
 * @Date: 2020-05-22 15:16:56
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-05-22 16:10:32
 */

package dbops

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:123456@/video_server?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}

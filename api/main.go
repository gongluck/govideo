/*
 * @Author: gongluck
 * @Date: 2020-05-08 11:49:21
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-05-08 12:42:29
 */

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8000", r)
}

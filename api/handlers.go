/*
 * @Author: gongluck
 * @Date: 2020-05-08 11:54:35
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-05-08 12:07:08
 */

package main

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// CreateUser create user handler
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "Create User Handler")
}

// Login login handler
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}

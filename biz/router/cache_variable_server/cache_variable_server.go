// Code generated by hertz generator. DO NOT EDIT.

package cache_variable_server

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	cache_variable_server "qinglin.org/cache_variable/biz/handler/cache_variable_server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.GET("/delete_value", append(_deletevalueMw(), cache_variable_server.DeleteValue)...)
	root.GET("/get_value", append(_getvalueMw(), cache_variable_server.GetValue)...)
	root.GET("/set_value", append(_setvalueMw(), cache_variable_server.SetValue)...)
}
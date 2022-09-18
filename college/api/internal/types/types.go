// Code generated by goctl. DO NOT EDIT.
package types

type CollegeRequest struct {
	Id       int64  `path:"id"`
	Name     string `path:"name"`
	Code     string `path:"code"`
	Level    int8   `path:"level"`
	Province string `path:"province"`
	City     string `path:"city"`
	District string `path:"district"`
}

type CollegeResponse struct {
	Message string `json:"message"`
}
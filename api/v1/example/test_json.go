package example

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"micro-wiki/model/common/response"
)

type JsonApi struct {
}

type JsonReq struct {
	Name       string     `json:"name"`
	Id         int        `json:"id"`
	TestStruct TestStruct `json:"test_struct" form:"test_struct"`
}

type TestStruct struct {
	TestStr string `json:"test_str"`
}

func (j *JsonApi) JsonApiFunc(c *gin.Context) {
	var jsonReq JsonReq
	if err := c.BindJSON(&jsonReq); err != nil {
		// handle error
	}
	fmt.Println(jsonReq.Name)
	fmt.Println(jsonReq.Id)

	fmt.Println(jsonReq.TestStruct)
	//fmt.Println(jsonReq.testStruct.id)
	//fmt.Println(jsonReq.testStruct.testStr)
	response.Ok(c)
}

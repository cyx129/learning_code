package model

type User struct {
	Username string `json:"name" form:"name"`
	Age      uint8  `json:"age" form:"age"`
	Mobile   string `json:"mobile" form:"mobile"`
	Sex      string `json:"sex" form:"sex"`
	Address  string `json:"address" form:"address"`
	Id       uint16 `json:"id" form:"id"`
}

/**
{
	"id":5,
	"name":"xxw5",
	"mobile":"13888888888",
	"age":18,
	"address":"北京",
	"Sex":"男"
}
*/
package models

import "fmt"

type BModel struct {
	Param string `json:"param"`
	MSG   string `json:"msg"`
}

// 结构体标准输出函数，一般不需要自定义
func (this *BModel) String() string {
	return fmt.Sprintf("param: %s, msg: %s", this.Param, this.MSG)
}

func InitB(param string) *BModel {
	return &BModel{param, "get b"}
}

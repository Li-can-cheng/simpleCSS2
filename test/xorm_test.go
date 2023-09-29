package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"simpleCSS2/core/models"
	"testing"
	"xorm.io/xorm"
)

func TestXormTest(t *testing.T) {

	//与数据库建立连接
	engine, err := xorm.NewEngine("mysql", "root:1234@tcp(127.0.0.1:3306)/cloud-disk?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		t.Fatal(err)
	}

	//对数据进行切片
	data := make([]*models.UserBasic, 0)
	err = engine.Find(&data)
	if err != nil {
		t.Fatal(err)
	}

	//转成json格式
	b, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	//美化并输出
	dst := new(bytes.Buffer)
	err = json.Indent(dst, b, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(dst.String())

}

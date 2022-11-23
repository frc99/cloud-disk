package test

import (
	"bytes"
	"cloud-disk/core/models"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"xorm.io/xorm"
)

func TestXorm(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", "root:541999@/cloud-disk?charset=utf8")
	if err != nil {
		t.Fatal(err)
	}
	data := make([]*models.UserBasic, 0)
	err = engine.Find(&data)
	if err != nil {
		t.Fatal(err)
	}
	//b:[]byte
	b, err := json.Marshal(data)
	if err != nil {
		return
	}
	dst := new(bytes.Buffer)
	err = json.Indent(dst, b, "", "")
	//fmt.Println(data)
	if err != nil {
		return
	}
	fmt.Println(dst.String())

}

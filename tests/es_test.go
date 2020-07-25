package test

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"movie/services/elasticsearch"
	"path/filepath"
	"runtime"
	"testing"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

type ResData struct {
	Title string
}

func TestEs(t *testing.T) {
	Search()
	//Add()
	//Update()
	//Delete()
}

func Search() {
	sort := []map[string]string{map[string]string{"id": "desc"}}
	query := map[string]interface{}{
		"bool": map[string]interface{}{
			"must": map[string]interface{}{
				"term": map[string]interface{}{
					"title": "商品",
				},
			},
		},
	}
	res := elasticsearch.EsSearch("es_demo", query, 0, 10, sort)
	total := res.Total
	var resData []ResData
	for _, value := range res.Hits {
		var data ResData
		if err := json.Unmarshal([]byte(value.Source), &data); err != nil {
			fmt.Println("failed")
		}
		resData = append(resData, data)
	}
	fmt.Println(total)
	fmt.Println(resData)
}

func Add() {
	body := map[string]interface{}{
		"id":    3,
		"title": "张商品三",
	}
	res := elasticsearch.EsAdd("es_demo", "user-1", body)
	if res {
		fmt.Println("success")
	} else {
		fmt.Println("failed")
	}
}

func Update() {
	body := map[string]interface{}{
		"id":    1,
		"title": "李三",
	}
	res := elasticsearch.EsUpdate("es_demo", "user-1", body)
	if res {
		fmt.Println("success")
	} else {
		fmt.Println("failed")
	}
}

func Delete() {
	res := elasticsearch.EsDelete("es_demo", "user-1")
	if res {
		fmt.Println("success")
	} else {
		fmt.Println("failed")
	}
}

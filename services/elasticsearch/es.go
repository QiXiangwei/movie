package elasticsearch

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
)

type TotalData struct {
	Value    int
	Relation string
}

type HitsJsonData struct {
	Source json.RawMessage `json:"_source"`
}

type HitsData struct {
	Total TotalData      `json:"total"`
	Hits  []HitsJsonData `json:"hits"`
}

type ReqSearchData struct {
	Hits HitsData `json:"hits"`
}

func EsSearch(indexName string, query map[string]interface{}, from int, size int, sort []map[string]string) HitsData {
	var (
		err error
		str string
		std ReqSearchData
	)
	searchQuery := map[string]interface{}{
		"query": query,
		"from":  from,
		"size":  size,
		"sort":  sort,
	}
	req := httplib.Post("http://127.0.0.1:9200/" + indexName + "/_search")
	_, err = req.JSONBody(searchQuery)
	if str, err = req.String(); err != nil {
		fmt.Println(err.Error())
	}
	err = json.Unmarshal([]byte(str), &std)
	return std.Hits
}

func EsAdd(indexName string, id string, body map[string]interface{}) bool {
	var (
		err error
	)
	req := httplib.Post("http://127.0.0.1:9200/" + indexName + "/_doc/" + id)
	_, err = req.JSONBody(body)
	if _, err = req.String(); err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func EsUpdate(indexName string, id string, body map[string]interface{}) bool {
	var (
		err error
	)
	bodyData := map[string]interface{}{
		"doc": body,
	}
	req := httplib.Post("http://127.0.0.1:9200/" + indexName + "/_doc/" + id + "/_update")
	_, err = req.JSONBody(bodyData)
	if _, err = req.String(); err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func EsDelete(indexName string, id string) bool {
	var (
		err error
	)
	req := httplib.Delete("http://127.0.0.1:9200/" + indexName + "/_doc/" + id)
	if _, err = req.String(); err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

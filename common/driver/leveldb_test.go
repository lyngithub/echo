package driver

import (
	"encoding/json"
	"fmt"
	"testing"
)

var (
	db  *LevelDB
	err error
)

func init() {
	db, err = CreateLevelDB("./leveldb_data")
	if err != nil {
		panic(err)
	}
}

func TestLevelDB_Put(t *testing.T) {
	type Demo1 struct {
		Phone     string
		ChannelId int
		Reason    string
	}

	d := new(Demo1)
	d.Phone = "18322323223"
	d.ChannelId = 2
	d.Reason = "流量控制"

	err = db.Put("202205270800-A01165345035721503010953", &d)

	type Demo2 struct {
		Phone     string
		ChannelId int
		Reason    string
	}

	d2 := new(Demo2)
	d2.Phone = "18745682512"
	d2.ChannelId = 1
	d2.Reason = "流量控制"

	err = db.Put("202205270800-A01165345035721503010954", &d2)
}

func TestLevelDB_Get(t *testing.T) {
	res, err := db.Get("202205270800-A01165345035721503010954")
	if err != nil {
		panic(err)
	}
	fmt.Printf("test: %s\n", res)

	res2, err := db.Get("202205270800-A01165345035721503010954")
	if err != nil {
		panic(err)
	}

	type Demo struct {
		Phone     string
		ChannelId int
		Reason    string
	}
	d := new(Demo)

	json.Unmarshal([]byte(res2), &d)

	fmt.Printf("Phone: %s\n", d.Phone)
	fmt.Printf("ChannelId: %d\n", d.ChannelId)
	fmt.Printf("Reason: %s\n", d.Reason)

}

func TestLevelDB_Has(t *testing.T) {
	if ok, _ := db.Has("202205270800-A01165345035721503010954"); !ok {
		fmt.Printf("key 不存在: %v", ok)
	}
	fmt.Printf("key 存在")
}

func TestLevelDB_SelectAll(t *testing.T) {
	iter := db.SelectAll()
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()
		fmt.Printf("k: %s\n", k)
		fmt.Printf("v: %s\n", v)
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		panic(err)
	}
}

func TestLevelDB_Delete(t *testing.T) {
	err = db.Delete("202205270800-A01165345035721503010954")
	if err != nil {
		panic(err)
	}
}

func TestLevelDB_DeleteAll(t *testing.T) {
	ok, _ := db.DeleteAll()
	if err != nil {
		panic(err)
	}
	if !ok {
		fmt.Println("删除失败")
	}
	fmt.Println("全部删除成功")
}

func TestLevelDB_SelectPrefixSubsetKeyAll(t *testing.T) {
	content, _ := db.SelectPrefixSubsetKeyAll("202205270800-")
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
}

func TestLevelDB_CountPrefixKey(t *testing.T) {
	sum, _ := db.CountPrefixSubsetKey("202205270800-")
	if err != nil {
		panic(err)
	}
	fmt.Printf("sum: %d", sum)
}

func TestLevelDB_CountAll(t *testing.T) {
	sum, _ := db.CountAll()
	if err != nil {
		panic(err)
	}
	fmt.Printf("sum: %d", sum)
}

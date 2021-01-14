package db

import (
	"reflect"

	"github.com/soafter/maggr/kit/msg"
)

type Db struct {
	DbFolder   string
	DateFormat string
	List       map[string]interface{}
}

var (
	Msg        = msg.Msg{}
	DbFolder   = "./db/"
	DateFormat = "2006/01/02/15:04:05"
)

//取数据
func (d Db) Take(dbkind string, t string, q map[string]string) []map[string]interface{} {
	//call.Db.Take("base", "settings", map[string]string{"where":"name='karla'"})
	args := []reflect.Value{reflect.ValueOf(t), reflect.ValueOf(q)}
	queryCall := reflect.ValueOf(d.List[dbkind]).MethodByName("Take").Call(args)[0]
	return queryCall.Interface().([]map[string]interface{})
}

//写数据
func (d Db) Save(dbkind string, t string, kv map[string]interface{}, u ...string) int64 {
	if len(u) > 0 {
		//这里Update
		//call.Db.Save("base", "settings", map[string]interface{}{"name":"'ginkwan'"}, "name='karla'")
		args := []reflect.Value{reflect.ValueOf(t), reflect.ValueOf(kv), reflect.ValueOf(u[0])}
		queryCall := reflect.ValueOf(d.List[dbkind]).MethodByName("Save").Call(args)[0]
		return queryCall.Interface().(int64)
	} else {
		//这里Insert
		//call.Db.Save("base", "settings", map[string]interface{}{"name":"'ginkwan'"})
		args := []reflect.Value{reflect.ValueOf(t), reflect.ValueOf(kv)}
		queryCall := reflect.ValueOf(d.List[dbkind]).MethodByName("Save").Call(args)[0]
		return queryCall.Interface().(int64)
	}
}

//删数据
func (d Db) Delete(dbkind string, t, w string) bool {
	//call.Db.Delete("base", "settings", "name='karla'")
	args := []reflect.Value{reflect.ValueOf(t), reflect.ValueOf(w)}
	queryCall := reflect.ValueOf(d.List[dbkind]).MethodByName("Delete").Call(args)[0]
	return queryCall.Interface().(bool)
}

//查询语句
func (d Db) Query(dbkind string, q string) []map[string]string {
	args := []reflect.Value{reflect.ValueOf(q)}
	queryCall := reflect.ValueOf(d.List[dbkind]).MethodByName("Query").Call(args)[0]
	return queryCall.Interface().([]map[string]string)
}

//执行语句
func (d Db) Exec(dbkind string, q string) interface{} {
	args := []reflect.Value{reflect.ValueOf(q)}
	execCall := reflect.ValueOf(d.List[dbkind]).MethodByName("Exec").Call(args)[0]
	return execCall.Interface()
}

//执行存储过程
func (d Db) Proc(dbkind string, q string) interface{} {
	args := []reflect.Value{reflect.ValueOf(q)}
	execCall := reflect.ValueOf(d.List[dbkind]).MethodByName("Proc").Call(args)[0]
	return execCall.Interface()
}

//打开数据库
func (d Db) Open(dbkind string) {
	reflect.ValueOf(d.List[dbkind]).MethodByName("Open").Call(nil)
}

//关闭数据库
func (d Db) Close(dbkind string, q map[string]string) {
	reflect.ValueOf(d.List[dbkind]).MethodByName("Close").Call(nil)
}

//重新打开数据库
func (d Db) ReOpen(dbkind string, q map[string]string) []map[string]string {
	var re = make([]map[string]string, 0)
	return re
}

/*
func DbInit() {

}
*/

func Init(cfg Db) Db {
	var (
		setDbFolder   *string = &DbFolder
		setDateFormat *string = &DateFormat
	)
	if len(cfg.DbFolder) > 0 {
		*setDbFolder = cfg.DbFolder
	}
	if len(cfg.DateFormat) > 0 {
		*setDateFormat = cfg.DateFormat
	}
	cfg.List = make(map[string]interface{})
	//Msg.Print("db folder：" + DbFolder)

	return cfg
}

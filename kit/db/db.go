package db

import (
	"os"

	"github.com/soafter/maggr/kit/msg"
)

type Db struct {
	Msg msg.Msg
}

type Config struct {
	DbFolder string
}

func (d Db) Take(dbkind interface{}, q map[string]string) []map[string]string {
	var re = make([]map[string]string, 0)
	d.Msg.Debug(dbkind)
	return re
}

func (d Db) Save(dbkind interface{}, q map[string]string) []map[string]string {
	var re = make([]map[string]string, 0)
	return re
}
func (d Db) Delete(dbkind interface{}, q map[string]string) []map[string]string {
	var re = make([]map[string]string, 0)
	return re
}
func (d Db) Query(dbkind interface{}, q map[string]string) []map[string]string {
	var re = make([]map[string]string, 0)
	return re
}
func (d Db) Exec(dbkind interface{}, q map[string]string) []map[string]string {
	var re = make([]map[string]string, 0)
	return re
}
func (d Db) Proc(dbkind interface{}, q map[string]string) []map[string]string {
	var re = make([]map[string]string, 0)
	return re
}

func (d Db) Open(dbkind interface{}, q map[string]string) []map[string]string {
	var re = make([]map[string]string, 0)
	return re
}

func (d Db) Close(dbkind interface{}, q map[string]string) []map[string]string {
	var re = make([]map[string]string, 0)
	return re
}

func (d Db) ReOpen(dbkind interface{}, q map[string]string) []map[string]string {
	var re = make([]map[string]string, 0)
	return re
}

func Init(cfg ...Config) Db {
	var db = Db{}

	msg := msg.Init()

	dbFolder := "./db/"

	if len(cfg) > 0 {
		if len(cfg[0].DbFolder) > 0 {
			dbFolder = cfg[0].DbFolder
		}
	}

	os.MkdirAll(dbFolder, 0764)
	db.Msg = msg
	//msg.Debug(dbFolder)

	return db
}

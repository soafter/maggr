package msg

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"strings"
	"time"

	badger "github.com/dgraph-io/badger/v2"
)

//var baseDB = &ql.DBOper{DbFile: "DB/base.ql"}
//var sysSettings map[string]string
//var dbVersion = make(map[string]int)

type Msg struct {
}

type Config struct {
	DateFormat string
	DbFolder   string
}

//KV库操作预设变量
//var KV *buntdb.DB
//var KVconfig buntdb.Config

func (m Msg) Print(parm ...interface{}) {
	toTime := time.Now().Format("2006/01/02/15:04:05")
	if len(parm) > 0 {
		fmt.Printf("\033[0;36mP\033[0m\033[0;32m[\033[0m%s\033[0;32m]\033[0m%s\n", toTime, fmt.Sprint(parm...))
	} else {
		fmt.Println("no Print content")
	}
}
func (m Msg) Log(parm ...interface{}) {
	toTime := time.Now().Format("2006/01/02/15:04:05")
	rand.Seed(int64(time.Now().Unix()))
	//rsb := strconv.Itoa(rand.Intn(1000))
	//logkey := strconv.FormatInt(time.Now().UnixNano(), 10) + "." + rsb

	if len(parm) > 0 {
		fmt.Printf("\033[0;33mL\033[0m\033[0;32m[\033[0m%s\033[0;32m]\033[0m\033[0;33m%s\033[0m\n", toTime, fmt.Sprint(parm...))

		/*
			err := KV.Update(func(tx *buntdb.Tx) error {
				_, _, err := tx.Set(logkey, `{"msg":"[L]`+fmt.Sprint(parm...)+`","date":"`+logkey+`"}`, nil)
				return err
			})
			if err != nil {
				fmt.Printf("\033[0;31mE[%s]%s\033[0m\n", toTime, fmt.Sprint("SAVE LOG ERROR!!"))
			}
		*/
	} else {
		fmt.Printf("\033[0;31mE[%s]%s\033[0m\n", toTime, fmt.Sprint("No print parameters"))
	}
}

func (m Msg) Debug(parm ...interface{}) {
	toTime := time.Now().Format("2006/01/02/15:04:05")
	/*
		rand.Seed(int64(time.Now().Unix()))
		rsb := strconv.Itoa(rand.Intn(1000))
		logkey := strconv.FormatInt(time.Now().UnixNano(), 10) + "." + rsb
	*/
	//lognano := strconv.FormatInt(time.Now().UnixNano(), 10)
	msgtype := reflect.TypeOf(parm[0])

	if len(parm) > 0 {
		_, file, line, _ := m.DebugCaller(2)
		ffile := strings.Split(file, "/")
		filename := strings.Join(ffile[len(ffile)-2:], "/")
		fmt.Printf("\033[0;31mD\033[0m\033[0;32m[\033[0m%s\033[0;32m]\033[0m\033[0;31m%s(%s)\033[0m<\033[0;36m%s\033[0m>%s\n", toTime, filename, fmt.Sprint(line), msgtype, fmt.Sprint(parm...))

		/*
			err := KV.Update(func(tx *buntdb.Tx) error {
				_, _, err := tx.Set(lognano, `{"msg":"[D]`+filename+`(`+fmt.Sprint(line)+`)`+fmt.Sprint(parm...)+`","date":`+lognano+`}`, nil)
				return err
			})
			if err != nil {
				fmt.Printf("\033[0;31mE[%s]%s\033[0m\n", toTime, fmt.Sprint("SAVE LOG ERROR!!"))
			}
		*/

	} else {
		fmt.Printf("\033[0;31mE[%s]%s\033[0m\n", toTime, fmt.Sprint("No print parameters"))
	}
}

/*
func (m Msg) GetLog(parm ...interface{}) [][2]string {
	//竟然没有写参数说明？？？
	toTime := time.Now().Format("2006/01/02/15:04:05")
	var reData = make([][2]string, 0)

	if parm[0] == "test" {
		if len(parm) > 1 {

		}
	} else {
		if len(parm) == 1 {
			KV.View(func(tx *buntdb.Tx) error {
				value, err := tx.Get(parm[0].(string))
				if err != nil {
					fmt.Printf("\033[0;31mE[%s]%s\033[0m\n", toTime, err)
				}
				var setRe [2]string
				setRe[0], setRe[1] = parm[0].(string), value
				reData = append(reData, setRe)
				return nil
			})
		} else if len(parm) == 2 {
			KV.View(func(tx *buntdb.Tx) error {
				tx.AscendRange("mlog", `{"date":`+parm[0].(string)+`}`, `{"date":`+parm[1].(string)+`}`, func(k, v string) bool {
					var setRe [2]string
					setRe[0], setRe[1] = k, v
					reData = append(reData, setRe)
					return true
				})
				return nil
			})
		}
	}
	return reData
}
*/

func (m Msg) DebugCaller(skip int) (name, file string, line int, ok bool) {
	var pc uintptr
	if pc, file, line, ok = runtime.Caller(skip); !ok {
		return
	}
	name = runtime.FuncForPC(pc).Name()
	return
}

func Init(cfg ...Config) Msg {
	//是不是直接把init的内容放入来？
	var msg = Msg{}

	dateFormat := "2006/01/02/15:04:05"
	dbFolder := "./db/"

	if len(cfg) > 0 {
		if len(cfg[0].DateFormat) > 0 {
			dateFormat = cfg[0].DateFormat
		}
		if len(cfg[0].DbFolder) > 0 {
			dbFolder = cfg[0].DbFolder
		}
	}

	//要在这里判断，如果没有这个目录就建一个
	toTime := time.Now().Format(dateFormat)
	fmt.Printf("\033[0;31mT[%s]%s\033[0m\n", toTime, dbFolder)
	db, err := badger.Open(badger.DefaultOptions(dbFolder + "badger"))
	if err != nil {
		fmt.Printf("\033[0;31mE[%s]%s\033[0m\n", toTime, err)
	}
	defer db.Close()

	/*
		os.MkdirAll(dbFolder+"kv/", 0764)

		var err error = nil
		KV, err = buntdb.Open(dbFolder + "kv/logs.db")
		if err != nil {
			fmt.Printf("\033[0;31mE[%s]%s\033[0m\n", toTime, err)
		}
		KVconfig.SyncPolicy = buntdb.Always
		if err := KV.ReadConfig(&KVconfig); err != nil {
			fmt.Printf("\033[0;31mE[%s]%s\033[0m\n", toTime, err)
		}
		if err := KV.SetConfig(KVconfig); err != nil {
			fmt.Printf("\033[0;31mE[%s]%s\033[0m\n", toTime, err)
		}
		//KV.CreateIndex("mlog", "*", buntdb.Desc(buntdb.IndexJSON("date")))
		KV.CreateIndex("mlog", "*", buntdb.IndexJSON("date"))
	*/
	return msg
}

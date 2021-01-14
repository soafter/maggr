package msg

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/boltdb/bolt"
)

type Msg struct {
	DbFolder   string
	DateFormat string
}

//预设变量
var (
	Kv         *bolt.DB
	DbFolder   = "./db/"
	DateFormat = "2006/01/02/15:04:05"
)

func (m Msg) Print(parm ...interface{}) {
	toTime := time.Now().Format(DateFormat)
	if len(parm) > 0 {
		fmt.Printf("\033[0;36mP\033[0m\033[0;32m[\033[0m%s\033[0;32m]\033[0m%s\n", toTime, fmt.Sprint(parm...))
	} else {
		fmt.Println("no Print content")
	}
}
func (m Msg) Log(parm ...interface{}) {
	toTime := time.Now().Format(DateFormat)
	rand.Seed(int64(time.Now().Unix()))
	rsb := strconv.Itoa(rand.Intn(1000))
	logkey := strconv.FormatInt(time.Now().UnixNano(), 10) + "." + rsb

	if len(parm) > 0 {
		fmt.Printf("\033[0;33mL\033[0m\033[0;32m[\033[0m%s\033[0;32m]\033[0m\033[0;33m%s\033[0m\n", toTime, fmt.Sprint(parm...))
		var err error = nil
		Kv, err = bolt.Open(DbFolder+"log.db", 0600, nil)
		if err != nil {
			m.Debug(err)
		}
		defer Kv.Close()
		uerr := Kv.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("Log"))
			err := b.Put([]byte("msglog_"+logkey), []byte(`{"msg":"[L]`+fmt.Sprint(parm...)+`","date":"`+logkey+`"}`))
			return err
		})
		if uerr != nil {
			m.Debug(err)
		}
		//下面是一段读取代码，kv库的尝试暂且这样，不要纠结了。
		/*
			Kv.View(func(tx *bolt.Tx) error {
				c := tx.Bucket([]byte("Log")).Cursor()

				prefix := []byte("msglog_")
				for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {
					m.Debug(string(k))
					m.Debug(string(v))
				}
				return nil

			})
		*/

	} else {
		fmt.Printf("\033[0;31mE[%s]%s\033[0m\n", toTime, fmt.Sprint("No print parameters"))
	}
}

func (m Msg) Debug(parm ...interface{}) {
	//Debug不再记录到Log数据库里，要记住
	toTime := time.Now().Format(DateFormat)
	msgtype := reflect.TypeOf(parm[0])
	if len(parm) > 0 {
		_, file, line, _ := m.DebugCaller(2)
		ffile := strings.Split(file, "/")
		filename := strings.Join(ffile[len(ffile)-2:], "/")
		fmt.Printf("\033[0;31mD\033[0m\033[0;32m[\033[0m%s\033[0;32m]\033[0m\033[0;31m%s(%s)\033[0m<\033[0;36m%s\033[0m>%s\n", toTime, filename, fmt.Sprint(line), msgtype, fmt.Sprint(parm...))

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

/*
func (m Msg) I18n(s string) string {
	saveDate := strconv.FormatInt(time.Now().UnixNano(), 10) //新建保存数据
	reStr := s //传入的文本
	newSave := false //假设不是新文本
	if len(s) > 0 { //有文本长度
		KV.View(func(tx *buntdb.Tx) error { //打开数据库
			value, err := tx.Get(s) //取传入的文本
			if err != nil {
				reStr = s
				newSave = true
				return nil
			} else {
				if Config.Sys["Lang"] == "default" {
					reStr = s
				} else {
					reInfo := I18nText{}
					if err := json.Unmarshal([]byte(value), &reInfo); err != nil {
						Debug(err)
					}
					if len(reInfo.Lang[Config.Sys["Lang"]]) > 0 {
						reStr = reInfo.Lang[Config.Sys["Lang"]]
					} else {
						reStr = s
					}
				}
			}
			return nil
		})
	}
	if newSave {
		KV.Update(func(tx2 *buntdb.Tx) error {
			saveText := I18nText{}
			saveText.Date = saveDate
			addLang := map[string]string{"zh-CHS": "", "zh-CHT": "", "en": "", "jp": ""}
			saveText.Lang = addLang
			saveTextJson, jerr := json.Marshal(saveText)
			if jerr != nil {
				Debug(jerr)
			}
			_, _, err := tx2.Set(s, string(saveTextJson), nil)
			if err != nil {
				Debug(err)
			}
			Log("添加了一条新的多语言记录：", s)
			return nil
		})
	}
	return reStr
}
a*/

func (m Msg) DebugCaller(skip int) (name, file string, line int, ok bool) {
	var pc uintptr
	if pc, file, line, ok = runtime.Caller(skip); !ok {
		return
	}
	name = runtime.FuncForPC(pc).Name()
	return
}

func Init(cfg Msg) Msg {
	//是不是直接把init的内容放入来？
	var (
		setDbFolder   *string = &DbFolder
		setDateFormat *string = &DateFormat
		err           error   = nil
	)
	if len(cfg.DbFolder) > 0 {
		*setDbFolder = cfg.DbFolder
	}
	if len(cfg.DateFormat) > 0 {
		*setDateFormat = cfg.DateFormat
	}

	//要在这里判断，如果没有这个目录就建一个<===========待
	Kv, err = bolt.Open(DbFolder+"log.db", 0600, nil)
	if err != nil {
		cfg.Debug(err)
	}
	defer Kv.Close()
	Kv.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Log"))
		if b == nil {
			b, err = tx.CreateBucket([]byte("Log"))
			if err != nil {
				cfg.Debug(err)
			}
		}
		/*
			if err := b.Put([]byte("0"), []byte("new Bucket")); err != nil {
				cfg.Debug(err)
			}
		*/
		return nil
	})

	//cfg.Print("msg db folder：" + DbFolder)
	return cfg
}

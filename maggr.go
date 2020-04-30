package maggr

import (
	"github.com/soafter/maggr/kit/db"
	"github.com/soafter/maggr/kit/msg"
)

type Maggr struct {
	PassObfus  string
	DateFormat string
	Msg        msg.Msg
	Db         db.Db
}

type Config struct {
	DbFolder   string //文件类数据库目录路径
	DateFormat string //日期显示格式
	PassObfus  string //密码混淆字符
}

func Init(cfg ...Config) Maggr {
	var maggr = Maggr{}
	//Default
	dbFolder := "./db/"
	passObfus := "bullcalf"
	dateFormat := "2006/01/02/15:04:05"

	if len(cfg) > 0 {
		if len(cfg[0].DbFolder) > 0 {
			dbFolder = cfg[0].DbFolder
		}
		if len(cfg[0].PassObfus) > 0 {
			passObfus = cfg[0].PassObfus
		}
		if len(cfg[0].DateFormat) > 0 {
			dateFormat = cfg[0].DateFormat
		}
	}

	//Global
	maggr.PassObfus = passObfus
	maggr.DateFormat = dateFormat

	//Kit Setting
	//就这么样。。。添加一个组件要做三处处理
	maggr.Msg = msg.Init(msg.Config{DbFolder: dbFolder})
	maggr.Db = db.Init(db.Config{DbFolder: dbFolder})

	return maggr
}

package maggr

import (
	"github.com/soafter/maggr/kit/db"
	"github.com/soafter/maggr/kit/msg"
)

var (
	Cfg ConfigFile
	Ini = InitConfig{
		DbFolder:   "./db/",
		PassObfus:  "bullcalf",
		DateFormat: "2006/01/02/15:04:05",
		Version:    "2.0.0"}
	Msg = msg.Msg{}
	Db  = db.Db{}
)

func Setting(cfg InitConfig) error {
	var (
		err           error   = nil
		setDbFolder   *string = &Ini.DbFolder
		setPassObfus  *string = &Ini.DateFormat
		setDateFormat *string = &Ini.PassObfus
	)
	if len(cfg.DbFolder) > 0 {
		*setDbFolder = cfg.DbFolder
	}
	if len(cfg.PassObfus) > 0 {
		*setPassObfus = cfg.PassObfus
	}
	if len(cfg.DateFormat) > 0 {
		*setDateFormat = cfg.DateFormat
	}

	var msgSetting *msg.Msg = &Msg
	*msgSetting = msg.Init(msg.Msg{
		DbFolder:   Ini.DbFolder,
		DateFormat: Ini.DateFormat})

	var dbSetting *db.Db = &Db
	*dbSetting = db.Init(db.Db{
		DbFolder:   Ini.DbFolder,
		DateFormat: Ini.DateFormat})
	return err
}

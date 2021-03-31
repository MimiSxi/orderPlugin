package model

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	logger logrus.FieldLogger
}

func (g *Logger) Print(v ...interface{}) {
	switch v[0] {
	case "sql":
		g.logger.WithFields(
			logrus.Fields{
				"module":   "gorm",
				"type":     "sql",
				"rows":     v[5],
				"duration": v[2],
				"src_ref":  v[1],
				"values":   v[4],
			},
		).Debug(v[3])
	case "log":
		g.logger.WithFields(logrus.Fields{"module": "gorm", "type": "log"}).Print(v[2])
	}
}

var db *gorm.DB

func NewDB(gdb *gorm.DB) (err error) {
	db = gdb
	db.LogMode(true)
	db.SetLogger(&Logger{logrus.StandardLogger()})
	db.AutoMigrate(&OrderInfo{})
	db.AutoMigrate(&AliPayInfo{})
	return nil
}

func OpenDatabase(Driver string,
	DSN string,
	Idle int,
	Active int,
	IdleTimeout time.Duration) (err error) {

	db, err = gorm.Open(Driver, DSN)
	if err != nil {
		logrus.Errorf("Open database failed: %v", err)
		return err
	}

	if err = db.DB().Ping(); err != nil {
		return err
	}
	db.DB().SetMaxIdleConns(Idle)
	db.DB().SetMaxOpenConns(Active)
	db.DB().SetConnMaxLifetime(time.Hour)
	db.SingularTable(true)

	return NewDB(db)
}

func Run(f func(*gorm.DB)) {
	f(db)
}

type database struct {
	User     string `json:user`
	Password string `json:password`
	Host     string `json:host`
	Port     int    `json:port`
	Database string `json:database`
}

type AliProp struct {
	AppId       string `json:appid`
	Url         string `json:url`
	RedirectUrl string `json:redirectUrl`
}

type config struct {
	Db  database `json:db`
	Ali AliProp  `json:ali`
}

var AppId string
var Url string
var RedirectUrl string

// 获取配置
func Setconfig(jsonStr string) error {
	con := config{}
	err := json.Unmarshal([]byte(jsonStr), &con)
	AppId = con.Ali.AppId
	Url = con.Ali.Url
	RedirectUrl = con.Ali.RedirectUrl
	return err
}

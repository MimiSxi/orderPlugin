package orderPlugin

import (
	"encoding/json"
	"fmt"
	"github.com/Fiber-Man/orderPlugin/model"
	"github.com/Fiber-Man/orderPlugin/schema"
	"github.com/sirupsen/logrus"

	"github.com/Fiber-Man/funplugin"
	"github.com/Fiber-Man/funplugin/plugin"

	"github.com/jinzhu/gorm"
)

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

type app struct {
	pls    funplugin.PluginManger
	config config
}

func (a *app) Version() string {
	return "order version"
}

func (a *app) Init(db *gorm.DB) error {
	if db != nil {
		model.NewDB(db)
	} else {
		strDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", a.config.Db.User, a.config.Db.Password, a.config.Db.Host, a.config.Db.Port, a.config.Db.Database)
		if err := model.OpenDatabase("mysql", strDSN, 10, 100, 1000); err != nil {
			logrus.Errorf("failed to connect database: %v", err)
		}
	}
	schema.Init()
	jsonStr, err := json.Marshal(a.config)
	err = model.Setconfig(string(jsonStr))
	return err
}

func (a *app) Schema() funplugin.Schema {
	return schema.NewPlugSchema(a.pls)
}

func (a *app) Query(arg ...interface{}) (interface{}, error) {
	return nil, nil
}

func (a *app) Func(params interface{}) (interface{}, error) {
	return nil, nil
}

func (a *app) String() string {
	return "order"
}

func (a *app) Setup() error {
	model.Run(func(db *gorm.DB) {})
	return nil
}

func NewPlugin(config string, pls funplugin.PluginManger) funplugin.Plugin {
	plugin.New(pls)

	a := &app{
		pls: pls,
	}

	if err := json.Unmarshal([]byte(config), &a.config); err != nil {
		logrus.Errorf("error +v%", err)
		return nil
	}

	return a
}

package model

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/smartwalle/alipay/v3"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// "github.com/smartwalle/alipay/v3"

const (
	//appID      = "2021002132662399"
	privateKey = "MIIEpAIBAAKCAQEA5xKoySfUg3hLIZtlGy+L0eLGtumEbv0hilGEwaYV5DAN4zPyx5eteKvgZimdXrRYnDvZ4mh49bbua1+XnChTxNLv4w4qUfBzSlnGmJDmyV/N7129lRaIvYYMiFOLyYBDPVODCkOZwpha5NU2WaZfVT6tvofV3mA7wWeu6mGCOaTytpA9/fTReSNZFBPluLV06x4UcZki4AEEXdycec+NNoJQ/eul4foyQCEItzDyOmrSYbkf/2UIifDGduC+xtkmqpkuMJejlRTLwUPY20p5CPmynyp+odzCUpmm1xFfrEtKtKXOFJmsBfr4K/zGAze5yJx2/IIbZhtoN15IcvOGewIDAQABAoIBAQC8BNTOCNjEuQb5K4ZTXpa4i3wBrXUTEmlOMRKCt2+souVJ8CUl/ucp/0CyID5qpvhK9/BMZ5G07cqGF9w3NiEjUDfdWtNYpPxKjU4pKg5/4LKiiHYQb6uH+yELdF+T8AfGSMOhgGwGiQ28kTiOLe/4Xu3k0IZXUZqNvp33HKxn1aHR35IUFjt8g1nBLLtUEJ6PNZuC4LUOrTXBK8GnbTy5l0RRYEuJklOigPsRhu9rgG23qDqEnW7ObDMkwh3Wkx84+s6S8SugJ8Nfw3QJ66QA14Exq9c99uVcoMFP7mArGKXYAld4d8HgLAippT5YEKWXG94vZy9itAO+YCKcZ9uRAoGBAPOpHedLqisug6Iny2u9zypfM7lInmWNR6PzaOcuC45tMlMkC9ya2ohY9JE9XTairYHuISktQU1R3pOYA7xGUQFZ8BWLP39SiM/kpyXoIYDaddU73gP5SqwYOhIDSj1NXuT2pcF/TQD9wDOtjbtOWRdwEQfMa7uLPSvudO5q0X+FAoGBAPLGWg6hbiDlk1jFGQiwtd5l6aHt2Af1CH1LdhXZyy0jOCH8wy31OZ6vSA2oRMYwxB2hn2KXdN/2xvvq6CxZ1XnL7oaLNfLYXmYmkS2q1bd+y0nxwhdGsbKWem/A1gGYLBQW42kxkwPr5/Y2B9jNW52/fhxwBBluB8mw8UIpCM3/AoGAJvqC4iFkk4vZWvNqw02V+n1IVPec/zneoAesXG8tQheN2WcGzr+m/fDdDu72Hmtfvk1N2Lx4mdni9VF4J4JIKyMsGQYxnjih0kANzS6ZTXelKfttxMz4eRdXEtKb6bqa153tXkrzEpmFSb8V0UTzU6CF2O2GvnXDz2dSJWHJKdECgYAgE4sEkdmuKQcN3ITRPB/bcZWr2nQHoR1tCJJikrMglJ2vB+l14gep6rjXbRshIIJY8+jOKvq7OKzTzha8/WWSQRqT1kLbgjD+yCu4X/D63JrZe0LMtn91/CHTMCRWc5enU9raJD2rb/jm8/6Xa5KmRg3QjhBMl9gZkvJdbnSGWwKBgQChfi3SExsfem4gUc9flnE+XzgnjLRkheZ7r/Ge8J0ZkJFBN0auQOj10XBDyAhJ71juWfY6MB1YdQbZTGBBcxwFjfXCguSFJjqJjyYiRMO/MXjbZTXYCMUfUmUndLsqp8XX3AIncnrFGJPm1+WPS+ev4lja/UwTMuqvOQed8ZtLaQ=="
)

func GetTimeTick64() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetFormatTime(time time.Time) string {
	return time.Format("20060102")
}

func GenerateCode() string {
	//todo rand????????????redis
	date := GetFormatTime(time.Now())
	r := rand.Intn(1000)
	code := fmt.Sprintf("%s%d%03d", date, GetTimeTick64(), r)
	return code
}

// ???????????????
type AliPayInfo struct {
	ID          uint        `gorm:"primary_key" gqlschema:"query!;querys;delete!" description:"??????id"`
	OrderId     uint        `gorm:"NOT NULL;" gqlschema:"create!;querys" description:"??????id"`
	Title       string      `gorm:"Type:varchar(255);DEFAULT:'';NOT NULL;" gqlschema:"create!;querys" description:"??????"`
	TradeNo     string      `gorm:"Type:varchar(255);DEFAULT:'';NOT NULL;" gqlschema:"querys" description:"??????????????????"`
	OutTradeNo  string      `gorm:"Type:varchar(255);DEFAULT:'';NOT NULL;" gqlschema:"querys" description:"?????????????????????"`
	SellerId    string      `gorm:"Type:varchar(255);DEFAULT:'';NOT NULL;" gqlschema:"querys" description:"????????????????????????"`
	TotalAmount string      `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"querys" description:"???????????????????????????????????????????????????????????????????????????[0.01,100000000]"`
	TradeStatus TradeStatus `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"querys" description:"????????????"`
	PayTime     time.Time   `gorm:"DEFAULT:'1970-1-1 00:00:00';" description:"????????????" gqlschema:"querys"`
	CreatedAt   time.Time   `description:"????????????" gqlschema:"querys"`
	UpdatedAt   time.Time   `description:"????????????" gqlschema:"querys"`
	DeletedAt   *time.Time
	v2          int    `gorm:"-" exclude:"true"`
	PayURL      string `gorm:"-"`
}

type AliPayInfos struct {
	TotalCount int
	Edges      []AliPayInfo
}

func (o AliPayInfo) Query(params graphql.ResolveParams) (AliPayInfo, error) {
	p := params.Args
	err := db.Where(p).First(&o).Error
	return o, err
}

func (o AliPayInfo) Querys(params graphql.ResolveParams) (AliPayInfos, error) {
	var result AliPayInfos

	dbselect := GenSelet(db, params)
	dbcount := GenWhere(db.Model(o), params)

	err := dbselect.Find(&result.Edges).Error
	if err != nil {
		return result, err
	}
	err = dbcount.Count(&result.TotalCount).Error
	return result, err
}

func (o AliPayInfo) Create(params graphql.ResolveParams) (AliPayInfo, error) {
	par := params.Args
	client, err := alipay.New(AppId, privateKey, true)
	client.LoadAliPayPublicKey("MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwER3MvLuMWj2Eh3c2ZC5rsqyh0B1BuYJb2+Ok4gWf9kTACkn3lihpWa4AhgTa7XKoEH7JvDPvCG9l9pSOyvlBIC4XpLoMA58Uw2LUlQt+4iZQ5zYtjMuwBE0fjpCaWrNBo1KlSZ20H3GIqIfJjwFZF+ArM32/SZ+snptg1uaW5ClJSak0QIkGIt3MQmPogePhl3AaT7Hmhk/gAyYuAMATLMzmV5HfbBqjUst9tg46W8zLOr4p1eiFQPZpn9cwyvbmRITW7qeSwoH2mlzewQelxP5o6F7eMytvBkWLoN7TjtESq7u7MOEiqwWZjoeI6jO5AuSTtFEoYh4h0eGLpKC8wIDAQAB")
	// ??? key ?????????????????????????????????
	if err != nil {
		return o, err
	}
	var p = alipay.TradePagePay{}
	o.Title = par["title"].(string)
	p.NotifyURL = Url + "/alireturn"
	p.ReturnURL = Url + "/alipay"
	p.Subject = o.Title
	p.OutTradeNo = GenerateCode()
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	o.OrderId = uint(par["orderId"].(int))
	order := &OrderInfo{}
	err = db.Where("id = ?", o.OrderId).First(&order).Error
	if err != nil {
		return o, err
	}
	TotalAmount := order.GoodsPrice + order.FreightPrice
	o.TotalAmount = strconv.FormatFloat(TotalAmount, 'f', 2, 64)
	p.TotalAmount = o.TotalAmount
	notfound := db.Where("order_id = ?", o.OrderId).First(&AliPayInfo{}).RecordNotFound()
	if !notfound {
		// orderid??????
		err = db.Where("order_id = ?", o.OrderId).First(&o).Error
		if err != nil {
			return o, err
		}
		p.OutTradeNo = o.OutTradeNo
	}
	o.OutTradeNo = p.OutTradeNo
	url, err := client.TradePagePay(p)
	if err != nil {
		fmt.Println(err)
	}
	o.PayURL = url.String()
	if !notfound {
		err = db.Save(&o).Error
		return o, err
	}
	o.TradeStatus = TradeStatusWaitBuyerPay
	err = db.Create(&o).Error
	return o, err
}

// ??????????????????
func Alireturn(c *gin.Context) {
	client, err := alipay.New(AppId, privateKey, true)
	client.LoadAliPayPublicKey("MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwER3MvLuMWj2Eh3c2ZC5rsqyh0B1BuYJb2+Ok4gWf9kTACkn3lihpWa4AhgTa7XKoEH7JvDPvCG9l9pSOyvlBIC4XpLoMA58Uw2LUlQt+4iZQ5zYtjMuwBE0fjpCaWrNBo1KlSZ20H3GIqIfJjwFZF+ArM32/SZ+snptg1uaW5ClJSak0QIkGIt3MQmPogePhl3AaT7Hmhk/gAyYuAMATLMzmV5HfbBqjUst9tg46W8zLOr4p1eiFQPZpn9cwyvbmRITW7qeSwoH2mlzewQelxP5o6F7eMytvBkWLoN7TjtESq7u7MOEiqwWZjoeI6jO5AuSTtFEoYh4h0eGLpKC8wIDAQAB")
	// ??? key ?????????????????????????????????
	if err != nil {
		fmt.Println(err)
	}
	c.Request.ParseForm()
	ok, err := client.VerifySign(c.Request.Form)
	fmt.Println(ok, err)
	return
}

// ????????????????????????
func Alipay(c *gin.Context) {
	client, err := alipay.New(AppId, privateKey, true)
	client.LoadAliPayPublicKey("MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwER3MvLuMWj2Eh3c2ZC5rsqyh0B1BuYJb2+Ok4gWf9kTACkn3lihpWa4AhgTa7XKoEH7JvDPvCG9l9pSOyvlBIC4XpLoMA58Uw2LUlQt+4iZQ5zYtjMuwBE0fjpCaWrNBo1KlSZ20H3GIqIfJjwFZF+ArM32/SZ+snptg1uaW5ClJSak0QIkGIt3MQmPogePhl3AaT7Hmhk/gAyYuAMATLMzmV5HfbBqjUst9tg46W8zLOr4p1eiFQPZpn9cwyvbmRITW7qeSwoH2mlzewQelxP5o6F7eMytvBkWLoN7TjtESq7u7MOEiqwWZjoeI6jO5AuSTtFEoYh4h0eGLpKC8wIDAQAB")
	// ??? key ?????????????????????????????????
	if err != nil {
		fmt.Println(err)
	}
	var noti, _ = client.GetTradeNotification(c.Request)
	ali := &AliPayInfo{}
	err = db.Where("out_trade_no = ?", noti.OutTradeNo).First(&ali).Error
	if err != nil {
		return
	}
	ali.PayTime = time.Now()
	ali.SellerId = noti.SellerId
	ali.TotalAmount = noti.TotalAmount
	ali.TradeStatus = TradeStatusSuccess
	ali.TradeNo = noti.TradeNo
	err = db.Save(&ali).Error
	order := &OrderInfo{}
	db.Where("id = ?", ali.OrderId).First(&order)
	order.PayTime = time.Now()
	order.PaymentId = ali.ID
	order.PayWay = ALIPAY
	order.Status = TO_BE_DELIVER
	err = db.Save(&order).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	if RedirectUrl != "nil" {
		redirect := fmt.Sprintf("%s/orderId-%v?type=TO_BE_DELIVER", RedirectUrl, ali.OrderId)
		c.Redirect(http.StatusMovedPermanently, redirect)
		return
	}
	alipay.AckNotification(c.Writer) // ????????????????????????
	return
}

func (o AliPayInfo) Delete(params graphql.ResolveParams) (AliPayInfo, error) {
	v, ok := params.Source.(AliPayInfo)
	if !ok {
		return o, errors.New("delete param")
	}
	if len(v.TradeNo) > 0 {
		return v, errors.New("???????????????,????????????")
	}
	err := db.Delete(&v).Error
	return v, err
}

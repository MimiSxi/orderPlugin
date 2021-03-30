package model

import (
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/smartwalle/alipay/v3"
	"time"
)

const (
	appID      = "2021000117621625"
	privateKey = "MIIEpAIBAAKCAQEA5xKoySfUg3hLIZtlGy+L0eLGtumEbv0hilGEwaYV5DAN4zPyx5eteKvgZimdXrRYnDvZ4mh49bbua1+XnChTxNLv4w4qUfBzSlnGmJDmyV/N7129lRaIvYYMiFOLyYBDPVODCkOZwpha5NU2WaZfVT6tvofV3mA7wWeu6mGCOaTytpA9/fTReSNZFBPluLV06x4UcZki4AEEXdycec+NNoJQ/eul4foyQCEItzDyOmrSYbkf/2UIifDGduC+xtkmqpkuMJejlRTLwUPY20p5CPmynyp+odzCUpmm1xFfrEtKtKXOFJmsBfr4K/zGAze5yJx2/IIbZhtoN15IcvOGewIDAQABAoIBAQC8BNTOCNjEuQb5K4ZTXpa4i3wBrXUTEmlOMRKCt2+souVJ8CUl/ucp/0CyID5qpvhK9/BMZ5G07cqGF9w3NiEjUDfdWtNYpPxKjU4pKg5/4LKiiHYQb6uH+yELdF+T8AfGSMOhgGwGiQ28kTiOLe/4Xu3k0IZXUZqNvp33HKxn1aHR35IUFjt8g1nBLLtUEJ6PNZuC4LUOrTXBK8GnbTy5l0RRYEuJklOigPsRhu9rgG23qDqEnW7ObDMkwh3Wkx84+s6S8SugJ8Nfw3QJ66QA14Exq9c99uVcoMFP7mArGKXYAld4d8HgLAippT5YEKWXG94vZy9itAO+YCKcZ9uRAoGBAPOpHedLqisug6Iny2u9zypfM7lInmWNR6PzaOcuC45tMlMkC9ya2ohY9JE9XTairYHuISktQU1R3pOYA7xGUQFZ8BWLP39SiM/kpyXoIYDaddU73gP5SqwYOhIDSj1NXuT2pcF/TQD9wDOtjbtOWRdwEQfMa7uLPSvudO5q0X+FAoGBAPLGWg6hbiDlk1jFGQiwtd5l6aHt2Af1CH1LdhXZyy0jOCH8wy31OZ6vSA2oRMYwxB2hn2KXdN/2xvvq6CxZ1XnL7oaLNfLYXmYmkS2q1bd+y0nxwhdGsbKWem/A1gGYLBQW42kxkwPr5/Y2B9jNW52/fhxwBBluB8mw8UIpCM3/AoGAJvqC4iFkk4vZWvNqw02V+n1IVPec/zneoAesXG8tQheN2WcGzr+m/fDdDu72Hmtfvk1N2Lx4mdni9VF4J4JIKyMsGQYxnjih0kANzS6ZTXelKfttxMz4eRdXEtKb6bqa153tXkrzEpmFSb8V0UTzU6CF2O2GvnXDz2dSJWHJKdECgYAgE4sEkdmuKQcN3ITRPB/bcZWr2nQHoR1tCJJikrMglJ2vB+l14gep6rjXbRshIIJY8+jOKvq7OKzTzha8/WWSQRqT1kLbgjD+yCu4X/D63JrZe0LMtn91/CHTMCRWc5enU9raJD2rb/jm8/6Xa5KmRg3QjhBMl9gZkvJdbnSGWwKBgQChfi3SExsfem4gUc9flnE+XzgnjLRkheZ7r/Ge8J0ZkJFBN0auQOj10XBDyAhJ71juWfY6MB1YdQbZTGBBcxwFjfXCguSFJjqJjyYiRMO/MXjbZTXYCMUfUmUndLsqp8XX3AIncnrFGJPm1+WPS+ev4lja/UwTMuqvOQed8ZtLaQ=="
)

// 支付宝付款
type AliPayInfo struct {
	ID        uint      `gorm:"primary_key" gqlschema:"query!;querys" description:"订单id"`
	OrderId   uint      `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create!;querys" description:"订单id"`
	CreatedAt time.Time `description:"创建时间" gqlschema:"querys"`
	UpdatedAt time.Time `description:"更新时间" gqlschema:"querys"`
	DeletedAt *time.Time
	v2        int `gorm:"-" exclude:"true"`
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
	q := params.Args
	client, err := alipay.New(appID, privateKey, false)
	client.LoadAliPayPublicKey("aliPublicKey")
	// 将 key 的验证调整到初始化阶段
	if err != nil {
		fmt.Println(err)
		return o,err
	}

	var p = alipay.TradePagePay{}
	p.NotifyURL = "http://xxx"
	p.ReturnURL = "http://xxx"
	p.Subject = "标题"
	p.OutTradeNo = fmt.Sprintf("order_%d", q["orderId"].(int))
	p.TotalAmount = "10.00"
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, err := client.TradePagePay(p)
	if err != nil {
		fmt.Println(err)
	}

	var payURL = url.String()
	fmt.Println(payURL)
	// 这个 payURL 即是用于支付的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。


	return o, errors.New(payURL)
}

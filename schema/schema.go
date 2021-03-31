package schema

import (
	"errors"
	"github.com/Fiber-Man/funplugin"
	"github.com/Fiber-Man/funplugin/plugin"
	"github.com/Fiber-Man/orderPlugin/model"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

var orderInfoSchema *funplugin.ObjectSchema
var alipaySchema *funplugin.ObjectSchema

var load = false

func Init() {
	union, err := plugin.AutoField("albumorder,")
	if err != nil {
		panic(errors.New("not have object type"))
	}
	orderInfoSchema.GraphQLType.AddFieldConfig("goods", union)
}

func marge(oc *funplugin.ObjectSchema) {
	for k, v := range oc.Query {
		queryFields[k] = v
	}
	for k, v := range oc.Mutation {
		mutationFields[k] = v
	}
}

var queryFields = graphql.Fields{
	// "account":  &queryAccount,
	// "accounts": &queryAccountList,
	// "authority":  &queryAuthority,
	// "authoritys": &queryAuthorityList,
}

var mutationFields = graphql.Fields{
	// "createAccount": &createAccount,
	// "updateAccount": &updateAccount,
}

// NewSchema 用于插件主程序调用
func NewPlugSchema(pls funplugin.PluginManger) funplugin.Schema {
	if load != true {

		orderInfoSchema, _ = pls.NewSchemaBuilder(model.OrderInfo{})
		marge(orderInfoSchema)

		alipaySchema, _ = pls.NewSchemaBuilder(model.AliPayInfo{})
		marge(alipaySchema)

		load = true
	}

	// roleSchema, _ := pls.NewSchemaBuilder(model.Role{})
	// marge(roleSchema)

	// roleAccountSchema, _ := pls.NewSchemaBuilder(model.RoleAccount{})
	// marge(roleAccountSchema)
	return funplugin.Schema{
		Object: map[string]*graphql.Object{
			// "account": accountType,

			"orderInfo":  orderInfoSchema.GraphQLType,
			"alipayInfo": alipaySchema.GraphQLType,

			// "role":        roleSchema.GraphQLType,
			// "roleaccount": roleAccountSchema.GraphQLType,
		},
		Query:    queryFields,
		Mutation: mutationFields,
		Get: map[string]func(c *gin.Context){
			"/alireturn": model.Alireturn,
			"/alipay":    model.Alipay,
		},
	}
}

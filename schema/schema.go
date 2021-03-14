package schema

import (
	"errors"
	"github.com/Fiber-Man/funplugin"
	"github.com/Fiber-Man/funplugin/plugin"
	"github.com/Fiber-Man/orderPlugin/model"
	"github.com/graphql-go/graphql"
)

var orderInfoSchema *funplugin.ObjectSchema

var load = false

func Init() {
	//union, err := plugin.AutoField("albumOrder,nothing")
	//if err != nil {
	//	panic(errors.New("not have object type"))
	//}
	//orderInfoSchema.GraphQLType.AddFieldConfig("goods", union)

	union, err := plugin.AutoField("albumOrder")
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

		load = true
	}

	// roleSchema, _ := pls.NewSchemaBuilder(model.Role{})
	// marge(roleSchema)

	// roleAccountSchema, _ := pls.NewSchemaBuilder(model.RoleAccount{})
	// marge(roleAccountSchema)

	return funplugin.Schema{
		Object: map[string]*graphql.Object{
			// "account": accountType,

			"orderInfo": orderInfoSchema.GraphQLType,

			// "role":        roleSchema.GraphQLType,
			// "roleaccount": roleAccountSchema.GraphQLType,
		},
		Query:    queryFields,
		Mutation: mutationFields,
	}
}

package role

import (
	"api/db"
	"api/logs"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

var Adapter *gormadapter.Adapter
var Enforcer *casbin.Enforcer

var rbac_model = `[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act`

func init() {
	if err := InitCasBin(db.DB); err != nil {
		logs.Log.Fatal(err.Error())
		return
	}
}

func InitCasBin(database *gorm.DB) error {
	var err error
	//gormadapter.TurnOffAutoMigrate(database)
	//database.AutoMigrate(&gormadapter.CasbinRule{})
	Adapter, err = gormadapter.NewAdapterByDB(database)
	if err != nil {
		logs.Log.Error(err.Error())
		return err
	}
	fromString, err := model.NewModelFromString(rbac_model)
	if err != nil {
		logs.Log.Error(err.Error())
		return err
	}
	if Enforcer, err = casbin.NewEnforcer(fromString, Adapter); err != nil {
		logs.Log.Error(err.Error())
		return err
	}

	return nil
}

package models
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/lunny/xorm"
	"github.com/revel/revel"
)

var  engine *xorm.Engine

func init() {
	var err error
	engine,err = xorm.NewEngine("mysql","root:@/xing100?charset=utf8")
	if err != nil {
		revel.ERROR.Fatalln("mysql connect error:",err)
	}
	engine.SetTableMapper(xorm.NewPrefixMapper(xorm.SnakeMapper{},"xing100b2c_"))
	engine.SetColumnMapper(xorm.SnakeMapper{})
	engine.ShowSQL=true
}
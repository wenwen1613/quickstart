package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["quickstart/controllers:UserController"] = append(beego.GlobalControllerRouter["quickstart/controllers:UserController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/user/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

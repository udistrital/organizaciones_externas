// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/organizaciones_externas/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/arl",
			beego.NSInclude(
				&controllers.ArlController{},
			),
		),

		beego.NSNamespace("/banco",
			beego.NSInclude(
				&controllers.BancoController{},
			),
		),

		beego.NSNamespace("/caja_compensacion",
			beego.NSInclude(
				&controllers.CajaCompensacionController{},
			),
		),

		beego.NSNamespace("/entidad_aseguradora",
			beego.NSInclude(
				&controllers.EntidadAseguradoraController{},
			),
		),

		beego.NSNamespace("/eps",
			beego.NSInclude(
				&controllers.EpsController{},
			),
		),

		beego.NSNamespace("/fondo_pension",
			beego.NSInclude(
				&controllers.FondoPensionController{},
			),
		),

		beego.NSNamespace("/sucursal",
			beego.NSInclude(
				&controllers.SucursalController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

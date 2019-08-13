package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:ArlController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:ArlController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:ArlController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:ArlController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:ArlController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:ArlController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:ArlController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:ArlController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:ArlController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:ArlController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:BancoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:BancoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:BancoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:BancoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:BancoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:CajaCompensacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:CajaCompensacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:CajaCompensacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:CajaCompensacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:CajaCompensacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:CajaCompensacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:CajaCompensacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:CajaCompensacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:CajaCompensacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:CajaCompensacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:EntidadAseguradoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:EntidadAseguradoraController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:EntidadAseguradoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:EntidadAseguradoraController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:EntidadAseguradoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:EntidadAseguradoraController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:EntidadAseguradoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:EntidadAseguradoraController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:EntidadAseguradoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:EntidadAseguradoraController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:EpsController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:EpsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:EpsController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:EpsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:EpsController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:EpsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:EpsController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:EpsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:EpsController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:EpsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:FondoPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:FondoPensionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:FondoPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:FondoPensionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:FondoPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:FondoPensionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:FondoPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:FondoPensionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:FondoPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:FondoPensionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:SucursalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:SucursalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:SucursalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:SucursalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:SucursalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:SucursalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:SucursalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:SucursalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:SucursalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizaciones_externas/controllers:SucursalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

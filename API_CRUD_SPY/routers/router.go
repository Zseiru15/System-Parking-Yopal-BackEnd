// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/Zseiru15/API_CRUD_SPY/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/Pagos",
			beego.NSInclude(
				&controllers.PagosController{},
			),
		),

		beego.NSNamespace("/Estacionamientos_Promociones",
			beego.NSInclude(
				&controllers.EstacionamientosPromocionesController{},
			),
		),

		beego.NSNamespace("/Promociones",
			beego.NSInclude(
				&controllers.PromocionesController{},
			),
		),

		beego.NSNamespace("/Credenciales",
			beego.NSInclude(
				&controllers.CredencialesController{},
			),
		),

		beego.NSNamespace("/Estacionamientos",
			beego.NSInclude(
				&controllers.EstacionamientosController{},
			),
		),

		beego.NSNamespace("/Comentarios",
			beego.NSInclude(
				&controllers.ComentariosController{},
			),
		),

		beego.NSNamespace("/Cargos",
			beego.NSInclude(
				&controllers.CargosController{},
			),
		),

		beego.NSNamespace("/Roles",
			beego.NSInclude(
				&controllers.RolesController{},
			),
		),

		beego.NSNamespace("/Reservas",
			beego.NSInclude(
				&controllers.ReservasController{},
			),
		),

		beego.NSNamespace("/Tipo_Pagos",
			beego.NSInclude(
				&controllers.TipoPagosController{},
			),
		),

		beego.NSNamespace("/Roles_Usuario",
			beego.NSInclude(
				&controllers.RolesUsuarioController{},
			),
		),

		beego.NSNamespace("/Ranking",
			beego.NSInclude(
				&controllers.RankingController{},
			),
		),

		beego.NSNamespace("/Slots",
			beego.NSInclude(
				&controllers.SlotsController{},
			),
		),

		beego.NSNamespace("/Usuarios",
			beego.NSInclude(
				&controllers.UsuariosController{},
			),
		),

		beego.NSNamespace("/Vehiculos",
			beego.NSInclude(
				&controllers.VehiculosController{},
			),
		),

		beego.NSNamespace("/Administradores_Estacionamientos",
			beego.NSInclude(
				&controllers.AdministradoresEstacionamientosController{},
			),
		),

		beego.NSNamespace("/Administradores_Empleados",
			beego.NSInclude(
				&controllers.AdministradoresEmpleadosController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

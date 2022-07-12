package routes

import (
	"mini_project/controllers"

	"github.com/labstack/echo/v4"
	// mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// eJwt := e.Group("/jwt")
	// eJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))

	// route admin
	e.GET("/admin", controllers.GetAdminController)
	e.DELETE("/admin/:Id", controllers.DeleteAdminController)
	e.PUT("/admin/:Id", controllers.UpdateAdminController)
	e.POST("/admin", controllers.CreateAdminController)

	// route kost
	e.GET("/kost", controllers.GetKostController)
	e.POST("/kost", controllers.CreateKostController)
	e.DELETE("/kost/:Id", controllers.DeleteKostController)
	e.PUT("/kost/:Id", controllers.UpdateKostController)

	// route reservation
	e.GET("/pemesanan", controllers.GetPemesananController)
	e.POST("/pemesanan", controllers.CreatePemesananController)
	e.DELETE("/pemesanan/:Id", controllers.DeletePemesananController)
	e.PUT("/pemesanan/:Id", controllers.UpdatePemesananController)

	// route pemilik
	e.GET("/pemilik", controllers.GetPemilikController)
	e.POST("/pemilik", controllers.CreatePemilikController)
	e.DELETE("/pemilik/:Id", controllers.DeletePemilikController)
	e.PUT("/pemilik/:Id", controllers.UpdatePemilikController)

	// route pengunjung
	e.GET("/pengunjung", controllers.GetPengunjungController)
	e.GET("/pengunjung/:Id", controllers.GetPengunjungByIdController)
	e.DELETE("/pengunjung/:Id", controllers.DeletePengunjungController)
	e.PUT("/pengunjung/:Id", controllers.UpdatePengunjungController)
	e.POST("/pengunjung", controllers.CreatePengunjungController)

	return e
}

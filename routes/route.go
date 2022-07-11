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
	e.DELETE("/admin/:id", controllers.DeleteAdminController)
	e.PUT("/admin/:id", controllers.UpdateAdminController)
	e.POST("/admin", controllers.CreateAdminController)

	// route kost
	e.GET("/kost", controllers.GetKostController)
	e.POST("/kost", controllers.CreateKostController)
	e.DELETE("/kost/:id", controllers.DeleteKostController)
	e.PUT("/kost/:id", controllers.UpdateKostController)

	// route reservation
	e.GET("/pemesanan", controllers.GetPemesananController)
	e.POST("/pemesanan", controllers.CreatePemesananController)
	e.DELETE("/pemesanan/:id", controllers.DeletePemesananController)
	e.PUT("/pemesanan/:id", controllers.UpdatePemesananController)

	// route pemilik
	e.GET("/pemilik", controllers.GetPemilikController)
	e.POST("/pemilik", controllers.CreatePemilikController)
	e.DELETE("/pemilik/:id", controllers.DeletePemilikController)
	e.PUT("/pemilik/:id", controllers.UpdatePemilikController)

	// route pengunjung
	e.GET("/pengunjung", controllers.GetPengunjungController)
	e.GET("/pengunjung/:id", controllers.GetPengunjungByIdController)
	e.DELETE("/pengunjung/:id", controllers.DeletePengunjungController)
	e.PUT("/pengunjung/:id", controllers.UpdatePengunjungController)
	e.POST("/pengunjung", controllers.CreatePengunjungController)

	return e
}

package controllers

import (
	"mini_project/config"
	"mini_project/middleware"
	"mini_project/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginControllerAdmin(c echo.Context) error {
	admin := model.Admin{}
	c.Bind(&admin)

	err := config.DB.Where("email = ? AND password = ?", admin.Email, admin.Password).First(&admin).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	token, err := middleware.CreateToken(admin.Id, admin.Nama)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	adminResponse := model.AdminResponse{admin.Id, admin.Nama, admin.Email, token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success",
		"user":    adminResponse,
	})
}

func LoginControllerPemilik(c echo.Context) error {
	pemilik := model.Pemilik{}
	c.Bind(&pemilik)

	err := config.DB.Where("email = ? AND password = ?", pemilik.Email, pemilik.Password).First(&pemilik).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	token, err := middleware.CreateToken(pemilik.Id, pemilik.Nama)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	pemilikResponse := model.PemilikResponse{pemilik.Id, pemilik.Nama, pemilik.Email, token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success",
		"user":    pemilikResponse,
	})
}

func LoginControllerPengunjung(c echo.Context) error {
	pengunjung := model.Pengunjung{}
	c.Bind(&pengunjung)

	err := config.DB.Where("email = ? AND password = ?", pengunjung.Email, pengunjung.Password).First(&pengunjung).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	token, err := middleware.CreateToken(pengunjung.Id, pengunjung.Nama)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	pengunjungResponse := model.PengunjungResponse{pengunjung.Id, pengunjung.Nama, pengunjung.Email, token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success",
		"user":    pengunjungResponse,
	})
}

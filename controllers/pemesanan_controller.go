package controllers

import (
	"mini_project/config"
	"mini_project/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreatePemesananController(c echo.Context) error {
	pemesanan := model.Pemesanan{}
	c.Bind(&pemesanan)

	err := config.DB.Create(&pemesanan).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage":   "success create pemesanan",
		"pemesanan": pemesanan,
	})
}

func GetPemesananController(c echo.Context) error {
	var pemesanan []model.Pemesanan

	err := config.DB.Find(&pemesanan).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    pemesanan,
	})
}

func DeletePemesananController(c echo.Context) error {
	stringId := c.Param("ID_Pemesanan")
	err := config.DB.Delete(&model.Pemesanan{}, "ID_Pemesanan = ?", stringId).Debug().Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete pemesanan with ID_Pemesanan `" + stringId + "`",
	})
}

func UpdatePemesananController(c echo.Context) error {
	pemesanan := model.Pemesanan{}
	c.Bind(&pemesanan)
	stringId := c.Param("ID_Pemesanan")
	err := config.DB.Model(&pemesanan).Where("ID_Pemesanan = ?", stringId).Updates(pemesanan).Debug().Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "success",
		"pemesanan": pemesanan,
	})
}

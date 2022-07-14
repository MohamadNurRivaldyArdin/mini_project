package controllers

import (
	"mini_project/config"
	"mini_project/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreatePengunjungController(c echo.Context) error {
	pengunjung := model.Pengunjung{}
	c.Bind(&pengunjung)

	err := config.DB.Create(&pengunjung).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage":    "success create account",
		"pengunjung": pengunjung,
	})
}

func GetPengunjungByIdController(c echo.Context) error {
	var pengunjung model.Pengunjung

	stringId := c.Param("Id")
	err := config.DB.First(&pengunjung, "Id = ?", stringId).Debug().Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "success get pengunjung by Id",
		"pengunjung": pengunjung,
	})
}

func GetPengunjungController(c echo.Context) error {
	var pengunjung []model.Pengunjung

	err := config.DB.Find(&pengunjung).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    pengunjung,
	})
}

func DeletePengunjungController(c echo.Context) error {
	stringId := c.Param("ID")
	err := config.DB.Delete(&model.Pengunjung{}, "ID = ?", stringId).Debug().Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete account with IdPengunjung `" + stringId + "`",
	})
}

func UpdatePengunjungController(c echo.Context) error {
	pengunjung := model.Pengunjung{}
	c.Bind(&pengunjung)
	stringId := c.Param("Id")
	err := config.DB.Model(&pengunjung).Where("Id = ?", stringId).Updates(pengunjung).Debug().Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "success",
		"pengunjung": pengunjung,
	})
}

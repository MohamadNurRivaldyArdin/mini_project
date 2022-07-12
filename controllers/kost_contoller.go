package controllers

import (
	"mini_project/config"
	"mini_project/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateKostController(c echo.Context) error {
	kost := model.Kost{}
	c.Bind(&kost)

	err := config.DB.Save(&kost).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success create kost",
		"kost":    kost,
	})
}

func GetKostController(c echo.Context) error {
	var kost []model.Kost

	err := config.DB.Find(&kost).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    kost,
	})
}

func DeleteKostController(c echo.Context) error {
	stringId := c.Param("Id")
	err := config.DB.Delete(&model.Kost{}, "Id = ?", stringId).Debug().Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete kost with Id `" + stringId + "`",
	})
}

func UpdateKostController(c echo.Context) error {
	kost := model.Kost{}
	c.Bind(&kost)
	stringId := c.Param("Id")
	err := config.DB.Model(&kost).Where("Id = ?", stringId).Updates(kost).Debug().Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"kost":    kost,
	})
}

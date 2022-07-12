package controllers

import (
	"mini_project/config"
	"mini_project/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateAdminController(c echo.Context) error {
	admin := model.Admin{}
	c.Bind(&admin)

	err := config.DB.Create(&admin).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success create account",
		"admin":   admin,
	})
}

func GetAdminController(c echo.Context) error {
	var admin []model.Admin

	err := config.DB.Find(&admin).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    admin,
	})
}

func DeleteAdminController(c echo.Context) error {
	stringId := c.Param("Id")
	err := config.DB.Delete(&model.Admin{}, "Id = ?", stringId).Debug().Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete account with IdAdmin `" + stringId + "`",
	})
}

func UpdateAdminController(c echo.Context) error {
	admin := model.Admin{}
	c.Bind(&admin)
	stringId := c.Param("Id")
	err := config.DB.Model(&admin).Where("Id = ?", stringId).Updates(admin).Debug().Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"admin":   admin,
	})
}

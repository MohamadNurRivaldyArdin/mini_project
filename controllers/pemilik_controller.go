package controllers

import (
	"mini_project/config"
	"mini_project/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreatePemilikController(c echo.Context) error {
	pemilik := model.Pemilik{}
	c.Bind(&pemilik)

	err := config.DB.Save(&pemilik).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success create pemilik",
		"pemilik": pemilik,
	})
}

func GetPemilikController(c echo.Context) error {
	var pemilik []model.Pemilik

	err := config.DB.Find(&pemilik).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    pemilik,
	})
}

func DeletePemilikController(c echo.Context) error {
	stringId := c.Param("ID")
	err := config.DB.Delete(&model.Pemilik{}, "ID = ?", stringId).Debug().Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete account with ID `" + stringId + "`",
	})
}

func UpdatePemilikController(c echo.Context) error {
	pemilik := model.Pemilik{}
	c.Bind(&pemilik)
	stringId := c.Param("ID")
	err := config.DB.Model(&pemilik).Where("ID = ?", stringId).Updates(pemilik).Debug().Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"pemilik": pemilik,
	})
}

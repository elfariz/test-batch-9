package controllers

import (
	"net/http"
	"unjuk-keterampilan/configs"
	"unjuk-keterampilan/models"

	"github.com/labstack/echo/v4"
)

func GetSekolah(c echo.Context) error {
	var sekolah []models.Sekolah

	result := configs.DB.Find(&sekolah)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed get data from Sekolah",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Success get Data",
		Data:    sekolah,
	})
}

func GetDetailSekolah(c echo.Context) error {
	nama := c.Param("nama_sekolah")
	var sekolah []models.Sekolah

	result := configs.DB.Where("nama_sekolah LIKE ?", "%"+nama+"%").Find(&sekolah)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed get data from Sekolah",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Success get Data",
		Data:    sekolah,
	})

}

func GetDetailSekolahById(c echo.Context) error {
	id := c.Param("id")
	var sekolah []models.Sekolah

	result := configs.DB.Find(&sekolah, "id = ?", id)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed get data from Sekolah",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Success get Data",
		Data:    sekolah,
	})

}

func AddSekolah(c echo.Context) error {
	var SekolahReq models.Sekolah
	c.Bind(&SekolahReq)

	result := configs.DB.Create(&SekolahReq)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed to insert Sekolah",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, models.BaseResponse{
		Status:  true,
		Message: "Success Insert Data",
		Data:    SekolahReq,
	})

}

func EditSekolah(c echo.Context) error {
	id := c.Param("id")
	var SekolahReq models.Sekolah
	c.Bind(&SekolahReq)

	result := configs.DB.Where("id = ?", id).Updates(&SekolahReq)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed to insert Sekolah",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Success Edit Data",
		Data:    SekolahReq,
	})

}

func DeleteSekolah(c echo.Context) error {
	id := c.Param("id")
	var Sekolah models.Sekolah

	result := configs.DB.Delete(&Sekolah, id)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed to Delete Sekolah",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Success Delete Data",
		Data:    nil,
	})

}

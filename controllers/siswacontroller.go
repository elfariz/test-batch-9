package controllers

import (
	"net/http"
	"unjuk-keterampilan/configs"
	"unjuk-keterampilan/models"

	"github.com/labstack/echo/v4"
)

func GetSiswa(c echo.Context) error {
	var siswas []models.Siswa

	result := configs.DB.Preload("Sekolah").Find(&siswas)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed get data from Siswa",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Success get Data",
		Data:    siswas,
	})
}

func GetDetailSiswa(c echo.Context) error {
	nisn := c.Param("nisn")
	var siswas []models.Siswa

	result := configs.DB.Preload("Sekolah").Find(&siswas, "nisn = ?", nisn)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed get data from Siswa",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Success get Data",
		Data:    siswas,
	})
}

func AddSiswa(c echo.Context) error {
	var SiswaReq models.Siswa
	c.Bind(&SiswaReq)

	result := configs.DB.Preload("Sekolah").Create(&SiswaReq)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed to insert Siswa",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, models.BaseResponse{
		Status:  true,
		Message: "Success Insert Data",
		Data:    SiswaReq,
	})

}

func EditSiswa(c echo.Context) error {
	id := c.Param("id")
	var SiswaReq models.Siswa
	c.Bind(&SiswaReq)

	result := configs.DB.Preload("Sekolah").Where("id = ?", id).Updates(&SiswaReq)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed to Update Siswa",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, models.BaseResponse{
		Status:  true,
		Message: "Success Edit Data",
		Data:    SiswaReq,
	})
}

func DeleteSiswa(c echo.Context) error {
	id := c.Param("id")
	var Siswa models.Siswa

	result := configs.DB.Delete(&Siswa, id)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed to Delete Siswa",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, models.BaseResponse{
		Status:  true,
		Message: "Success Delete Data",
		Data:    nil,
	})
}

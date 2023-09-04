package routes

import (
	"os"
	"unjuk-keterampilan/controllers"

	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	e.POST("/register", controllers.RegisterController)
	e.POST("/login", controllers.LoginController)

	eAuth := e.Group("")
	eAuth.Use(echojwt.JWT([]byte(os.Getenv("PRIVATE_KEY_JWT"))))
	eLog := eAuth.Group("")

	eLog.GET("siswa", controllers.GetSiswa)           //Get ALL Data Siswa
	e.GET("siswa/:nisn", controllers.GetDetailSiswa)  //Get Data Siswa By NISN
	eLog.POST("siswa", controllers.AddSiswa)          //Add Data Siswa
	eLog.PUT("siswa/:id", controllers.EditSiswa)      //Update Data Siswa
	eLog.DELETE("siswa/:id", controllers.DeleteSiswa) //Deleta Data Siswa

	eLog.GET("/sekolah", controllers.GetSekolah)                         //Get ALL Data Sekolah
	e.GET("/sekolah_search/:nama_sekolah", controllers.GetDetailSekolah) //Get Data Sekolah By Nama Sekolah
	eLog.GET("/sekolah/:id", controllers.GetDetailSekolahById)           //Get Data Sekolah By Id
	eLog.POST("/sekolah", controllers.AddSekolah)                        //Add Data Sekolah
	eLog.PUT("/sekolah/:id", controllers.EditSekolah)                    //Update Data Sekolah
	eLog.DELETE("/sekolah/:id", controllers.DeleteSekolah)               //Delete Data Sekolah
}

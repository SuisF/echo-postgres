package main

import (
	"fmt"
	"net/http"
	"strings"

	"echo-postgres/database"
	//"echo-postgres/models"
	"github.com/labstack/echo/v4"
)

type GetModel struct {
	Id uint `json:"id"`
	Jadwal string `json:"jadwal"`
	Inspector string `json:"inspector"`
	Kondisi string `json:"kondisi"`
}

type GetModelArray struct {
	Id uint `json:"id"`
	Jadwal []string `json:"jadwal"`
	Inspector []string `json:"inspector"`
	Kondisi []string `json:"kondisi"`
}


func main() {
	e := echo.New()

	database.Connect()



	
	e.POST("/inspection", func(c echo.Context) error {
		
		var allOfStruct struct {
			Jadwal string `json:"jadwal"`
			Inspector string `json:"inspector"`
			Kondisi	string `json:"kondisi"`
		}
		errBind := c.Bind(&allOfStruct);
		if errBind != nil {
			fmt.Println(errBind.Error())
		}
		

		dataToPostgres := fmt.Sprintf("INSERT INTO inspection(jadwal, inspector, kondisi) VALUES ('%s', '%s', '%s')",
						allOfStruct.Jadwal, allOfStruct.Inspector, allOfStruct.Kondisi)
		err := database.DB.Exec(dataToPostgres).Error
		if err != nil {
			fmt.Println("Error when inserting")
			fmt.Println(err.Error())
		}

		return c.JSON(http.StatusOK, echo.Map{
			"message" : "Insert Data Success",
		})

	})
	


	e.GET("/", func (c echo.Context) error {
		getModelArray := []GetModel{}
		getModelArray2 := []GetModelArray{}

		database.DB.Raw("SELECT * FROM public.inspection").Scan(&getModelArray)
	
		for i := range getModelArray{
			getModelArray[i].Jadwal = strings.TrimPrefix(getModelArray[i].Jadwal, "{")
			getModelArray[i].Jadwal = strings.TrimSuffix(getModelArray[i].Jadwal, "}")

			getModelArray[i].Inspector = strings.TrimPrefix(getModelArray[i].Inspector, "[")
			getModelArray[i].Inspector = strings.TrimSuffix(getModelArray[i].Inspector, "]")

			getModelArray[i].Kondisi = strings.TrimPrefix(getModelArray[i].Kondisi, "[{")
			getModelArray[i].Kondisi = strings.TrimSuffix(getModelArray[i].Kondisi, "}]")

			arrJ := strings.Split(getModelArray[i].Jadwal, ",")
			arrI := strings.Split(getModelArray[i].Inspector, ",")
			arrK := strings.Split(getModelArray[i].Kondisi, ",")

			getModelArray2 = append(getModelArray2, GetModelArray{
				Id: getModelArray[i].Id,
				Jadwal: arrJ,
				Inspector: arrI,
				Kondisi: arrK,
			})
		}

		return c.JSON(200, map[string]interface{}{
			"data": getModelArray2,
		})

	})

	e.Logger.Fatal(e.Start(":2000"))

}
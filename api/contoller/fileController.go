package contoller

import (
	"bufio"
	"fmt"

	"github.com/Jose-Gomez-c/challenge/api/services"
	"github.com/gin-gonic/gin"
)

type FileController interface {
	FillDataBase() gin.HandlerFunc
}

type fileControllerLayer struct {
	fileService services.UploadFileService
}

func NewfileController(fileService services.UploadFileService) FileController {
	return &fileControllerLayer{fileService: fileService}
}

func (fileControllerLayer *fileControllerLayer) FillDataBase() gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.Bind("multipart/form-data") != nil {
			fmt.Println("Fallo no es multi")
			return
		}
		form, _, err := context.Request.FormFile("file")
		if err != nil {
			fmt.Println("error con el context multipart")
			return
		}
		defer form.Close()

		scanner := bufio.NewScanner(form)
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println("Error al leer la primera línea: ", err)
			}
			return
		}
		for scanner.Scan() {
			info := scanner.Text()
			fmt.Println(info)
			result := fileControllerLayer.fileService.FillDataBase(info, "https://api.mercadolibre.com")
			fmt.Println(result)
		}
		// str := string(content)
		// result := fileControllerLayer.fileService.FillDataBase(str, "https://api.mercadolibre.com")
		// fmt.Println(result)

	}

}

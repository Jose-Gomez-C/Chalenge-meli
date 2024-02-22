package contoller

import (
	"fmt"
	"io"

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
		form, err := context.MultipartForm()
		if err != nil {
			fmt.Println("error con el context multipart")
			return
		}
		files := form.File["file"]
		if files == nil {
			fmt.Println("Erro con la llave")
			return
		}
		for _, file := range files {
			data, err := file.Open()
			if err != nil {
				fmt.Println("error abrir el archivo")
				return
			}
			defer data.Close()
			content, err := io.ReadAll(data)
			if err != nil {
				fmt.Println(err)
			}

			str := string(content)
			result := fileControllerLayer.fileService.FillDataBase(str, "https://api.mercadolibre.com")
			fmt.Println(result)

		}

	}
}

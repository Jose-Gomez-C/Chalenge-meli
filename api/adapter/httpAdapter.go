package adapter

import (
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type HttpAdapter interface {
	GetWithQuery(url string, query []string) (result string, err error)
	Get(url string) (result string, err error)
}

type httpAdapterLayer struct {
	context *gin.Engine
}

func NewHttpAdapter(c *gin.Engine) HttpAdapter {
	return &httpAdapterLayer{context: c}
}

func (httpAdapter httpAdapterLayer) GetWithQuery(url string, query []string) (result string, err error) {
	urlCompleted := url + strings.Join(query, ",")
	println(urlCompleted)
	response, err := http.Get(urlCompleted)
	if err != nil {
		println("Error en el request", err)
		return "", err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	return string(body), err
}

func (httpAdapter httpAdapterLayer) Get(url string) (result string, err error) {
	println(url)
	response, err := http.Get(url)
	if err != nil {
		println("Error en el request", err)
		return "", err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	return string(body), err
}

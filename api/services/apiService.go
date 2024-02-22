package services

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Jose-Gomez-c/challenge/api/adapter"
	"github.com/Jose-Gomez-c/challenge/api/model"
)

type ApiService interface {
	getItemInfoByProductsId(productsId []model.ProductId, url string) []model.ResponseItem
	getNameByCategoryId(url string, body model.BodyItem) model.ResponseCategory
	getDescriptionByCurrencyID(url string, body model.BodyItem) model.ResponseCurrency
	getNicknameBySellerID(url string, body model.BodyItem) model.ResponseUser
}

type apiServiceLayer struct {
	httpAdapter   adapter.HttpAdapter
	resdisAdapter adapter.RedisAdapter
}

func NewApiservice(httpAdapter adapter.HttpAdapter, redis adapter.RedisAdapter) ApiService {
	return &apiServiceLayer{httpAdapter: httpAdapter, resdisAdapter: redis}
}

func (api apiServiceLayer) getItemInfoByProductsId(productsId []model.ProductId, url string) []model.ResponseItem {
	var response []model.ResponseItem
	var finalUrl = url + "/items?ids="
	id := makeIdsForApi(productsId)
	body, err := api.httpAdapter.GetWithQuery(finalUrl, id)
	if err != nil {
		fmt.Println("no hay body")
	}
	if err := json.Unmarshal([]byte(body), &response); err != nil {
		fmt.Println("no hay body", err)
	}
	return response
}

func (api apiServiceLayer) getNameByCategoryId(url string, body model.BodyItem) model.ResponseCategory {
	value, err := api.resdisAdapter.GetCache(body.CategoryID)
	var response model.ResponseCategory
	if err != nil {
		info, err := api.httpAdapter.Get(url + "/categories/" + body.CategoryID)
		if err != nil {
			fmt.Println("no hay body")
		}
		if err := json.Unmarshal([]byte(info), &response); err != nil {
			fmt.Println("no hay body", err)
		}
		errCache := api.resdisAdapter.SendCache(body.CategoryID, response.Name)
		if errCache != nil {
			fmt.Println("error al enviar cache", errCache)
		}
	} else {
		response = model.ResponseCategory{Id: body.CategoryID, Name: value}
	}

	return response
}

func (api apiServiceLayer) getDescriptionByCurrencyID(url string, body model.BodyItem) model.ResponseCurrency {
	value, err := api.resdisAdapter.GetCache(body.CurrencyID)
	var response model.ResponseCurrency
	if err != nil {

		info, err := api.httpAdapter.Get(url + "/currencies/" + body.CurrencyID)
		if err != nil {
			fmt.Println("no hay body")
		}
		if err := json.Unmarshal([]byte(info), &response); err != nil {
			fmt.Println("no hay body", err)
		}
		errCache := api.resdisAdapter.SendCache(body.CurrencyID, response.Description)
		if errCache != nil {
			fmt.Println("error al enviar cache", errCache)
		}
	} else {
		response = model.ResponseCurrency{Id: body.CurrencyID, Description: value}
	}
	return response
}

func (api apiServiceLayer) getNicknameBySellerID(url string, body model.BodyItem) model.ResponseUser {
	value, err := api.resdisAdapter.GetCache(strconv.Itoa(body.SellerID))
	var response model.ResponseUser
	if err != nil {
		info, err := api.httpAdapter.Get(url + "/users/" + strconv.Itoa(body.SellerID))
		if err != nil {
			fmt.Println("no hay body")
		}
		if err := json.Unmarshal([]byte(info), &response); err != nil {
			fmt.Println("no hay body", err)
		}
		errCache := api.resdisAdapter.SendCache(strconv.Itoa(body.SellerID), response.Nickname)
		if errCache != nil {
			fmt.Println("error al enviar cache", errCache)
		}
	} else {
		response = model.ResponseUser{Id: body.SellerID, Nickname: value}
	}
	return response
}

func makeIdsForApi(productsId []model.ProductId) []string {
	var result []string
	for _, product := range productsId {
		id := product.Site + strconv.Itoa(product.Id)
		result = append(result, id)
	}
	return result
}

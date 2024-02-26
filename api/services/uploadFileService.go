package services

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Jose-Gomez-c/challenge/api/model"
	"github.com/Jose-Gomez-c/challenge/api/repositories"
)

type UploadFileService interface {
	FillDataBase(data string, url string) []model.ResponseItem
}

type uploadFileServiceLayer struct {
	apiService     ApiService
	itemRepository repositories.ItemRepository
}

func NewUploadServices(api ApiService, repository repositories.ItemRepository) UploadFileService {
	return &uploadFileServiceLayer{
		apiService:     api,
		itemRepository: repository,
	}
}

func (layer *uploadFileServiceLayer) getProductId(data string) []model.ProductId {
	var producIds []model.ProductId
	values := strings.Split(data, ",")
	id, err := strconv.Atoi(strings.ReplaceAll(values[1], "\r", ""))
	if err != nil {
		fmt.Println("Error al convertir entero: ", err)
	}
	info := mewProducId(id, values[0])
	producIds = append(producIds, info)
	return producIds
}

func (layer *uploadFileServiceLayer) FillDataBase(data string, url string) []model.ResponseItem {
	var haveInfo []model.BodyItem
	var noInfo []string
	productIds := layer.getProductId(data)
	body := layer.apiService.getItemInfoByProductsId(productIds, url)
	for _, item := range body {
		if item.Code == 200 {
			haveInfo = append(haveInfo, item.Body)
		} else {
			noInfo = append(noInfo, item.Body.Id)
		}
	}
	itemModel, err := layer.getAllInfo(haveInfo, url)
	if err != nil {
		fmt.Println("error al obtener info", err)
	}

	layer.sendItemsToDb(itemModel, noInfo)
	return body
}

func (layer *uploadFileServiceLayer) getAllInfo(items []model.BodyItem, url string) ([]model.Items, error) {
	var itemModel []model.Items
	for _, info := range items {
		category := layer.apiService.getNameByCategoryId(url, info)
		currency := layer.apiService.getDescriptionByCurrencyID(url, info)
		seller := layer.apiService.getNicknameBySellerID(url, info)
		producId, err := getProductIdFromString(info.Id)
		if err != nil {
			fmt.Println("error al cconvertir el id")
			return nil, err
		}
		item := newItemModel(category, currency, seller, producId, info)
		itemModel = append(itemModel, item)
	}
	return itemModel, nil
}

func (layer *uploadFileServiceLayer) sendItemsToDb(items []model.Items, noInfo []string) {
	var noInfoItems []model.Items
	layer.itemRepository.SaveInBatch(items)
	for _, bi := range noInfo {
		productId, _ := getProductIdFromString(bi)
		item := newItemModelWithOutInfo(productId)
		noInfoItems = append(noInfoItems, item)
	}
	layer.itemRepository.SaveInBatch(noInfoItems)
}

func mewProducId(id int, site string) model.ProductId {
	return model.ProductId{
		Id:   id,
		Site: site,
	}
}

func newItemModel(category model.ResponseCategory, currency model.ResponseCurrency, seller model.ResponseUser, productId model.ProductId, allInfo model.BodyItem) model.Items {
	return model.Items{
		SiteId:       productId.Site,
		Id:           productId.Id,
		Price:        allInfo.Price,
		NameCategory: category.Name,
		Description:  currency.Description,
		Nickname:     seller.Nickname,
	}
}
func newItemModelWithOutInfo(productId model.ProductId) model.Items {
	return model.Items{Id: productId.Id, SiteId: productId.Site, Price: 0, NameCategory: "Not found", Description: "Not found", Nickname: "Not found"}
}

func getProductIdFromString(id string) (model.ProductId, error) {
	re := regexp.MustCompile(`(\d+)`)
	matches := re.FindAllStringSubmatch(id, -1)
	if len(matches) > 0 && len(matches[0]) > 1 {
		numberStr := matches[0][1]
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			return mewProducId(0, ""), err
		}
		return mewProducId(number, id[:len(id)-len(numberStr)]), nil
	}

	// Si no se encontró ningún número, devolver un error
	return mewProducId(0, ""), fmt.Errorf("code not found")
}

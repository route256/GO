package products

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"gitlab.ozon.dev/go/classroom-20/experts/product-service/internal/domain/model"
)

var productsMap map[int64]*model.Product

func IsProductsReady() bool {
	return len(productsMap) > 0
}

func init() {

	defaultDataFileName := "init/data.json"

	if _, err := os.Stat(defaultDataFileName); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("init/data.json is absent ... unit-test? %v", err)
		return
	}

	InitWithDataFileName(defaultDataFileName)

}

func InitWithDataFileName(filename string) {

	var initMap map[string]model.Product
	fdata, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(fdata, &initMap)

	productsMap = make(map[int64]*model.Product, len(initMap))
	for k, v := range initMap {
		i, _ := strconv.ParseInt(k, 10, 64)
		v.SKU = i
		productsMap[i] = &v
	}
}

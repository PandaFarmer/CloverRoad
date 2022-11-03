package model3ds

import (
	"fmt"
	// "strconv"
	"strings"

	"github.com/PandaFarmer/CloverRoad/backend-fiber/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetModel3Ds(c *fiber.Ctx) error {
	var model3ds []models.Model3D

	if result := h.DB.Find(&model3ds); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&model3ds)
}

type SearchModel3DRequestBody struct {
	Keywords   string `form:"keywords"`
	MaxResults int    `form:"max_results"`
	// LastUpdated time.Time `form:"last_updated"`
	MinPrice float64 `form:"min_price"`
	MaxPrice float64 `form:"max_price"`
}

func (h handler) GetModel3DsSearch(c *fiber.Ctx) error {
	fmt.Println("processing model3d search request")
	body := SearchModel3DRequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	keywords := body.Keywords
	maxResults := body.MaxResults
	// lastUpdated := body.LastUpdated
	minPrice := body.MinPrice
	maxPrice := body.MaxPrice

	// fmt.Println("JAJAJAJAJAJAJ")
	// keywords := c.Params("keywords")
	// maxResults, _ := strconv.ParseInt(c.Params("max_results"), 10, 0)
	// fmt.Println(c.Params("lastUpdated"))
	// lastUpdated, err := time.Parse(time.UnixDate, c.Params("lastUpdated"))
	// if err != nil {
	// 	return fiber.NewError(fiber.StatusNotFound, err.Error())
	// }
	// fmt.Println("minPrice: " + c.Params("minPrice"))
	// fmt.Println("maxPrice: " + c.Params("maxPrice"))
	// minPrice, _ := strconv.ParseFloat(c.Params("minPrice"), 32)
	// maxPrice, _ := strconv.ParseFloat(c.Params("maxPrice"), 32)
	// fmt.Sprintln("type of minPrice: %T", minPrice)
	//Author, Title, maybe Description + price range, on sale?

	var model3ds []models.Model3D
	parsedKeywords := strings.Fields(keywords)
	//.Find changes result type? :(
	queryDB := h.DB.Where("Price >= ?", minPrice).Where("Price <= ?", maxPrice)

	for _, parsedKeyword := range parsedKeywords {
		queryDB = queryDB.Where("? in Title", parsedKeyword).Or("? in Author", parsedKeyword).Or("? in Description", parsedKeyword)
	}
	// if result := queryDB.Limit(int(maxResults)).Where("updated_at <= ?", lastUpdated).Find(&model3ds); result.Error != nil {
	if result := queryDB.Limit(int(maxResults)).Find(&model3ds); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&model3ds)
}

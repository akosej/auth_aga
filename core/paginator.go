package core

import (
	"fmt"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Pagination struct {
	Limit      int         `json:"limit,omitempty,query:limit"`
	Page       int         `json:"page,omitempty;query:page"`
	Sort       string      `json:"sort,omitempty;query:sort"`
	TotalItems int64       `json:"totalItems"`
	TotalPages int         `json:"totalPages"`
	HasPrev    bool        `json:"hasPrev"`
	HasNext    bool        `json:"hasNext"`
	Items      interface{} `json:"items"`
	PrevURL    string      `json:"prevURL,omitempty"`
	NextURL    string      `json:"nextURL,omitempty"`
}

func Paginate(query *gorm.DB, itemsPtr interface{}, ctx *fiber.Ctx) (Pagination, error) {
	var pagination Pagination

	limit, err := strconv.Atoi(ctx.Query("limit", "47"))
	if err != nil {
		return pagination, err
	}
	page, err := strconv.Atoi(ctx.Query("page", "1"))
	if err != nil {
		return pagination, err
	}
	sort := ctx.Query("sort", "")

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return pagination, err
	}

	totalPages := int(math.Ceil(float64(count) / float64(limit)))

	offset := (page - 1) * limit
	if offset < 0 {
		offset = 0
	}
	if err := query.Offset(offset).Limit(limit).Order(sort).Find(itemsPtr).Error; err != nil {
		return pagination, err
	}

	baseUrl := ctx.Path()
	if page > 1 {
		prevPage := page - 1
		prevUrl := fmt.Sprintf("%s?page=%d&limit=%d&sort=%s", baseUrl, prevPage, limit, sort)
		pagination.PrevURL = prevUrl
	}

	if page < totalPages {
		nextPage := page + 1
		nextUrl := fmt.Sprintf("%s?page=%d&limit=%d&sort=%s", baseUrl, nextPage, limit, sort)
		pagination.NextURL = nextUrl
	}

	pagination.Limit = limit
	pagination.Page = page
	pagination.Sort = sort
	pagination.TotalItems = count
	pagination.TotalPages = totalPages
	pagination.HasPrev = page > 1
	pagination.HasNext = page < totalPages
	pagination.Items = itemsPtr

	return pagination, nil
}

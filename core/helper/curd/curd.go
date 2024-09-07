package curd

import (
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"lazyfury.github.com/yoomall-server/core/driver"
	"lazyfury.github.com/yoomall-server/core/helper/response"
)

type CRUD struct {
	DB    *driver.DB
	Model interface{}
}

func (c *CRUD) GetDB() *driver.DB {
	return c.DB
}

func (c *CRUD) GetTableName() string {
	defer recover()

	model := c.Model
	if model == nil {
		return ""
	}
	fn, ok := model.(interface {
		TableName() string
	})

	if !ok {
		return ""
	}
	return fn.TableName()
}

func (c *CRUD) GetList(list any, page int, limit int, params map[string]string) (tx *gorm.DB) {
	return c.DB.Model(c.Model).Where(params).Limit(limit).Offset((page - 1) * limit)
}

// get list handler
func (c *CRUD) GetListHandler(list any) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
		var params map[string]string = make(map[string]string)
		ctx.ShouldBindQuery(&params)
		delete(params, "page")
		delete(params, "limit")
		log.Info("get list", "page", page, "limit", limit, "params", params)

		// query
		query := c.DB.Model(c.Model).Where(params)
		var count int64 = 0
		query.Count(&count)

		err := query.Limit(limit).Offset((page - 1) * limit).Find(list).Error

		if err != nil {
			response.Error(response.ErrInternalError, err.Error()).Done(ctx)
			return
		}

		response.Success(map[string]any{
			"data":  list,
			"total": count,
			"page":  page,
			"limit": limit,
			"pages": count / int64(limit),
		}).Done(ctx)
	}
}

func (c *CRUD) GetById(id int64) (interface{}, error) {
	var model interface{}
	err := c.DB.Model(c.Model).Where("id = ?", id).First(&model).Error
	return model, err
}

func (c *CRUD) Create(model interface{}) error {
	return c.DB.Model(c.Model).Create(model).Error
}

func (c *CRUD) Update(model interface{}) error {
	return c.DB.Model(c.Model).Save(model).Error
}

func (c *CRUD) Delete(model interface{}) error {
	return c.DB.Model(c.Model).Delete(model).Error
}

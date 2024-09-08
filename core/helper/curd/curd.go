package curd

import (
	"regexp"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"lazyfury.github.com/yoomall-server/core/driver"
	"lazyfury.github.com/yoomall-server/core/helper/response"
	"lazyfury.github.com/yoomall-server/core/helper/utils"
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

func (c *CRUD) Where(params map[string]interface{}) *gorm.DB {
	tx := c.DB.Model(c.Model)
	for k, v := range params {

		// log.Info("where", "key", k, "value", v)
		if regexp.MustCompile(`^(\S+)__(in|not_in|like|eq|gt|lt|is_null|is_not_null)$`).MatchString(k) {
			delete(params, k)
			find := regexp.MustCompile(`^(\S+)__(in|not in|like|eq|gt|lt)$`).FindStringSubmatch(k)
			key := find[1]
			action := find[2]
			log.Info("find pattern", "key", key, "action", action)
			switch action {
			case "in":
				tx = tx.Where(key+" IN (?)", utils.TryInterfaceToStringToArray(v))
			case "not_in":
				tx = tx.Where(key+" NOT IN (?)", utils.TryInterfaceToStringToArray(v))
			case "like":
				tx = tx.Where(key+" LIKE ?", "%"+v.(string)+"%")
			case "eq":
				tx = tx.Where(key+" = ?", v)
			case "gt":
				tx = tx.Where(key+" > ?", v)
			case "lt":
				tx = tx.Where(key+" < ?", v)
			case "is_null":
				tx = tx.Where(key + " IS NULL")
			case "is_not_null":
				tx = tx.Where(key + " IS NOT NULL")
			}

			continue
		}

		// sort
		if regexp.MustCompile(`^(\S+)__(desc|asc)$`).MatchString(k) {
			delete(params, k)
			find := regexp.MustCompile(`^(\S+)__(desc|asc)$`).FindStringSubmatch(k)
			key := find[1]
			action := find[2]
			// log.Info("find pattern", "key", key, "action", action)
			switch action {
			case "desc":
				tx = tx.Order(key + " DESC")
			case "asc":
				tx = tx.Order(key + " ASC")
			}
		}

		if v == "" {
			delete(params, k)
		}

	}
	tx = tx.Where(params)
	return tx
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
		query := c.Where(utils.StringMapToInterfaceMap(params))
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

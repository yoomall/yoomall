package curd

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

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

func (c *CRUD) GetList(ctx *gin.Context) map[string]any {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	var params map[string]string = make(map[string]string)
	ctx.ShouldBindQuery(&params)
	delete(params, "page")
	delete(params, "limit")
	return map[string]any{
		"query": c.Where(utils.StringMapToInterfaceMap(params)),
		"page":  page,
		"limit": limit,
	}
}

var searchAct = []string{"in", "not_in", "like", "eq", "gt", "lt", "is_null", "is_not_null", "asc", "desc"}

func (c *CRUD) Where(params map[string]interface{}) *gorm.DB {
	tx := c.DB.Model(c.Model)
	for k, v := range params {

		// 外键查询
		fk_reg := fmt.Sprintf(`^(\S+)__(\S+)__fk__(%s)$`, strings.Join(searchAct, "|"))
		if find := regexp.MustCompile(fk_reg).FindStringSubmatch(k); len(find) > 0 {
			delete(params, k)
			preload := find[1]
			key := find[2]
			opration := find[3]

			if preload != "" {
				tx = tx.Joins(preload)
				tx = c.superWhere(opration, tx, fmt.Sprintf("%s.%s", preload, key), v)
			}
			continue
		}

		where_reg := fmt.Sprintf(`^(\S+)__(%s)$`, strings.Join(searchAct, "|"))
		if find := regexp.MustCompile(where_reg).FindStringSubmatch(k); len(find) > 0 {
			delete(params, k)
			key := find[1]
			action := find[2]
			// log.Info("find pattern", "key", key, "action", action)
			tx = c.superWhere(action, tx, key, v)
			continue
		}

		if v == "" {
			delete(params, k)
		}

	}
	tx = tx.Where(params)
	return tx
}

func (*CRUD) superWhere(action string, tx *gorm.DB, key string, v interface{}) *gorm.DB {
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
	// desc
	case "desc":
		tx = tx.Order(key + " DESC")
	case "asc":
		tx = tx.Order(key + " ASC")
	}

	return tx
}

// get list handler
func (c *CRUD) GetListHandler(list any, extraWhere func(tx *gorm.DB) *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		getListCurd := c.GetList(ctx)
		var query *gorm.DB = getListCurd["query"].(*gorm.DB)
		page := getListCurd["page"].(int)
		limit := getListCurd["limit"].(int)

		// query
		if extraWhere != nil {
			query = extraWhere(query)
		}
		var count int64 = 0
		query.Count(&count)

		err := query.Offset((page - 1) * limit).Limit(limit).Find(list).Error

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

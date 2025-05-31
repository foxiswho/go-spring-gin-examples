package repositoryDemo

import (
	"github.com/gin-gonic/gin"
	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/spring-core/util/syslog"
	"github/foxiswho/go-spring-gin-examples/infrastructure/entityDemo"
	"gorm.io/gorm"
	"reflect"
)

func init() {
	gs.Provide(&UserRepository{Entity: &entityDemo.UserEntity{}}).Init(func(s *UserRepository) {
		syslog.Debugf("%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type UserRepository struct {
	db     *gorm.DB `autowire:"?"`
	ctx    *gin.Context
	Entity *entityDemo.UserEntity
}

func (b *UserRepository) Db() *gorm.DB {
	return b.db
}

// 保存
func (b *UserRepository) Save(info *entityDemo.UserEntity) error {
	if e := b.db.Save(info); e != nil {
		return e.Error
	}
	return nil
}

// 保存
func (b *UserRepository) SaveAll(ts []*entityDemo.UserEntity) error {
	if ts != nil && len(ts) > 0 {
		for _, info := range ts {
			if e := b.db.Save(info); e != nil {
				return e.Error
			}
		}
	}
	return nil
}

// Update 更新 更新属性，只会更新非零值的字段
func (b *UserRepository) Update(info entityDemo.UserEntity, id int64) error {
	result := b.db.Model(b.Entity).Where("id=?", id).Updates(&info)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FindByIdString 根据主键查询
//
//	@Description:
//	@receiver b
//	@param id
//	@return info
//	@return result 是否查询到值
//	@return err
func (b *UserRepository) FindByIdString(id string) (info *entityDemo.UserEntity, result bool) {
	tx := b.db.Where("id=?", id).First(&info)
	if tx.Error != nil {
		return nil, false
	}
	if tx.RowsAffected == 0 {
		return nil, false
	}
	return info, true
}

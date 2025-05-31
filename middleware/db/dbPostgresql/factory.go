package dbPostgresql

import (
	"errors"
	"fmt"
	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/spring-core/util/syslog"
	"github/foxiswho/go-spring-gin-examples/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 数据库工厂 Postgresql
type Factory struct {
	database config.Database `value:"${database}"`
}

func (factory *Factory) CreateDB() (*gorm.DB, error) {
	syslog.Debugf("数据库连接地址： enabled:%+v,URL:=>%s", factory.database.Enabled, factory.database.URL)
	fmt.Printf("数据库连接地址： enabled:%+v,URL:=>%s", factory.database.Enabled, factory.database.URL)
	if !factory.database.Enabled {
		syslog.Debugf("未启用数据库： %s", factory.database.Enabled)
		return nil, nil
	}
	db, err := gorm.Open(postgres.Open(factory.database.URL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		syslog.Errorf("open gorm postgresql %s error: %v", factory.database.URL, err)
		panic(errors.New(err.Error()))
		return nil, err
	}
	dbName := db.Migrator().CurrentDatabase()
	// check if db exists
	stmt := fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s';", dbName)
	rs := db.Raw(stmt)
	if rs.Error != nil {
		return nil, rs.Error
	}
	// if not create it
	var rec = make(map[string]interface{})
	if rs.Find(rec); len(rec) == 0 {
		stmt := fmt.Sprintf("CREATE DATABASE %s;", dbName)
		if rs := db.Exec(stmt); rs.Error != nil {
			return nil, rs.Error
		}

		// close db connection
		sql, err := db.DB()
		defer func() {
			_ = sql.Close()
		}()
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}

func init() {
	gs.Object(&Factory{})
	gs.Provide((*Factory).CreateDB).
		//指定名称
		Name("GormDb").
		//当指定类型/名称的 Bean 不存在时激活
		Condition(
			gs.OnProperty("database.enabled").HavingValue("true").MatchIfMissing(),
			// GormDB 不存在
			gs.OnMissingBean[*gorm.DB]("GormDB"))
}

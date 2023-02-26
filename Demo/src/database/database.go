package database

import (
	"Demo/src/config"
	"fmt"
	"time"

	"github.com/allegro/bigcache"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	GlobalCache *bigcache.BigCache
)

func Init() {
	// Connect to DB
	var err error
	dsn := config.Appconfig.GetString("database.username") + ":" + config.Appconfig.GetString("database.password") + "@tcp(" + config.Appconfig.GetString("database.host") + ":" + config.Appconfig.GetString("database.port") + ")/" + config.Appconfig.GetString("database.name") + "?parseTime=true"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect to DB: %w", err))
	}
	GlobalCache, err = bigcache.NewBigCache(bigcache.DefaultConfig(30 * time.Minute))
	if err != nil {
		panic(fmt.Errorf("failed to initialize cahce: %w", err))
	}
}

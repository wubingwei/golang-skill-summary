package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"github.com/wubingwei/veigar/config"
	"github.com/wubingwei/veigar/dbmodel"
)

var (
	store              *gorm.DB
	AdministratorStore dbmodel.AdministratorStore
)

func init() {
	store = NewStore(&config.Mysql)
	AdministratorStore = &administrator{}
}

func NewStore(DB *config.MysqlDB) *gorm.DB {
	connParams := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Asia%%2FShanghai", DB.User, DB.Password, DB.Host, DB.Port, DB.DataBase)
	logrus.Debugf("Connecting to DB: %s", connParams)
	db, err := gorm.Open("mysql", connParams)
	if err != nil {
		logrus.Errorf("Error during connecting to DB: %v", err)
	}
	return db
}

func GetDB() *gorm.DB {
	db := store.New()
	return db
}

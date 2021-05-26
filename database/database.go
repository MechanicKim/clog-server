package database

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type CLog struct {
	ID    int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Title string `gorm:"not null" form:"title" json:"title"`
	Desc  string `gorm:"not null" form:"desc" json:"desc"`
	Days  int    `gorm:"not null" form:"days" json:"days"`
	Start string `gorm:"not null" form:"start" json:"start"`
	End   string `gorm:"not null" form:"end" json:"end"`
	Total int    `gorm:"not null" form:"total" json:"total"`
}

type CLogDay struct {
	ID     int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	CLogID int    `gorm:"not null" form:"cLogId" json:"cLogId"`
	Date   string `gorm:"not null" form:"date" json:"date"`
	Count  int    `gorm:"default:0" form:"count" json:"count"`
}

var myDB *gorm.DB

func Open() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // Disable color
		},
	)

	db, err := gorm.Open(sqlite.Open("clog.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&CLog{})
	db.AutoMigrate(&CLogDay{})

	fmt.Println("type:", reflect.ValueOf(db).Type())
	myDB = db
}

func AddCLog(cLog CLog) CLog {
	myDB.Create(&cLog)
	return cLog
}

func GetCLogs() []CLog {
	var cLogs []CLog
	myDB.Find(&cLogs)
	return cLogs
}

func GetCLog(id int) CLog {
	var cLog CLog
	myDB.First(&cLog, id)
	return cLog
}

func DeleteCLog(id int) {
	cLog := GetCLog(id)
	myDB.Delete(&cLog, id)
	days := GetCLogDays(id)
	myDB.Delete(&days)
}

func UpdateCLog(cLog CLog) CLog {
	myDB.Save(&cLog)
	return cLog
}

func AddCLogDays(days []CLogDay) []CLogDay {
	myDB.Create(&days)
	return days
}

func GetCLogDays(cLogId int) []CLogDay {
	var days []CLogDay
	myDB.Where(&CLogDay{CLogID: cLogId}).Find(&days)
	return days
}

func UpdateCLogDay(day CLogDay) CLogDay {
	myDB.Save(&day)
	return day
}

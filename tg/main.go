package tg

import (
	"github.com/jinzhu/gorm"
	"github.com/petuhovskiy/telegram"

	"github.com/lodthe/bdaytracker-go/conf"
)

// General keeps the general fields for a session
type General struct {
	Bot    *telegram.Bot
	DB     *gorm.DB
	Config conf.Config
}
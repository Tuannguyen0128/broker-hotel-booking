package repositories

import (
	"broker-hotel-booking/configs"
	"broker-hotel-booking/internal/repositories/log"
	"broker-hotel-booking/internal/util"
)

type (
	Repositories struct {
		LogRepo *log.LogRepository
	}
)

func New() *Repositories {
	conf := configs.Load("./configs/config.yaml")
	logRepo := log.New(util.InitConnection(conf.MongoDB))

	return &Repositories{
		LogRepo: logRepo,
	}
}

package cron_auth_server

import (
	"fmt"
	"time"

	interfaces "github.com/bibin-zoz/social-match-room-svc/pkg/repository/interface"
	cron "github.com/robfig/cron/v3"
)

type CronJob struct {
	roomRepository interfaces.RoomRepository
	Location       *time.Location
}

func NewCronJob(roomRepository interfaces.RoomRepository) *CronJob {
	locationInd, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		fmt.Println("error at time place", err)
	}
	return &CronJob{roomRepository: roomRepository,
		Location: locationInd}
}

func (c *CronJob) StartCronInAuthService() {
	newCron := cron.New()

	newCron.AddFunc("*@daily", func() {
		c.roomRepository.CloseExpiredRooms()
	})

	newCron.Start()
}

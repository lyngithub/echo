package job

import (
	"github.com/robfig/cron"
)

func InitConsole() {
	var c = cron.New()
	var err error

	// 每分钟执行一次
	err = c.AddFunc("0 */1 * * * ?", func() {
		//_ = services.JobTasks.GetCoinPrice()
	})

	if err != nil {
		return
	}

	c.Start()
	select {}
}

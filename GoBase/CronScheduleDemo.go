package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/robfig/cron"
)

func TestSchedule(t *testing.T) {

	jobCron := cron.New()

	//秒 分 时 日 月 星期
	//0 0 8 * * *

	/*
		/ 字符可用于指定值的增量。
		? 问号可以用于月日和星期字段。用于指定“无特定值”。当您需要在这两个字段中的一个字段中指定某些内容，而不是另一个字段时，这很有用。
	*/

	//jobCronMain.AddFunc("0 0 8 * * *", EveryDayHello)
	//jobCronMain.AddFunc("0 0 9 * * *", CheckWsClientOnline)
	//jobCronMain.AddFunc("0 0 15 * * *", CheckWsClientOnline)
	//jobCronMain.AddFunc("0 0/5 * * * ?", SayHello)
	//jobCronMain.AddFunc("10 0/5 * * * ?", SayHello)
	//jobCronMain.AddFunc("0 0 0/8 * * *", EveryDayHello)
	//jobCronMain.AddFunc("0 30 10-13 ? * WED,FRI", SayHello)
	//jobCronMain.AddFunc("0 0/30 8-9 5,20 * ?", SayHello)
	//jobCronMain.AddFunc("0 0/30 8-9 ? * FRI#3", SayHello)

	jobCron.AddFunc("0/10 * * * * *", func() { fmt.Println("workFunc come here.", time.Now()) })

	jobCron.AddFunc("@every 2s", func() { fmt.Println("workFunc every 2s.", time.Now()) })

	//带自定义参数的定时任务
	paramsJob := &ParamsJob{P1: "some string", P2: true}
	jobCron.AddJob("0/10 * * * * *", paramsJob)

	jobCron.Schedule(cron.Every(time.Second), cron.FuncJob(func() { fmt.Println("workFunc Every Second.", time.Now()) }))

	jobCron.Schedule(cron.Every(time.Second*5), cron.FuncJob(func() { fmt.Println("workFunc Every 5 Second.", time.Now()) }))

	jobCron.Schedule(cron.Every(time.Minute), cron.FuncJob(func() { fmt.Println("workFunc Every Minute.", time.Now()) }))

	jobCron.Schedule(cron.Every(time.Hour), cron.FuncJob(func() { fmt.Println("workFunc Every Hour.", time.Now()) }))

	jobCron.Start()

}

type ParamsJob struct {
	P1 string
	P2 bool
}

func (_self *ParamsJob) Run() {
	//执行逻辑
}

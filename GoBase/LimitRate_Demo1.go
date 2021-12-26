package main

import (
	"context"
	"time"
)
import "golang.org/x/time/rate"

func LimitRate_Demo1_main(){
	//限流器

	//NewLimiter(r Limit, b int)
	//r 代表每秒可以向Token桶中产生多少token
	//b 代表Token桶的容量大小

	//每100ms往桶中放一个Token
	limiter := rate.NewLimiter(rate.Every(time.Millisecond * 100), 1)

	//三种消费Token的方式：Wait、Allow、Reserve

	//Wait方式
	//可以设置context的Deadline或者Timeout，来决定此次Wait的最长时间
	limiter.Wait(context.Background())	//Wait实际上就是WaitN(ctx,1)
	//limiter.WaitN(context.Background(),2)

	//Allow方式
	//截止到某一时刻，目前桶中数目是否至少为n个，满足则返回true，同时从桶中消费n个token
	//反之返回不消费Token，false
	limiter.AllowN(time.Now(), 1)

	//Reserve方式
	//当调用完成后，无论Token是否充足，都会返回一个 *Reservation
	//可以调用该对象的Delay()方法，该方法返回了需要等待的时间。如果等待时间为0，则说明不用等待。
	//必须等到等待时间之后，才能进行接下来的工作，如果不想等待，可以调用Cancel()方法，该方法会将Token归还
	limiter.ReserveN(time.Now(), 1)


	//在运行时动态调整限流器参数
	limiter.SetLimit(rate.Every(time.Millisecond * 10))
	limiter.SetBurst(20)


}
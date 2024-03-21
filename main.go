package main

import (
	"context"
	"log"
	"task/flood/config"
	"task/flood/control"

	"time"
)

const (
    fileName string = "config.yaml"
)

// FloodControl интерфейс, который нужно реализовать.
// Рекомендуем создать директорию-пакет, в которой будет находиться реализация.
type FloodControl interface {
	// Check возвращает false если достигнут лимит максимально разрешенного
	// кол-ва запросов согласно заданным правилам флуд контроля.
	Check(ctx context.Context, userID int64) (bool, error)
}

func main() {
	config,err:=config.Read(fileName)
	if err!=nil{
		log.Fatalf("cannot read config: %v",err)
	}

	log.Printf("config: %v=timer, %v=countMax, %v=userId", config.Timer, config.CountMax, config.CheckUserID)

	floodControl:=control.NewFloodControlStruct(int(config.Timer), int(config.CountMax))

	ctx:=context.Background()

	for i:=1;i<=20;i++{
		log.Printf("check number: %v", i)
		ok,err:=floodControl.Check(ctx,config.CheckUserID)
		if err!=nil{
			log.Fatalf("cannot check request: %v",err)
		}
		if !ok{
			log.Println("too much requests")
		}
		time.Sleep(time.Millisecond * 500)
	}
}
package control

import (
	"context"
	"task/flood/storage"

	"time"
)

type FloodControlStruct struct{
	Time time.Duration
	CountMax int
}

func NewFloodControlStruct(timeNeed int, countMax int) *FloodControlStruct{
	FloodControl:=FloodControlStruct{
		Time: time.Duration(timeNeed)*time.Second,
		CountMax: countMax,
	}
	return &FloodControl
}

func (f *FloodControlStruct) Check(ctx context.Context, userID int64) (bool, error) {
	t:=time.Now()

	ok,err:=sqlite.CheckDb(f.CountMax,f.Time,t,int(userID),ctx)
	if err!=nil{
		return false, err
	}

	return ok,nil
}


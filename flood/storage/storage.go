package sqlite

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "modernc.org/sqlite"
)

func CheckDb(countMax int, timer time.Duration, t time.Time, userId int, ctx context.Context) (bool,error){
	log.Printf("open storage")
	db, err := sql.Open("sqlite", "storage.db")
	if err != nil {
		return false,err
	}

	log.Printf("take data from storage")

	rows, err := db.QueryContext(ctx,"SELECT id, time FROM flood WHERE userId=:userId", sql.Named("userId",userId))

	if err != nil {
		return false,err
	}

	mapDb:=make(map[time.Time]int)

	for rows.Next() {
		var timeString string
		var id int

		err := rows.Scan(&id, &timeString)
		if err != nil {
			return false,err
		}

		time, err := time.Parse("2006-01-02 15:04:05.000000000 +0000 UTC", timeString)
		if err != nil {
			return false,err
		}

		mapDb[time]=id
	}

	if err = rows.Err(); err != nil {
		return false,err
	}

	rows.Close()

	log.Printf("strating to delete old")

	tString:=t.Format("2006-01-02 15:04:05.000000000 +0000 UTC")

	time1, err := time.Parse("2006-01-02 15:04:05.000000000 +0000 UTC", tString)
	if err != nil {
		return false,err
	}

	for k,v:=range(mapDb){
		duration:=time1.Sub(k)
		if duration.Seconds()>timer.Seconds(){
			_, err = db.ExecContext(ctx,"DELETE FROM flood WHERE id=:id", sql.Named("id", v))
			if err != nil {
				return false,err
			}
			delete(mapDb,k)
		}
		
	}

	log.Printf("check for too much requests")

	if len(mapDb)>=countMax{
		return false,nil
	}

	log.Printf("take new data to storage")

	_, err = db.ExecContext(ctx,"INSERT INTO flood (time, userId) VALUES (:time, :userId)",
	sql.Named("time",tString),
	sql.Named("userId",userId))
	if err != nil {
		return false,err

	}

	return true,nil
}
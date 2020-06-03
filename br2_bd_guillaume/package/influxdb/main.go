package main

import (
    "fmt"
    "test_influxdb/measuretime"
    "test_influxdb/influxdb"
    "time"
)

func main() {

  start  := time.Now()
  start = measuretime.Measuretime_start()
  fmt.Printf("Création de la base de donnée sur localhost sur le port 8086")
  influxdb.Create_database_influxdb("192.168.0.10", "8086", "test_influxdb_go")
  measuretime.Measuretime_stop(start)

}

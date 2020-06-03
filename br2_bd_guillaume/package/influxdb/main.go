package main

import (
    "fmt"
    "test_influxdb/measuretime"
    "test_influxdb/influxdb"
    "time"
)

func main() {

  /*Création de la base de donnée*/

  start  := time.Now()
  start = measuretime.Measuretime_start()
  fmt.Printf("Création de la base de donnée sur localhost sur le port 8086")
  influxdb.Create_database_influxdb("192.168.0.10", "8086", "test_influxdb_go")
  measuretime.Measuretime_stop(start)
  
  /*Insertion de n éléments*/
  fmt.Printf("Insertion de 11 éléments")
  start = measuretime.Measuretime_start()
  rows := "11"
  ip := "192.168.0.10"
  port := "8086"
  influxdb.Write_database_influxdb(&rows, &ip, &port)
  measuretime.Measuretime_stop(start)
  
  /*Lecture de la database*/
  fmt.Printf("Lecture BD")
  start = measuretime.Measuretime_start()
  influxdb.Read_all_rows_database_influxdb("192.168.0.10", "8086")
  measuretime.Measuretime_stop(start)
  
  /*Lecture de n éléments*/
  fmt.Printf("Lecture du nième élément")
  start = measuretime.Measuretime_start()
  influxdb.Read_one_row_database_influxdb("11", "192.168.0.10", "8086") 
  measuretime.Measuretime_stop(start)
  
  /*Supression du nième élément*/
  fmt.Printf("Suppression du nième élément")
  start = measuretime.Measuretime_start()
  influxdb.Delete_one_row_database_influxdb("11", "192.168.0.10", "8086") 
  measuretime.Measuretime_stop(start)
  
  /*Suppresion de la database*/
  fmt.Printf("Suppression DB")
  start = measuretime.Measuretime_start()
  influxdb.Delete_database_influxdb("192.168.0.10", "8086") 
  measuretime.Measuretime_stop(start)
  

}

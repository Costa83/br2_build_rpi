/* InfluxDB test  : Insertion de N rows */ 
/* go mod init test_influxdb
/ go mod tidy */



package main

import (
    "fmt"
    "test_influxdb/measuretime"
    "test_influxdb/influxdb"
    "time"
)

func main() {


  ip := "192.168.0.10"
  port := "8086"
  bd := "test_influxdb_go"
  lignes := "1"

  fmt.Printf("PROGRAMME DE TEST DE INFLUXDB EN GO")
  fmt.Printf("## Lignes : %s",lignes)
  fmt.Printf("## Port : %s",port)
  fmt.Printf("## Ip : %s",ip)
  fmt.Printf("## BD : %s",bd)

  
 

  /*Création de la base de donnée*/
  start  := time.Now()
  start = measuretime.Measuretime_start()
  fmt.Printf("Création de la base de donnée sur localhost sur le port 8086")
  influxdb.Create_database_influxdb(ip, port, bd)
  measuretime.Measuretime_stop(start)
  
  /*Insertion de n lignes*/
  fmt.Printf("Insertion de 11 lignes")
  start = measuretime.Measuretime_start()
  
  influxdb.Write_database_influxdb(&lignes, &ip, &port)
  measuretime.Measuretime_stop(start)
  
  /*Lecture de la database*/
  fmt.Printf("Lecture BD")
  start = measuretime.Measuretime_start()
  influxdb.Read_all_rows_database_influxdb(ip, port)
  measuretime.Measuretime_stop(start)
  
  /*Lecture de n éléments*/
  fmt.Printf("Lecture du nième élément")
  start = measuretime.Measuretime_start()
  influxdb.Read_one_row_database_influxdb(lignes, ip, port) 
  measuretime.Measuretime_stop(start)
  
  /*Supression du nième élément*/
  fmt.Printf("Suppression du nième élément")
  start = measuretime.Measuretime_start()
  influxdb.Delete_one_row_database_influxdb(lignes, ip, port) 
  measuretime.Measuretime_stop(start)
  
  /*Suppresion de la database*/
  /*fmt.Printf("Suppression DB")
  start = measuretime.Measuretime_start()
  influxdb.Delete_database_influxdb(ip, port)
  measuretime.Measuretime_stop(start)*/
  

}

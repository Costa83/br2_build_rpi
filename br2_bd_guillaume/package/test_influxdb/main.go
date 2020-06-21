/* InfluxDB test  : Programme d'insertion de N lignes dans la base de donnée test_influxdb_go à l'IP 192.168.0.1 et sur le port 8086 */ 


/* go mod init test_influxdb
/  go mod tidy 
/  go build ou  env GOOS=linux GOARCH=arm GOARM=7 go build*/

package main

import (
    "fmt"
    "test_influxdb/measuretime"
    "test_influxdb/influxdb"
    "time"
    "os/exec"
    
)

func main() {
  
  ip := "192.168.0.10"
  port := "8086"
  bd := "test_influxdb_go"
  lignes := "1"
  une_ligne := "1"
  lignes_to_delete := "7 Source"
  lignes_to_print := "6 Source"

  fmt.Printf("PROGRAMME DE TEST D'INFLUXDB EN GO")
  fmt.Println("Entrez le nombre de ligne à tester entre 1 et 25 000 0000 : ")	
  fmt.Scanln(&lignes)
  fmt.Println("OK")
  fmt.Println("Entrez la ligne à afficher ex: 1 Source : ")	
  fmt.Scanln(&lignes_to_print)
  fmt.Println("OK")
  fmt.Println("Entrez la ligne à supprimer ex: 99 Source : ")	
  fmt.Scanln(&lignes_to_delete)
  fmt.Println("OK")	
    
  fmt.Printf("## Lignes : %s",lignes)
  fmt.Printf("## Port : %s",port)
  fmt.Printf("## Ip : %s",ip)
  fmt.Printf("## BD : %s",bd)

  /*Création de la base de donnée*/
  start  := time.Now()
  start = measuretime.Measuretime_start()
  fmt.Printf("Création de la base de donnée")
  influxdb.Create_database_influxdb(ip, port, bd)
  measuretime.Measuretime_stop(start)
  
  /*Insertion de n lignes*/
  fmt.Printf("Insertion de : ", lignes, "lignes")
  start = measuretime.Measuretime_start()
  
  influxdb.Write_database_influxdb(&lignes, &ip, &port)
  measuretime.Measuretime_stop(start)
  

  fmt.Printf("Lecture de l'élément : ", lignes)
  start = measuretime.Measuretime_start()
  influxdb.Read_one_row_database_influxdb(lignes_to_print, ip, port) 
  measuretime.Measuretime_stop(start)
  
  
  fmt.Printf("Suppression du nième élément, ", lignes)
  start = measuretime.Measuretime_start()
  influxdb.Delete_one_row_database_influxdb(lignes_to_delete, ip, port) 
  measuretime.Measuretime_stop(start)
    

 /*Insertion de 1 ligne*/
  fmt.Printf("Insertion de une ligne")
  start = measuretime.Measuretime_start()
  
  influxdb.Write_database_influxdb(&une_ligne, &ip, &port)
  measuretime.Measuretime_stop(start)
  

 /* fmt.Printf("Lecture de la BD")
  start = measuretime.Measuretime_start()
  influxdb.Read_all_rows_database_influxdb(ip, port)
  measuretime.Measuretime_stop(start)*/
  
  
    app := "ls"

    arg0 := "-lah"
    
    cmd := exec.Command(app, arg0)
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println(string(stdout))

  /*fmt.Printf("Suppression de la DB")
  start = measuretime.Measuretime_start()
  influxdb.Drop_database_influxdb(ip, port)
  measuretime.Measuretime_stop(start)*/
 

}

/* LevelDB test  : Insertion de N rows */ 
/* go mod init test_influxdb
/  go mod tidy */

/* LevelDB test  : Programme d'insertion de N lignes dans la base de donnée testleveldb */ 

package main

import (
    "fmt"
    "test_leveldb/measuretime"
    "test_leveldb/leveldb"
    "time"
    "os/exec"
)



func main() {

  lignes := "1"
  one_line := "1"
  path := "testleveldb"
  lignes_to_delete := "7 Source"
  lignes_to_print := "6 Source"
  print_bool := "true"

  

  fmt.Printf("PROGRAMME DE TEST DE LEVELDB EN GO")
  fmt.Println("Entrez le nombre de ligne à tester entre 1 et 25 000 0000 : ")	
  fmt.Scanln(&lignes)
  fmt.Println("OK")
  fmt.Println("Entrez la ligne à afficher ex: 1 Source : ")	
  fmt.Scanln(&lignes_to_print)
  fmt.Println("OK")
  fmt.Println("Entrez la ligne à supprimer ex: 99 Source : ")	
  fmt.Scanln(&lignes_to_delete)
  fmt.Println("OK")
  fmt.Printf("## path : %s",path)
  fmt.Printf("## lignes : %s",lignes)
 

  /*Création de la base de donnée*/
  start  := time.Now()
  start = measuretime.Measuretime_start()
  fmt.Printf("Création de la base de donnée sur le chemin : %s  ", path )
  leveldb.Create_database_leveldb(path)
  measuretime.Measuretime_stop(start)
  
  /*Insertion de n lignes*/
  start = measuretime.Measuretime_start()
  
  leveldb.Write_database_leveldb(&lignes, &path)
  measuretime.Measuretime_stop(start)
   
  
  /*Lecture de la database*/
  fmt.Printf("Lecture BD")
  start = measuretime.Measuretime_start()
  leveldb.Read_all_rows_database_leveldb(&path, &print_bool)
  measuretime.Measuretime_stop(start)
   
  /*Supression du nième élément*/
  fmt.Printf("Suppression du nième élément")
  start = measuretime.Measuretime_start()
  leveldb.Delete_one_row_database_leveldb(lignes_to_delete, path) 
  measuretime.Measuretime_stop(start)
  
   /*Insertion de 1 lignes*/
  fmt.Printf("Insertion de une ligne")
  start = measuretime.Measuretime_start()
  
  leveldb.Write_database_leveldb(&one_line, &path)
  measuretime.Measuretime_stop(start)
  
  /*Lecture de n éléments*/
  fmt.Printf("Lecture du nième élément")
  start = measuretime.Measuretime_start()
  leveldb.Read_one_row_database_leveldb(lignes_to_print, path, "true") 
  measuretime.Measuretime_stop(start)
  
   app := "ls"

   arg0 := "-lah"
   
   
    cmd := exec.Command(app, arg0)
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println(string(stdout))
  
  
  /*Suppresion de la database
  fmt.Printf("Suppression DB")
  start = measuretime.Measuretime_start()
  leveldb.Delete_database_leveldb(path)
  measuretime.Measuretime_stop(start)*/
  

}

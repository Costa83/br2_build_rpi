package leveldb

import (
    "github.com/syndtr/goleveldb/leveldb"
    "fmt"
)
/*Suppresion de n lignes "nb_row_to_delete" de la base de donnée "path" précisée dans le main*/
func Delete_one_row_database_leveldb(nb_row_to_delete string, path string) {

  fmt.Println("Delete one row nb#:  ", nb_row_to_delete)
  db, err := leveldb.OpenFile(path, nil)
	
	if err != nil {
		
		fmt.Println("Error openning LevelDB database: ", err.Error())
	}
	defer db.Close()

 	err = db.Delete([]byte(nb_row_to_delete), nil)





}

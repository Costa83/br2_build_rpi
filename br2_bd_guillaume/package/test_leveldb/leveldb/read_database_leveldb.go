package leveldb

import (
    "github.com/syndtr/goleveldb/leveldb"
    "fmt"
)
/*Lecture de toutes les lignes en base  "path" précisée dans le main et avec un flag permettant d'activer ou non l'affichage de la base de donnée sur le terminal qui est très long et gourmand en ressources*/
func Read_all_rows_database_leveldb(path *string, print *string){

	db, err := leveldb.OpenFile(*path, nil)
	
	if err != nil {
		
		fmt.Println("Error creating LevelDB database: ", err.Error())
	}
	defer db.Close()
	
	iter := db.NewIterator(nil, nil)
        
        for iter.Next() {
        	
        	key := iter.Key()
        	value := iter.Value()

        	
        	if *print == "true" {
        	
	        	fmt.Println("---------------------")
        		key_string := string(key)
        		value_string := string(value)
        		fmt.Println("Key : ", key_string)
        		fmt.Println("Value : ", value_string)
		}
	}
	
  	iter.Release()
  	err = iter.Error()
}






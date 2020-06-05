package leveldb

import (
    "github.com/syndtr/goleveldb/leveldb"
    "fmt"
)

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






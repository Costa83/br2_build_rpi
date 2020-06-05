package leveldb

import (
    "github.com/syndtr/goleveldb/leveldb"
    "fmt"
    "strconv"
    /*"math/rand"*/
)

func Write_database_leveldb(nb_row *string, path *string){

	    fmt.Println("Création LevelDB: ")
	    source_tags := [3]string{"FAKE0","FAKE1","FAKE2"}
	    type_tags := [11]string{"Vitesse","Frequence cardiaque", "Puissance développé", "Cadence de pédalage", "Distance", "Temperature", "Dépense énergétique", "Pourcentage de pente", "Dénivelé positif", "Vitesse du vent", "Luminosité/Panneau solaire"}
	    unit_tags := [11]string{"km_h","bpm","Watts","tmp","km","°C","Kcal","%","D+", "m/s","lux"}
	    
	    db, err := leveldb.OpenFile(*path, nil)
	    if err != nil {
	    
	    	fmt.Println("Error creating LevelDB database: ", err.Error())
	    }
	    defer db.Close()
	    
	    nb_row_int, err := strconv.Atoi(*nb_row)
	    for i:=0;i<(nb_row_int);i++ {
	    	
	    	batch := new(leveldb.Batch)
		batch.Put([]byte("source"), []byte(source_tags[i%3]))
		batch.Put([]byte("type"), []byte(type_tags[i%11]))
		batch.Put([]byte(type_tags[i%11]), []byte(unit_tags[i%11]))
		batch.Put([]byte("idtag"), []byte(strconv.Itoa(i)))
		batch.Put([]byte("value"), []byte(strconv.Itoa(/*rand.Intn(1000)*/1000)))
		/*batch.Put([]byte("id"), []byte(strconv.Itoa(rand.Intn(i))))*/
		batch.Put([]byte("id"), []byte(strconv.Itoa(i)))

		err = db.Write(batch, nil)
		if err != nil {
			fmt.Println("Error: ", err.Error())
   		}
    }
}


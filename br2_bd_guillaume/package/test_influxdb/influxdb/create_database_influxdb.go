package influxdb

import (
    "github.com/influxdata/influxdb1-client/v2"
    "fmt"
)

/*Création de la base de donnée "MyDB" sur une IP "IpAdress"  et port  "PortAddress" précisées dans le main*/

func Create_database_influxdb(IpAdress string, PortAddress string, MyDB string) {

  Addr := "http://" + IpAdress + ":" + PortAddress

	c, err := client.NewHTTPClient(client.HTTPConfig {
		Addr: Addr,
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()


	q := client.NewQuery("create database test_influxdb_go", MyDB, "")
  c.Query(q)

}

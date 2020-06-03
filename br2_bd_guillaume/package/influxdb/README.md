# INSTALLER INFLUXDB SUR LE PC HOTE, DEPUIS ZERO :


* On est sur une architecture client-serveur où ici le client et le serveur est situé sur le PC hote.
* Outils : GRAFANA, InfluxDB



## ETAPE 1: Installer GRAFANA :

Grafana est de loin l'écran de monitoring le plus populaire depuis 2019.
Utilisé par SpaceX, Bloomberg et Booking.com.
Il peut être utilisé pour supervisé différentes ressources :

SQL (MySQL or PostgreSQL), NoSQL, time series BD (InfluxDB / TimescaleDB)

Depuis le site officiel :https://grafana.com/grafana/download 

Ubuntu and Debian(64 Bit)SHA256: 2f307d8ec8ef72aa2831efe1d1d82d5975a801276d6ebadebb9ab0e35a59bb86 

 
$ sudo apt-get install -y adduser libfontconfig1
$ wget https://dl.grafana.com/oss/release/grafana_6.7.2_amd64.deb
$ sudo dpkg -i grafana_6.7.2_amd64.deb



$ sudo systemctl daemon-reload
$ sudo systemctl start grafana-server
$ sudo systemctl status grafana-server

Configure the Grafana server to start at boot:
$ sudo systemctl enable grafana-server.service

Start the server with init.d
$ sudo service grafana-server start



https://grafana.com/docs/grafana/latest/guides/getting_started/
Log in for the first time
Open your web browser and go to http://localhost:3000/. 3000 is the default HTTP port that Grafana listens to if you haven’t configured a different port. 
    1. On the login page, type admin for the username and password. 
    2. Change your password.  « gucos83 » « gucos83 »

Tutoriel supplémentaire :
https://www.sugarbug.fr/atelier/maquettes/centreon_graphite_graphana/grafana/


## ETAPE 2: Installer InfluxDB: Install InfluxDB on Ubuntu 18.04

* Source : https://computingforgeeks.com/install-influxdb-on-ubuntu-18-04-and-debian-9/

* Installation of InfluxDB on Ubuntu 18.04 is done from Influxdata repository. Once the repo is added, the package can then be installed using an apt package manager. Add the InfluxData repository to the file  /etc/apt/sources.list.d/influxdb.list : 

$ echo "deb https://repos.influxdata.com/ubuntu bionic stable" | sudo tee /etc/apt/sources.list.d/influxdb.list


Update apt index and install influxdb 

$ sudo apt-get install -y influxdb

$ sudo apt-get install influxdb-client

Start and enable the service to start on boot up : 

$ sudo systemctl enable --now influxdb

$ sudo systemctl is-enabled influxdb

Check service status:

$ systemctl status influxdb

Open a terminal and check if influxdb is installed :

$ influxd




# INSTALLER GOLANG et les tests de InfluxDB en GO sur  UBUNTU HOST et Cible ARM :

- On est sur une architecture client-serveur où ici le client et le serveur est situé sur le PC hote.


## Prérequis 1 : Installer GRAFANA & INFLUXDB (CLIENT & SERVER) voir README ci-dessus

## Prérequis 2: Installer la commande go pourconstruire le package depuis l’hote:

$ sudo apt install golang-go

$ git clone https://go.googlesource.com/go goroot

$ cd goroot/

$ git checkout go1.14.2

$ git status

$ cd src/

$ ./all.bash 
 
$  go version


$ mkdir test_influxdb

$ git clone .....

$ cp test_influxdb/* test_influxdb    (influxdb/ measuretime/ go.mod go.sum main.go test_influxdb INFLUXDB_FROM_SRC.txt)

$ cd test_influxdb

$ rm -rf go.* 
$ rm -rf test_influxdb

$ go mod init test_influxdb
$ go mod tidy

(génère un go.mod et un go.sum)

$ env GOOS=linux GOARCH=arm GOARM=7 go build

Génère un exécutable arm

$ scp test_influxdb root@192.168.0.11:/usr/


On lance l’exécutable avec ./test_influxdb et les options du programme :

./main --Points=10 --Ip=localhost –Port=8086

########## Test INFLUX DB in Go  ########################
--- Enter the IpAdress of rows to insert : --Ip=...
--- Enter the PortAddress of rows to insert : --Port=...
--- Enter the number of rows to insert : --Points=...
--- Enter the Middle of rows to insert : --Middle=...
--- Enter the Last of rows to insert : --Last=...
##########################################################



### Pour L'intégration d'un package golang sur Buildroot (Pas fini -> Doc à retester=)

On git clone le réperoire golang_test dans le répertoire qui est appelé par BR2 (APP EXTERNES).

Dans package/influx_db_golang/
	Création d’un Config.in
	Création d’un influx_db_go_test.mk


#SUR BR2 :

Créer un répertoire influxdb_go dans package/

	Créer un .config
	Créer un .mk

#POUR RPI

Puis scp dans /usr/bin/ de la RPI de l’excutable.


Trois niveau : repository = ! modules, module!= direcrory, directory = package!= src files






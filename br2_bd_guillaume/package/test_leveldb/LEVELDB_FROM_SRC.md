# INSTALLER LEVELDB SUR LE PC HOTE, DEPUIS ZERO :


* Outils : LevelDB



## ETAPE 2: Installer LevelDB sur Ubuntu  (PC Hôte).

Pas d'installation de logiciels particuliers sur le PC hôte.

# INSTALLER GOLANG et les tests de LevelDB en GO sur  UBUNTU HOST et Cible ARM :

## Prérequis 2: Installer la commande go pour construire le package depuis l’hote:

* Source : https://github.com/syndtr/goleveldb
* Source : https://godoc.org/github.com/syndtr/goleveldb/leveldb
* Source : https://pypi.org/project/leveldb-cli/
* Source : https://computingforgeeks.com/how-to-install-leveldb-on-ubuntu-18-04-ubuntu-16-04/

$ sudo apt install golang-go

$ git clone https://go.googlesource.com/go goroot

$ cd goroot/

$ git checkout go1.14.2

$ git status

$ cd src/

$ ./all.bash 
 
$  go version


$ mkdir test_leveldb

$ git clone .....

$ cp test_leveldb/* test_leveldb   (leveldb/ measuretime/ go.mod go.sum main.go test_leveldb LEVELDB_FROM_SRC.md)

$ cd test_leveldb

$ rm -rf go.* 
$ rm -rf test_leveldb

$ go mod init test_leveldb
$ go mod tidy

(génère un go.mod et un go.sum)

$ env GOOS=linux GOARCH=arm GOARM=7 go build                         (ou go build sur pc hote)
(possibilité de donner un nom à l'executable générer avec l'option -o "nom de l'excutable en sortie".

Génère un exécutable arm

$ ./test_leveldb

On retrouve la base de donnée dans testleveldb/ (pc hote)

$ scp test_leveldb root@192.168.0.11:/usr/


On lance l’exécutable avec ./test_leveldb et les options du programme.






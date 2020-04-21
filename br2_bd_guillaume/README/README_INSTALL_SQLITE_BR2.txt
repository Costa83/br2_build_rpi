#!/bin/bash

# 1
echo 'Installing SQLITE ON BR2'


########################################################
		STEP 0

#Transfert sur sd card avec scp
scp target/usr/bin/sttSqliteInsertNRowsTest  root@10.18.3.214:/usr/bin

> depuis /Documents/STAGE_SMILE_STOCKAGE/DEV_STAGE/stt_env/output$ 
	-> make stt_ihm-reconfigure DEBUG=n
     	-> make stt_ihm-rebuild DEBUG=n
	-> make stt_ihm DEBUG=n
	-> make DEBUG=n


> depuis /Documents/STAGE_SMILE_STOCKAGE/DEV_STAGE/stt_env/output$

	#Upload des app binaires :
	->scp target/usr/bin/sttTestReadRowSqlite             root@10.18.2.128:/usr/bin

-> sur une clé USB montée dans /media/usb : 

-> cd /media/usb/tests/sqlite/

	test.sh 

A l'execution rediriger la sortie vers log_nbrows.txt

Contient les résultats de l'app de test.
########################################################


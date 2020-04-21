INSTALL ENV BR2 DEPUIS LE GIT :
#######################################


 (Utiliser le fichier DEFCONFI propre à la RPI4 et le fichier .config).

1- Se positionner dans ~/SMILE_STAGE/SOURCES_GIT/BR2_BUILD_RPI$


2- git clone https://github.com/Costa83/br2_build_rpi.git


3- cd br2_build_rpi/br2_bd_guillaume/

4- br2_build_rpi/br2_bd_guillaume$ make O=${PWD}/output BR2_EXTERNAL=${PWD} -C /opt/buildroot br2_install_rpi4_defconfig


5- br2_build_rpi/br2_bd_guillaume/output$ make menuconfig


6- On sélectionne les options suivantes ou alors on charge un fichier .config 

On sélectionne différentes fonctionnalités nécessaires au projet :
- Target packags → Networking applications : cocher « dropbear »
On configure un mot de passe Root : 
- System configuration → Root password :
« gucos83»
On configure l’interface réseau pas défault :
- System configuration :
(eth0) Network interface to configure through DHCP.
On configure les paramètres propes à la rpi 
- Target packages → Harware handling → firmware 
« rpi-wifi-firmware ».
« rpi-bt-firmware ».
On active le support du C++ : 
- Toolchain
« Enable C++ support »
.config contient toute la configuration.
project_defconfig contient le minimum pour reconstruite .config
Permet de réutiliser une configuration avec une version postérieur de Buildroot

7- Sauvegarde du fichier .config résultant de la configuration de l'outil menuconfig
/br2_bd_guillaume# cp .config ../configs/.br2_intall_rpi4_config

8- Lancer la compilation : 

/SMILE_STAGE/SOURCES_GIT/BR2_BUILD_RPI/br2_build_rpi/br2_bd_guillaume/output$ make






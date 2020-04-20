#!/bin/bash

# 1
echo 'Installing br2 for raspberry Pi 4'

make
gcc
g++
openssl-dev
ncurces-dev


Téléchargement de buildroot  depuis le site offciel :

wget https://buildroot.org/downloads/buildroot-2018.05.tar.bz2
cp buildroot-2018.05.tar.bz2 ~
cd ~
[~]$ tar -xjf buildroot-2018.05.tar.bz2


Installation de buildroot pour tous projets : 

[~]$ sudo su
[~]# cd /opt
[opt]# git clone git://git.buildroot.net/buildroot
[opt]# cd buidlroot
[buildroot]# mkdir dl
[buildroot]# chmod a+w ./dl


Création du répertoire de travail :

[repertoire projet]$ mkdir stt_env
[repertoire projet]$ cd stt_env
[repertoire projet/stt_env]$ touch external.desc
[repertoire projet/stt_env]$ touch Config.in
[repertoire projet/stt_env]$ touch external.mk
[repertoire projet/stt_env]$ mkdir board configs overlays patches package



Configuration du répertoire de travail : 

[repertoire projet/stt_env]$ cd configs
[repertoire projet/stt_env/configs]$ touch stt_env_defconfig	
[repertoire projet/stt_env/configs]$ cp opt/buildroot/configs/raspberry4_defconfig stt_env_defconfig



Préparation de l’environnement :

[repertoire projet]$make O=${PWD}/output BR2_EXTERNAL=${PWD} -C /opt/buildroot stt_env_defconfig

BR2_DEFCONFIG="$(BR2_EXTERNAL_PROJECT_PATH)/configs/br2_bd_guillaume_defconfig"
BR2_DL_DIR="/opt/buildroot/dl"
BR2_ARCH="arm"

BR2_arm=y
BR2_cortex_a72=y
BR2_ARM_FPU_NEON_VFPV4=y

BR2_TOOLCHAIN_BUILDROOT_CXX=y

BR2_SYSTEM_DHCP="eth0"

# Linux headers same as kernel, a 4.19 series
BR2_PACKAGE_HOST_LINUX_HEADERS_CUSTOM_4_19=y

BR2_LINUX_KERNEL=y
BR2_LINUX_KERNEL_CUSTOM_TARBALL=y
BR2_LINUX_KERNEL_CUSTOM_TARBALL_LOCATION="$(call github,raspberrypi,linux,4f2a4cc501c428c940549f39d5562e60404ac4f7)/linux-4f2a4cc501c428c940549f39d5562e60404ac4f7.tar.gz"
BR2_LINUX_KERNEL_DEFCONFIG="bcm2711"

# Build the DTB from the kernel sources
BR2_LINUX_KERNEL_DTS_SUPPORT=y
BR2_LINUX_KERNEL_INTREE_DTS_NAME="bcm2711-rpi-4-b"

BR2_LINUX_KERNEL_NEEDS_HOST_OPENSSL=y

BR2_PACKAGE_RPI_FIRMWARE=y
BR2_PACKAGE_RPI_FIRMWARE_VARIANT_PI4=y

# Required tools to create the SD image
BR2_PACKAGE_HOST_DOSFSTOOLS=y
BR2_PACKAGE_HOST_GENIMAGE=y
BR2_PACKAGE_HOST_MTOOLS=y

# Filesystem / image
BR2_TARGET_ROOTFS_EXT2=y
BR2_TARGET_ROOTFS_EXT2_4=y
BR2_TARGET_ROOTFS_EXT2_SIZE="120M"
# BR2_TARGET_ROOTFS_TAR is not set
BR2_ROOTFS_POST_BUILD_SCRIPT="board/raspberrypi4/post-build.sh"
BR2_ROOTFS_POST_IMAGE_SCRIPT="board/raspberrypi4/post-image.sh"
BR2_ROOTFS_POST_SCRIPT_ARGS="--add-miniuart-bt-overlay"


Ajout du projet à un git :

Permet d’identifer les étapes de construction du système.
[repertoire projet]$ git init
[repertoire projet]$ git add external.desc
[repertoire projet]$ git add external.mk
[repertoire projet]$ git add Config.in
Permet d’identifer les étapes de construction du système.
[project]$ git remote add origin  https://git.company.com/project
[project]$ git push origin master
$ git show –pretty=--name-only     #Pour afficher les fichiers qui ont été commit.


Utiliser les outils proposer par Buildroot :

external.mk : 

include $(sort $(wildcard $(BR2_EXTERNAL_PROJECT_PATH)
>>/package/*/*.mk))

Config.in : 

source "$BR2_EXTERNAL_PROJECT_PATH/package/my_app/Config.in"
source "$BR2_EXTERNAL_PROJECT_PATH/package/my_dev/Config.in"


[project]$ cd output
[output]$ make menuconfig

Génération de l'image système :


Vérifier que le fichier bcm2711 est bien sélectionné, que le fichier dts est bien sélectionné, ainsi que le fichier stt_env_defconfig, ainsi que l’option build device tree blob.

On sélectionne différentes fonctionnalités nécessaires au projet :
- Target packags → Networking applications : cocher « dropbear »
On configure un mot de passe Root : 
- System configuration → Root password :
« gucos83»
On configure l’interface réseau pas défault :
- System configuration :
(eth0) Network interface to configure through DHCP.
On configure les paramètres propes à la rpi 
- Target packages → Harware handling → firmware 
« rpi-wifi-firmware ».
« rpi-bt-firmware ».

On active le support du C++ : 
- Toolchain
« Enable C++ support »





.config contient toute la configuration.
project_defconfig contient le minimum pour reconstruite .config
Permet de réutiliser une configuration avec une version postérieur de Buildroot




$ make











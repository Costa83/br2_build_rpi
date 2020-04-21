################################################################################
#
# InfluxDB
#
################################################################################

INFLUXDB_VERSION = 1.7.10
INFLUXDB_SOURCE = influxdb-$(INFLUXDB_VERSION)_linux_armhf.tar.gz
INFLUXDB_SITE =  https://dl.influxdata.com/influxdb/releases
STT_IHM_SITE_METHOD = wget
INFLUXDB_LICENSE = MIT
# INFLUXDB_STAGING = YES

## URL https://github.com/influxdata/influxdb/archive/v1.2.3.tar.gz

define INFLUXDB_BUILD_CMDS
	echo "Source Compile"
endef

define INFLUXDB_INSTALL_STADING_CMDS
	$(INSTALL) -D -m 0755 $(@D)/
endef

$(eval $(generic-package))

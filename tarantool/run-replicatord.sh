#!/bin/sh
/opt/tarantool/replicatord \
	-c /opt/tarantool/replicatord.yaml \
	-l /var/lib/tarantool/log/replicatord.log \
	-i /var/lib/tarantool/pid/replicatord.pid &

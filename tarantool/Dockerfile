FROM tarantool/tarantool:2.7.2

COPY run-replicatord.sh /opt/tarantool
COPY replicatord.yaml /opt/tarantool
COPY replicatord /opt/tarantool
COPY box.lua /opt/tarantool

ENTRYPOINT ["docker-entrypoint.sh"]

HEALTHCHECK CMD tarantool_is_up

EXPOSE 3301

CMD ["tarantool", "/opt/tarantool/box.lua"]

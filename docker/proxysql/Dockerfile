FROM proxysql/proxysql:2.0.13

RUN echo 'Acquire { https::Verify-Peer false }' > /etc/apt/apt.conf.d/99influxdata-cert

# Update packages
RUN apt update -y --fix-missing --allow-unauthenticated

# Install common packages
RUN apt install -y mysql-client --allow-unauthenticated

# Set custom config
ADD proxysql.cnf /etc/

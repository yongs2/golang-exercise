# database

## 01.mariadb

- start mariadb
```sh
./docker_mariadb.sh start

docker inspect mariadb | grep IPAddr

export MARIADB_HOST=172.17.0.4;
export MARIADB_PORT=3306;

go get github.com/go-sql-driver/mysql
```

# spec
```shell script
docker pull mariadb
```
```shell script
docker run -p 0.0.0.0:3306:3306  --name my-db -e MARIADB_ROOT_PASSWORD=password -d mariadb:latest
```
```shell script
docker exec -it my-db mysql -uroot -ppassword
```
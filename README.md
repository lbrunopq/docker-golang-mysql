# Docker + Golang + Mysql

## Network

```
- docker network create -d bridge pfa_net
```

## MySQL

```
- docker run -d --name db --network pfa_net -p 3306:3306 -v $(pwd)/mysql-data:/var/lib/mysql --env MYSQL_ROOT_PASSWORD=mysql --env MYSQL_DATABASE=challenge mysql:5.7
- docker exec -it db bash
- mysql -uroot -p <ENTER>
  digite a senha: mysql <ENTER>
  create database challenge; <ENTER>
  <EXIT>
  <EXIT>
```

## App

```
cd app
- docker build -t lbrunoq/golang-mysql .
- docker run -d --name golang_mysql --network pfa_net -p 8000:8000 --env MYSQL_ROOT_PASSWORD=mysql --env MYSQL_USER=root --env MYSQL_HOST=db --env MYSQL_PORT=3306 --env MYSQL_DATABASE=challenge lbrunoq/golang-mysql
```

## Nginx

```
- cd nginx
- docker build -t lbrunoq/nginx .
- docker run -d --name nginx --network pfa_net -p 8080:80 lbrunoq/nginx
```

## Publicando as imagens

```
- docker push lbrunoq/golang-mysql
- docker push lbrunoq/nginx
```

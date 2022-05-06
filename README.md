# bookstore_oauth-api
Oauth API


### Setup cassandra with docker
```bash
$ docker run --name cassandra -d \
    -e CASSANDRA_BROADCAST_ADDRESS=[YOUR-NODE-IP-ADDRESS] \
    -p 7000:7000 \
    -p 9042:9042 \
    --restart always\
    cassandra
```
Create keyspace<br/>
`create keyspace oauth with replication = {'class':'SimpleStrategy','replication_factor':1}`<br/>
Create table<br/>
`create table access_tokens(access_token varchar primary key, user_id bigint, client_id bigint, expires bigint);`<br/>
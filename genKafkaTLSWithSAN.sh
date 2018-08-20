#!/bin/bash

rm -rf kafkaTLSserver kafkaTLSclient
mkdir kafkaTLSserver kafkaTLSclient

cd ./kafkaTLSserver
# make ca
echo ca
openssl genrsa -out ca.key 2048
openssl req -x509 -new -nodes -key ca.key -subj "/C=CN/ST=Beijing/L=Beijing/CN=peersafe.com" -days 3650 -out ca.crt

# make server
echo server
openssl genrsa -out server.key 2048
openssl req -new -key server.key -subj "/C=CN/ST=Beijing/L=Beijing/CN=kafkaserver" -out server.csr -reqexts SAN -config <(cat ../openssl.cnf.base <(printf "[SAN]\nsubjectAltName=DNS:kafka0,DNS:kafka1,DNS:kafka2,DNS:kafka3"))
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 3650 -extfile <(cat ../openssl.cnf.base <(printf "[SAN]\nsubjectAltName=DNS:kafka0,DNS:kafka1,DNS:kafka2,DNS:kafka3")) -extensions SAN

# make client
echo client
openssl genrsa -out client.key 2048
openssl req -new -key client.key -subj "/C=CN/ST=Beijing/L=Beijing/CN=client" -out client.csr
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 3650

# do transfer
echo transfer
openssl pkcs12 -export -in server.crt -inkey server.key -out server.pk12 -name server -passout pass:test1234
printf "test1234\ntest1234\nY\n\n" | keytool -importkeystore -deststorepass test1234 -destkeypass test1234 -destkeystore server.keystore.jks -srckeystore server.pk12 -srcstoretype PKCS12 -srcstorepass test1234 -alias server

# add singed cert to server jks
echo keytool
printf "test1234\ntest1234\nY\n\n" | keytool -keystore server.truststore.jks -alias CARoot -import -file ca.crt
printf "test1234\n\n" | keytool -keystore server.truststore.jks -alias server -import -file server.crt
printf "test1234\n\n" | keytool -keystore server.truststore.jks -alias client -import -file client.crt

cp ca.crt ../kafkaTLSclient
mv client.* ../kafkaTLSclient

cd ..

# IP SAN
#openssl req -new -key server.key -subj "/C=CN/ST=Beijing/L=Beijing/CN=kafkaserver" -out server.csr -reqexts SAN -config <(cat ./openssl.cnf.base <(printf "[SAN]\nsubjectAltName=IP:1.1.1.1,DNS:kafka0,DNS:kafka1,DNS:kafka2,DNS:kafka3"))

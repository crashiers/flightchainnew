#!/bin/bash

# Exit on first error
set -e

CHAINCODE_NAME=flightchaincode
CHAINCODE_VERSION=1.0
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

starttime=$(date +%s)
LANGUAGE=${1:-"node"}

cd $DIR
cd ./node
yarn
yarn run clean
yarn run build
CC_SRC_PATH=/opt/gopath/src/github.com/firstchaincode/node

# Clean the keystore
rm -rf ./hfc-key-store

cd $DIR

docker exec -e "CORE_PEER_LOCALMSPID=AirlineMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/airline.flight.com/users/Admin@airline.flight.com/msp" a896c63c68cc peer chaincode install -n "$CHAINCODE_NAME" -v "$CHAINCODE_VERSION" -p "$CC_SRC_PATH" -l "$LANGUAGE"

docker exec -e "CORE_PEER_LOCALMSPID=AirportMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/airport.flight.com/users/Admin@airport.flight.com/msp" a896c63c68cc peer chaincode install -n "$CHAINCODE_NAME" -v "$CHAINCODE_VERSION" -p "$CC_SRC_PATH" -l "$LANGUAGE"


docker exec -e "CORE_PEER_LOCALMSPID=AirlineMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/airline.flight.com/users/Admin@airline.flight.com/msp" a896c63c68cc peer chaincode instantiate -o orderer.example.com:7050 -C flightchannel -n "$CHAINCODE_NAME" -l "$LANGUAGE" -v "$CHAINCODE_VERSION" -c '{"function":"init","Args":["'$CHAINCODE_VERSION'"]}'

CONTAINER_NAME="dev-peer0.airline.flight.com-${CHAINCODE_NAME}"

CID=$(docker ps -q -f status=running -f name=^/${CONTAINER_NAME})

while [ ! "${CID}" ]; do
    CID=$(docker ps -q -f status=running -f name=^/${CONTAINER_NAME})
    echo "$CONTAINER_NAME not found";
    sleep 3;
done;

sleep 3;

docker exec -e "CORE_PEER_LOCALMSPID=AirlineMSP" -e "CORE_PEER_MSPCONFIGPATH/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/airline.flight.com/users/Admin@airline.flight.com/msp" a896c63c68cc peer chaincode invoke -o orderer.example.com:7050 -C flightchannel -n "$CHAINCODE_NAME" -c '{"function":"initLedger","Args":[""]}'

printf "\nTotal setup execution time : $(($(date +%s) - starttime)) secs ...\n\n\n"

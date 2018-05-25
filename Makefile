SHELL := /bin/bash

proto:
	protoc --go_out=. markets/*.proto
	protoc --go_out=. client/*.proto
	protoc --go_out=Mmarkets/market.proto=github.com/cryptowatch/proto/markets,Mmarkets/pair.proto=github.com/cryptowatch/proto/markets:. stream/*.proto

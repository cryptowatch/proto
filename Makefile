SHELL := /bin/bash

proto:
	protoc --go_out=. markets/*.proto
	protoc --go_out=. client/*.proto
	protoc --go_out=Mmarkets/market.proto=code.cryptowat.ch/proto/markets,Mmarkets/pair.proto=code.cryptowat.ch/proto/markets,Mmarkets/asset.proto=code.cryptowat.ch/proto/markets:. stream/*.proto
	protoc --go_out=Mstream/stream.proto=code.cryptowat.ch/proto/stream:. broker/*.proto

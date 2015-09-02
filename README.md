Gr
========

[![GoDoc](http://godoc.org/github.com/xzip/gr?status.png)](http://godoc.org/github.com/xzip/gr)
[![License](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/xzip/gr)

A Go (Golang) Redis 3.2 Client that aims to be EASY to use, and supports:
- Transparent Connection Pooling
- Pipelining
- Transactions
- Cluster Suppport
- Sentinel Support
- Support all 3.2 commands:
  - Cluster
  - Connection
  - Geo (Redis 3.2)
  - Hashes
  - HyperLogLog
  - Keys
  - Lists
  - Pub/sub
  - Scripting
  - Server
  - Sets
  - SortedSets
  - Strings
  - Transactions


#Testing
Create 3 redis-server

redis-server

redis-server --port 6780 --slaveof localhost 6379

redis-server --port 7000


Run go test -cover

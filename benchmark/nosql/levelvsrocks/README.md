# levelvsrocks
It's the file for leveldb vs rocksdb.

In this file, we'll test four types in each db:

* Batch write sequential key data.
* Read sequential key data.
* Batch delete sequential key data.

* Batch write random key data.
* Read random key data.
* Batch delete random key data.

# install go packages

```BASH
$ go get github.com/syndtr/goleveldb/leveldb
$ go get github.com/tecbot/gorocksdb
```

# run

```BASH
$ go run benchmark.go [total] // total default: 100000

// result
Total data count: 100000
LevelDB====================================================================================
Success to put sequential data, time pass: 138.208953ms
Success to get sequential data, time pass: 160.143581ms
Success to delete sequential data, time pass: 158.125617ms
====================================================================================
Numbers length: 100000, Random numbers length: 100000
Success to put random data, time pass: 181.3681ms
Success to get random data, time pass: 231.175262ms
Success to delete random data, time pass: 214.814096ms
RocksDB====================================================================================
Success to put sequential data, time pass: 47.899502ms
Success to get sequential data, time pass: 96.112914ms
Success to delete sequential data, time pass: 78.944891ms
====================================================================================
Numbers length: 100000, Random numbers length: 100000
Success to put random data, time pass: 126.584311ms
Success to get random data, time pass: 140.281799ms
Success to delete random data, time pass: 120.892461ms

```
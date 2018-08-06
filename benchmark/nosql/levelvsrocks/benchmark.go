package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/tecbot/gorocksdb"
)

// run level db benchmark
func runLevelBenchmark(db *leveldb.DB, numbers []int, isRandom bool) {
	batch := new(leveldb.Batch)
	start := time.Now()
	totalLen := len(numbers)
	typeString := ""

	if isRandom {
		typeString = "random"
	} else {
		typeString = "sequential"
	}

	for i := 0; i < totalLen; i++ {
		batch.Put([]byte(strconv.Itoa(numbers[i])), []byte(strconv.Itoa(i)))
	}

	err := db.Write(batch, nil)
	end := time.Now()
	difference := end.Sub(start)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Success to put %5s data, time pass: %v\n", typeString, difference)
	}
	for i := 0; i < totalLen; i++ {
		data, _ := db.Get([]byte(strconv.Itoa(numbers[i])), nil)
		if data != nil {
		}
	}

	end2 := time.Now()

	if err != nil {
		fmt.Println(err)
	} else {
		difference = end2.Sub(end)
		fmt.Printf("Success to get %s data, time pass: %v\n", typeString, difference)
	}

	batch.Reset()

	for i := 0; i < totalLen; i++ {
		batch.Delete([]byte(strconv.Itoa(numbers[i])))
	}

	err = db.Write(batch, nil)
	end3 := time.Now()

	if err != nil {
		fmt.Println(err)
	} else {
		difference = end3.Sub(end2)
		fmt.Printf("Success to delete %s data, time pass: %v\n", typeString, difference)
	}
}

// run rocks db benchmark
func runRocksBenchmark(db *gorocksdb.DB, numbers []int, isRandom bool) {
	batch := gorocksdb.NewWriteBatch()
	start := time.Now()
	totalLen := len(numbers)
	typeString := ""

	if isRandom {
		typeString = "random"
	} else {
		typeString = "sequential"
	}

	for i := 0; i < totalLen; i++ {
		batch.Put([]byte(strconv.Itoa(numbers[i])), []byte(strconv.Itoa(i)))
	}

	wo := gorocksdb.NewDefaultWriteOptions()
	err := db.Write(wo, batch)
	end := time.Now()
	difference := end.Sub(start)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Success to put %5s data, time pass: %v\n", typeString, difference)
	}

	ro := gorocksdb.NewDefaultReadOptions()

	for i := 0; i < totalLen; i++ {
		data, _ := db.Get(ro, []byte(strconv.Itoa(numbers[i])))
		if data != nil {
		}
	}

	end2 := time.Now()

	if err != nil {
		fmt.Println(err)
	} else {
		difference = end2.Sub(end)
		fmt.Printf("Success to get %s data, time pass: %v\n", typeString, difference)
	}

	batch.Clear()

	for i := 0; i < totalLen; i++ {
		batch.Delete([]byte(strconv.Itoa(numbers[i])))
	}

	err = db.Write(wo, batch)
	end3 := time.Now()

	if err != nil {
		fmt.Println(err)
	} else {
		difference = end3.Sub(end2)
		fmt.Printf("Success to delete %s data, time pass: %v\n", typeString, difference)
	}
}

func main() {
	total := 100000
	db, err := leveldb.OpenFile("./ldb", nil)
	numbers := make([]int, total)

	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < total; i++ {
		numbers[i] = i
	}

	fmt.Printf("Total data count: %d\n", total)
	fmt.Println("LevelDB====================================================================================")

	runLevelBenchmark(db, numbers, false)

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	randoms := make([]int, total)

	for i := 0; i < total; i++ {
		randoms[i] = random.Intn(total)
	}

	fmt.Println("====================================================================================")
	fmt.Printf("Numbers length: %d, Random numbers length: %d\n", len(numbers), len(randoms))

	runLevelBenchmark(db, randoms, true)

	defer db.Close()

	fmt.Println("RocksDB====================================================================================")

	bbto := gorocksdb.NewDefaultBlockBasedTableOptions()
	opts := gorocksdb.NewDefaultOptions()
	opts.SetBlockBasedTableFactory(bbto)
	opts.SetCreateIfMissing(true)
	rdb, _ := gorocksdb.OpenDb(opts, "./rdb")

	runRocksBenchmark(rdb, numbers, false)

	fmt.Println("====================================================================================")
	fmt.Printf("Numbers length: %d, Random numbers length: %d\n", len(numbers), len(randoms))

	runRocksBenchmark(rdb, randoms, true)

	defer rdb.Close()

}

package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/sudo-suhas/bulk-marshal/jsonutil"
	"github.com/sudo-suhas/bulk-marshal/model"
	"github.com/sudo-suhas/bulk-marshal/strategies"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func unmarshalFile(file string, v interface{}) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	return jsonutil.Unmarshal(data, v)
}

func main() {
	var bulk []interface{}
	var ps []*model.Person
	check(unmarshalFile("strategies/testdata/people.json", &ps))
	bulk = make([]interface{}, len(ps))
	for i, p := range ps {
		bulk[i] = p
	}

	start := time.Now()
	_, err := strategies.MarshalSerial(bulk)
	check(err)
	fmt.Println("Finished serial marshalling", time.Since(start))
	start = time.Now()
	_, err = strategies.MarshalParallel(8, bulk)
	check(err)
	fmt.Println("Finished parallel marshalling", time.Since(start))
	start = time.Now()
	_, err = strategies.MarshalPooled(8, bulk)
	check(err)
	fmt.Println("Finished pooled marshalling", time.Since(start))
}

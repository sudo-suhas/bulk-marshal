package strategies

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"testing"

	"github.com/sudo-suhas/bulk-marshal/jsonutil"
	"github.com/sudo-suhas/bulk-marshal/model"
)

var bulk []interface{}

func init() {
	var ps []*model.Person
	check(unmarshalFile("testdata/people.json", &ps))
	bulk = make([]interface{}, len(ps))
	for i, p := range ps {
		bulk[i] = p
	}
}

var result [][]byte

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

func BenchmarkMarshalStrategies(b *testing.B) {
	var (
		r   [][]byte
		err error
	)

	b.Run("Serial", func(b *testing.B) {
		b.ReportAllocs()

		// run the serial marshalling function b.N times
		for n := 0; n < b.N; n++ {
			r, err = MarshalSerial(bulk)
			check(err)
		}
		result = r
	})

	cpus := runtime.NumCPU()
	var strategies = []struct {
		name string
		fn   func(int, []interface{}) ([][]byte, error)
	}{
		{"Parallel", MarshalParallel},
		{"Pooled", MarshalPooled},
	}
	for _, s := range strategies {
		for i := 1; i <= 10; i++ {
			b.Run(fmt.Sprintf("%s/WorkerCnt-%dx", s.name, i), func(b *testing.B) {
				b.ReportAllocs()

				// run the parallel marshalling function b.N times
				for n := 0; n < b.N; n++ {
					r, err = s.fn(cpus*i, bulk)
					check(err)
				}
				result = r
			})
		}
	}
}

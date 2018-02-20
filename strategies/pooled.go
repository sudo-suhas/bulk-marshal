package strategies

import (
	"fmt"

	pool "gopkg.in/go-playground/pool.v3"

	"github.com/sudo-suhas/bulk-marshal/jsonutil"
)

type iData struct {
	i    int
	data []byte
}

func poolWorker(i int, val interface{}) pool.WorkFunc {
	return func(wu pool.WorkUnit) (interface{}, error) {
		if wu.IsCancelled() {
			// return values not used
			return nil, nil
		}

		data, err := jsonutil.Marshal(val)
		if err != nil {
			return nil, err
		}

		return &iData{i, data}, nil
	}
}

func MarshalPooled(poolSize int, bulk []interface{}) ([][]byte, error) {
	p := pool.NewLimited(uint(poolSize))
	defer p.Close()

	batch := p.Batch()

	go func() {
		for i, val := range bulk {
			batch.Queue(poolWorker(i, val))
		}

		// DO NOT FORGET THIS OR GOROUTINES WILL DEADLOCK
		// if calling Cancel() it calles QueueComplete() internally
		batch.QueueComplete()
	}()

	dataSlice := make([][]byte, len(bulk))
	for res := range batch.Results() {
		if err := res.Error(); err != nil {
			batch.Cancel()
			return nil, err
		}

		iRes, ok := res.Value().(*iData)
		if !ok {
			return nil, fmt.Errorf("bulk marshal got unexpected type: %T", res.Value())
		}
		dataSlice[iRes.i] = iRes.data
	}

	return dataSlice, nil
}

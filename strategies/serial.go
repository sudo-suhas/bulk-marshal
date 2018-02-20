package strategies

import "github.com/sudo-suhas/bulk-marshal/jsonutil"

func MarshalSerial(bulk []interface{}) ([][]byte, error) {
	dataSlice := make([][]byte, len(bulk))
	var err error

	for idx, val := range bulk {
		if dataSlice[idx], err = jsonutil.Marshal(val); err != nil {
			return nil, err
		}
	}

	return dataSlice, nil
}

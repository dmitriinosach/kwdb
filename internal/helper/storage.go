package helper

import (
	"hash/fnv"
	"kwdb/app/errorpkg"
)

func HashFunction(key string, partitions int) (int, error) {
	h := fnv.New32a()
	h.Write([]byte(key))
	number := int(h.Sum32()) % partitions

	if number > partitions {
		return 0, errorpkg.ErrHashFunctionIndexOutRange
	}

	return number, nil
}

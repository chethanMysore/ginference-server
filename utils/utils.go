package utils

import (
	"bytes"
	"io"
	"os"
)

type FilterError interface {
	ErrEmptyList() error
	ErrNotFound(params ...any) error
	// ErrNotFound(T any) string
}

func Filter[T FilterError](ipList []T, cond func(t T) bool) ([]T, error) {
	opList := []T{}
	if (len(ipList) > 0) && (cap(ipList) > 0) {
		for _, v := range ipList {
			if cond(v) {
				opList = append(opList, v)
			}
		}
		if (len(opList) > 0) && (cap(opList) > 0) {
			return opList, nil
		} else {
			var tErr T
			return []T{}, tErr.ErrNotFound()
		}
	} else {
		var tErr T
		return []T{}, tErr.ErrEmptyList()
	}
}

func First[T FilterError](ipList []T) (T, error) {
	if (len(ipList) > 0) && (cap(ipList) > 0) {
		return ipList[0], nil
	} else {
		var none T
		return none, none.ErrEmptyList()
	}
}

func ReadConfig(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func() {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()
	b := make([]byte, 32)
	strBuff := bytes.NewBufferString("")
	for {
		n, err := file.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		if n > 0 {
			_, err := strBuff.Write(b[:n])
			if err != nil {
				return "", err
			}
		}
	}
	return strBuff.String(), nil
}

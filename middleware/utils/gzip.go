package utils

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

func Gunzip(byteArray []byte) []byte {
	buf := bytes.NewBuffer(byteArray)
	gzipReader, _ := gzip.NewReader(buf)
	newByteArray, _ := ioutil.ReadAll(gzipReader)
	gzipReader.Close()
	return newByteArray
}

func Gzip(byteArray []byte) []byte {
	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf)
	gzipWriter.Write(byteArray)
	gzipWriter.Close()
	return buf.Bytes()
}

package main

import (
	"flag"
	"myblogsearch/crawler/convert"
	"os"
)

func main() {
	inputFilenamePtr := flag.String("i", "feed.xml", "input filename")
	outputFilenamePtr := flag.String("o", "feed.json", "output filename")
	flag.Parse()
	xmlFile, err := os.Open(*inputFilenamePtr)
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()
	jsonFile, err := os.Create(*outputFilenamePtr)
	if err != nil {
		panic(err)
	}
	convert.XmlToJson(xmlFile, jsonFile)
	jsonFile.Close()
}

package main

import (
	"flag"
	"fmt"
	"github.com/10cella/yomo-y3-stress-testing/internal/generator"
	"os"
)

func main() {
	file := flag.String("file", "./assets/json/json_test_data.bin", "file path of json data")
	count := flag.Int("count", 1000, "count of data")
	flag.Parse()

	f, _ := os.Create(*file)
	for i := 1; i <= *count; i++ {
		_, _ = f.Write(append(generator.NewJsonTestData().GenData(), []byte("\n")...))
		//_, _ = f.Write(append(generator.NewJsonTestData().GenDataBy(3), []byte("\n")...))
	}
	_ = f.Close()

	fmt.Printf("write data to file successfully, file=%v, count=%v\n", *file, *count)
}

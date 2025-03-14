package ali

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString("nshdk ndir lik day day")
	data, err := os.ReadFile("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	jm, err := os.ReadFile("ali.sh")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jm))
	date, err := os.ReadDir(".")

	for _, ele := range date {
		fmt.Println(ele.Name())
	}
	os.Remove("example.txt")
}

package pool

import "fmt"

func main() {
	// needed to split to N routines
	for _, domain := range getData() {
		if err := processDomain(domain); err != nil {
			fmt.Println("failed to process:", domain)
		}
	}
}

func processDomain(domain string) error {
	// do some LONG magic
	return nil
}

func getData() []string {
	res := []string{}
	for i := 0; i < 1000; i++ {
		res = append(res, fmt.Sprintf("%d.com", i))
	}
	return res
}

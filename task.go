package main // This is a basic Golang function that eliminate odd numbers in a range
import "fmt" // Basic format library included

func main() {
	var startofRange int
	var endofRange int
	fmt.Scan(&startofRange)
	fmt.Scan(&endofRange)
	for i := startofRange; i <= endofRange; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

}

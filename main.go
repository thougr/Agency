package main
import (
	"fmt"
	"time"
)
func asynchronous() {

	customer := GetCustomerDetails()
	destinations := GetDestinations(customer)
	var Ins [10]Info

	quotes := [10]chan Quoting{}
	weathers := [10]chan Weather{}

	for i := range quotes {
		quotes[i] = make(chan Quoting)
	}

	for i := range weathers {
		weathers[i] = make(chan Weather)
	}

	for index, dest := range destinations {
		i := index
		d := dest
		go func() {
			quotes[i] <- GetQuote(d)
		}()

		go func() {
			weathers[i] <- GetWeather(d)
		}()
	}

	for index, dest := range destinations {
		Ins[index] = Info{Destinations:dest, Quote:<-quotes[index], Weather:<-weathers[index]}
	}
}


func synchronous() {

	var Ins [10]Info
	customers := GetCustomerDetails()
	destinations := GetDestinations(customers)

	for index, dest := range destinations {
		q := GetQuote(dest)
		w := GetWeather(dest)
		Ins[index] = Info{Destinations:dest, Quote:q, Weather:w}
	}

}
func main()  {
  var c int
  fmt.Println("choose the modes you want")
  fmt.Println("0 synchronous")
  fmt.Println("1 asynchronous")
  fmt.Scanln(&c)
  if c == 0 {
    t1 := time.Now()
    synchronous()
    elapsed1 := time.Since(t1)
    fmt.Println("Naive Approach Time: ", elapsed1)
  } else if c == 1 {
    t2 := time.Now()
    asynchronous()
    elapsed2 := time.Since(t2)
    fmt.Println("Optimized Approach Time: ", elapsed2)
  }
}

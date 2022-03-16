package main

import "fmt"

func main() {
	fmt.Println("Deferred Execution in Golang")
	Space()

	// Simulate opening a connection and deferring its close
	OpenDBCon()
	defer CloseDBCon()

	// Simulate opening a connection to a grpc service and deferring its close
	OpenGrpcCon()
	defer CloseGrpcCon()

	// Simulate opening a connection to a message queue service and deferring its close
	OpenQueueCon()
	defer CloseQueueCon()

	Space()
	// Do some work
	SomeWork()

	Space()
	fmt.Println("Main function is ending...")
	Space()
}

// OpenDBCon mocks connecting to an actual DB using
// its driver.
func OpenDBCon() {
	fmt.Println("Connected to database !")
}

// CloseDBCon mocks disconnecting from an actual DB
func CloseDBCon() {
	fmt.Println("Closing the connection to DB...")
}

// OpenGrpcCon mocks connecting to a GRPC service
func OpenGrpcCon() {
	fmt.Println("Connected to grpc service!")
}

// CloseGrpcCon mocks closing the connection to a GRPC service
func CloseGrpcCon() {
	fmt.Println("Closing the connection to grpc service...")
}

// OpenGrpcCon mocks connecting to a GRPC service
func OpenQueueCon() {
	fmt.Println("Connected to message queue service!")
}

// CloseGrpcCon mocks closing the connection to a GRPC service
func CloseQueueCon() {
	fmt.Println("Closing the connection to message queue service...")
}

// SomeWork is literally just some processing
func SomeWork() {
	fmt.Println("Calculating the sum of million numbers...")
	var sum int
	for i := 1; i <= 10000000; i++ {
		sum += i
	}
	fmt.Println("Sum of first million numbers:", sum)
}

func Space() {
	fmt.Println("-----------------------------------------------------------")
}

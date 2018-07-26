package hw

// HelloWorld interface lists methods for HelloWorld microservice.
type HelloWorld interface {
	// Hello method returns a hello greetings message addressed to a person/bot with given name.
	Hello(name string) (*HelloRs, error)
}

// HelloRs is the response for Hello method.
type HelloRs struct {
	// Greetings.
	Greetings string
}

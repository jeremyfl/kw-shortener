package helpers

// Response basic struct of every response
type Response struct {
	Data    interface{} `json:"data" xml:"data"`       // Every object of data goes to this
	Message string      `json:"message" xml:"message"` // Define the response and object
}


package activemq

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee struct {
	ID     float32 `json:"id"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
}

func (q Employee) String() string {
	return fmt.Sprintf("ID: %.0f, Name: %s, Salary %.2f", q.ID, q.Name, q.Salary)
}

func TestMQClient(t *testing.T) {

	addr := os.Getenv("BROKER_ADDR")
	queueName := "/helloQueue"

	activeMQ := NewClient(addr)
	assert.NoError(t, activeMQ.Publish(queueName, []byte("This is a simple message")))

	err := activeMQ.Check()
	assert.NoError(t, err)

	// activeMQ.Subscribe("/employeeQueue", func(err error, message []byte) {
	// 	var employee Employee
	// 	if err := json.Unmarshal(message, &employee); err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("received", employee)
	// 	//publish the data back onto another queue
	// 	activeMQ.Publish("helloQueue", message)
	// })
}

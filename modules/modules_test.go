package modules

import (
	"fmt"
	"testing"

	"github.com/goburrow/modbus"
)

func TestModules(t *testing.T) {
	handler := modbus.NewTCPClientHandler("localhost:502")
	// Connect manually so that multiple requests are handled in one session
	err := handler.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer handler.Close()
	client := modbus.NewClient(handler)

	_, err = client.WriteMultipleRegisters(0, 4, []byte{0, 10, 0, 255, 1, 5, 0, 3})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	results, err := client.ReadHoldingRegisters(0, 3)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("results %v\n", results)
}

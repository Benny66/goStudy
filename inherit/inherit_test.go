package inherit

import (
	"fmt"
	"testing"
)

func TestInherit(t *testing.T) {
	teacher := Teacher{}
	teacher.ShowA()
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("show a")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("show b")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher show b")
}

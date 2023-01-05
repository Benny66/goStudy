package sensitive

import (
	"fmt"
	"testing"

	"github.com/importcjj/sensitive"
)

func TestSensitive(t *testing.T) {
	filter := sensitive.New()
	filter.LoadWordDict("./sensitive.txt")
	result := filter.Replace("赠送礼品", '*')
	fmt.Println(result)
}

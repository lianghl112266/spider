package fetcher

import (
	"fmt"
	"testing"
)

// It is very convenient to test whether your agent is available
func TestFetch(t *testing.T) {
	//if context, err := Fetch(`https://www.google.com`); err == nil {
	//	fmt.Printf("%s", context)
	//}
	if context, err := Fetch(`https://www.tianqi24.com/xiqing/history12.html`); err == nil {
		fmt.Printf("%s", context)
	}

}

package fetcher

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	//if context, err := Fetch(`https://www.google.com`); err == nil {
	//	fmt.Printf("%s", context)
	//}
	if context, err := Fetch(`https://www.tianqi24.com/xiqing/history12.html`); err == nil {
		fmt.Printf("%s", context)
	}

}

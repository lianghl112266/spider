package persist

import (
	"spider/src/spider/module"
	"testing"
)

func TestSave(t *testing.T) {
	err := Save(module.Profile{"1", "tom"})
	if err != nil {
		return
	}
}

package commons_test

import (
	"testing"

	"github.com/arturoeanton/Yuenyeung/commons"
)

func TestContains(t *testing.T) {

	if !commons.Contains([]string{"hola", "pepe"}, "hola") {
		t.Errorf("commons.Contains error")
	}

	if commons.Contains([]string{"hola", "pepe"}, "juan") {
		t.Errorf("commons.Contains error")
	}

}

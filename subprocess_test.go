package process

import (
	"bytes"
	"errors"
	"reflect"
	"testing"
)

func assertError(t *testing.T, err error) {
	if err != nil {
		panic(err)
	}
}

func TestNotFound(t *testing.T) {
	for _, testFunc := range []interface{}{
		NewProcess,
		NewProcessAllStdout,
	} {
		if err := testNotFound(testFunc); err != nil {
			t.Log(err)
			t.Fail()
		}
	}
}

func testNotFound(testFunc interface{}) error {
	result := reflect.ValueOf(testFunc).Call([]reflect.Value{
		reflect.ValueOf("this_is_random_command"),
	})
	verr := result[len(result)-1]
	if verr.Type().String() != "error" {
		return errors.New("not an error")
	}
	return nil
}

func TestSubprocessReadStdout(t *testing.T) {
	data, err := NewProcessAllStdout("echo", "123")
	assertError(t, err)
	t.Log(data)
	if !bytes.Equal(data, []byte("123\n")) {
		t.Log("data in equal")
		t.Fail()
	}
}

package optionparser

import (
	"testing"
)

func TestOption(t *testing.T) {
	f := Init()
	optionCalled := false
	f.HandleOption("test", func(args []string) {
		optionCalled = true
	})
	f.Parse([]string{"test"})
	if !optionCalled {
		t.Error("Expected option to be called")
	}
}

func TestDefaultHandler(t *testing.T) {
	f := Init()
	defaultHandler := false
	f.HandleOption("notused", func(args []string) {
		t.Error("Option should of not been called")
	})
	f.SetDefaultHandler(func() {
		defaultHandler = true
	})
	f.Parse([]string{"notoption"})
	if !defaultHandler {
		t.Error("Expected default handler to be called")
	}
}

func TestDefaultHandlerWithNoOptions(t *testing.T) {
	f := Init()
	defaultHandler := false
	f.HandleOption("notused", func(args []string) {
		t.Error("Option should of not been called")
	})
	f.SetDefaultHandler(func() {
		defaultHandler = true
	})
	f.Parse([]string{})
	if !defaultHandler {
		t.Error("Expected default handler to be called")
	}
}

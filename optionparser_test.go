package optionparser

import (
	"testing"
)

func TestOption(t *testing.T) {
	f := Init()
	optionCalled := false
	f.HandleOption("test", func(args []string, v ...interface{}) {
		optionCalled = true
	})
	f.Parse([]string{"test"})
	if !optionCalled {
		t.Error("Expected option to be called")
	}
}

func TestOptionWithParameter(t *testing.T) {
	f := Init()
	optionCalled := false
	f.HandleOptionWithParameters("test", func(args []string, v ...interface{}) {
		if v[0].(int) == 1 {
			optionCalled = true
		}
	}, 1)
	f.Parse([]string{"test"})
	if !optionCalled {
		t.Error("Expected option to be called")
	}
}

func TestOptionWithMultipleParameter(t *testing.T) {
	f := Init()
	optionCalled := false
	f.HandleOptionWithParameters("test", func(args []string, v ...interface{}) {
		if len(v) == 2 {
			optionCalled = true
		}
	}, 1, 2)
	f.Parse([]string{"test"})
	if !optionCalled {
		t.Error("Expected option to be called")
	}
}

func TestDefaultHandler(t *testing.T) {
	f := Init()
	defaultHandler := false
	f.HandleOption("notused", func(args []string, v ...interface{}) {
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
	f.HandleOption("notused", func(args []string, v ...interface{}) {
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

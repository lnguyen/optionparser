package optionparser

import (
)

type Option struct {
  Pattern string
  Callback func(args []string)
}

type Flags struct {
  Options []*Option
  DefaultHandler func()
}

type Callback func([]string)


func Init() *Flags {
  return &Flags{Options: make([]*Option, 0)}
}

func (f *Flags) HandleOption(pattern string, callback Callback) {
  f.Options = append(f.Options, &Option{Pattern: pattern, Callback: callback})
}

func (f *Flags) SetDefaultHandler(handler func()) {
  f.DefaultHandler = handler
}

func (f *Flags) Parse(args []string) {
  if len(args) == 0 {
    return
  }
  for _, option := range f.Options {
    if option.Callback != nil {
      if option.Pattern == args[0] {
        option.Callback(args[1:])
        return
      }
    }
  }
  if f.DefaultHandler != nil {
    f.DefaultHandler()
  }

}



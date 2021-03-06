package plugin

import (
	"strconv"

	"github.com/einride/protoc-gen-typescript-http/internal/codegen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type enumGenerator struct {
	enum protoreflect.EnumDescriptor
}

func (e enumGenerator) Type() string {
	return string(e.enum.Name())
}

func (e enumGenerator) Generate(f *codegen.File) {
	commentGenerator{descriptor: e.enum}.generateLeading(f, 0)
	f.P("export type ", e.Type(), " = ")
	rangeEnumValues(e.enum, func(value protoreflect.EnumValueDescriptor) {
		commentGenerator{descriptor: value}.generateLeading(f, 1)
		f.P(t(1), "| ", strconv.Quote(string(value.Name())))
	})
}
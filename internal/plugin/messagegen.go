package plugin

import (
	"go.einride.tech/protoc-gen-typescript-http/internal/codegen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type messageGenerator struct {
	pkg     protoreflect.FullName
	message protoreflect.MessageDescriptor
}

func (m messageGenerator) Generate(f *codegen.File, params map[string]string) {
	commentGenerator{descriptor: m.message}.generateLeading(f, 0)
	f.P("export type ", scopedDescriptorTypeName(m.pkg, m.message), " = {")
	rangeFields(m.message, func(field protoreflect.FieldDescriptor) {
		commentGenerator{descriptor: field}.generateLeading(f, 1)
		fieldType := typeFromField(m.pkg, field)
		isRepeated := field.IsList() || field.IsMap()
		if isRepeated || field.ContainingOneof() != nil || field.HasOptionalKeyword() {
			f.P(t(1), field.Name(), "?: ", fieldType.Reference(), ";")
		} else {
			f.P(t(1), field.Name(), ": ", fieldType.Reference(), ";")
		}
	})

	f.P("};")
	f.P()
}

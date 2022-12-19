package plugin

import (
	"go.einride.tech/protoc-gen-typescript-http/internal/codegen"
	"google.golang.org/genproto/googleapis/api/annotations"
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
		useRequiredBehavior := params["use_required_behavior"] == "true"
		isRequired := false
		if useRequiredBehavior {
			for _, behavior := range getFieldBehaviors(field) {
				if behavior == annotations.FieldBehavior_REQUIRED {
					isRequired = true
					break
				}
			}
		}
		if isRequired || !useRequiredBehavior && field.ContainingOneof() == nil && !field.HasOptionalKeyword() {
			f.P(t(1), field.Name(), ": ", fieldType.Reference(), " | undefined;")
		} else {
			f.P(t(1), field.Name(), "?: ", fieldType.Reference(), ";")
		}
	})

	f.P("};")
	f.P()
}

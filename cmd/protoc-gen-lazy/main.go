package main

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"unicode"
)

func main() {
	opts := protogen.Options{}
	opts.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}

			file := gen.NewGeneratedFile(f.GeneratedFilenamePrefix+"_lazy.pb.go", f.GoImportPath)

			file.P("package ", f.GoPackageName)

			for _, m := range f.Messages {
				messageName := "Lazy" + file.QualifiedGoIdent(m.GoIdent)
				lowerMessageName := lower(messageName)
				cursorOffsetsVar := lowerMessageName + "CursorsOffsets"
				messageVar := string(lowerMessageName[0])

				//file.Import("github.com/ClaudioAlbertin/go-fast-protobuf/pkg/parser")
				//file.Import("sync/atomic")
				file.P("type ", messageName, " struct {")
				file.P("	payload []byte")
				file.P("	cursors parser.Cursors")
				file.P("	currentIndex uint32")
				file.P("	message ", file.QualifiedGoIdent(m.GoIdent))
				file.P("}")
				file.P()

				for _, field := range m.Fields {
					returnType, _ := fieldGoType(file, field)
					file.P("func (", messageVar, " *", messageName, ") Get", field.GoName, "() ", returnType, " {")
					file.P("	const fieldNum = ", field.Desc.Number())
					file.P("	if ", messageVar, ".message.", field.GoName, " == \"\" {")
					file.P("		if !", messageVar, ".cursors.IsPresent(", cursorOffsetsVar, "[fieldNum]) {")
					file.P("			readUntil, _ := parser.FindCursors(", cursorOffsetsVar, ", ", messageVar, ".payload, ", messageVar, ".cursors, int(atomic.LoadUint32(&", messageVar, ".currentIndex)), fieldNum)")
					file.P("			atomic.StoreUint32(&", messageVar, ".currentIndex, uint32(readUntil))")
					file.P("		}")
					file.P()
					file.P("		start, end := ", messageVar, ".cursors.AtOffset(", cursorOffsetsVar, "[fieldNum])")
					file.P("		", messageVar, ".message.", field.GoName, " = string(", messageVar, ".payload[start:end])")
					file.P("	}")
					file.P()
					file.P("	return ", messageVar, ".message.", field.GoName)
					file.P("}")
					file.P()
				}

				file.P("func (", messageVar, " *", messageName, ") UnmarshalFrom(payload []byte) error {")
				file.P("	", messageVar, ".payload = payload")
				file.P("	", messageVar, ".cursors = parser.NewCursors(", len(m.Fields)*8, ")")
				file.P("	return nil")
				file.P("}")
				file.P()

				file.P("var ", cursorOffsetsVar, " = map[int32]uint32{")
				for i, field := range m.Fields {
					file.P(field.Desc.Number(), ": ", i*4, ",")
				}
				file.P("}")
				file.P()
			}
		}

		return nil
	})
}

func fieldGoType(g *protogen.GeneratedFile, field *protogen.Field) (goType string, pointer bool) {
	if field.Desc.IsWeak() {
		return "struct{}", false
	}

	pointer = field.Desc.HasPresence()
	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		goType = "bool"
	case protoreflect.EnumKind:
		goType = g.QualifiedGoIdent(field.Enum.GoIdent)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		goType = "int32"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		goType = "uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		goType = "int64"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		goType = "uint64"
	case protoreflect.FloatKind:
		goType = "float32"
	case protoreflect.DoubleKind:
		goType = "float64"
	case protoreflect.StringKind:
		goType = "string"
	case protoreflect.BytesKind:
		goType = "[]byte"
		pointer = false // rely on nullability of slices for presence
	case protoreflect.MessageKind, protoreflect.GroupKind:
		goType = "*" + g.QualifiedGoIdent(field.Message.GoIdent)
		pointer = false // pointer captured as part of the type
	}
	switch {
	case field.Desc.IsList():
		return "[]" + goType, false
	case field.Desc.IsMap():
		keyType, _ := fieldGoType(g, field.Message.Fields[0])
		valType, _ := fieldGoType(g, field.Message.Fields[1])
		return fmt.Sprintf("map[%v]%v", keyType, valType), false
	}
	return goType, pointer
}

func lower(s string) string {
	if len(s) == 0 {
		return s
	}

	r := []rune(s)
	r[0] = unicode.ToLower(r[0])

	return string(r)
}

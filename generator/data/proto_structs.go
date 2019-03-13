package data

import (
	"log"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/yoozoo/protocli/util"
)

// define file/message/field structs to be used in language generators
// wrap protoc-gen-go/descriptor to provide helper methods

// GenerateReq is the code-gen request struct passed to generators
type GenerateReq struct {
	Files      map[string]*ProtoFile
	PackageMap map[string]*ProtoFile
	MessageMap map[string]*ProtoMessage
	EnumMap    map[string]*ProtoEnum
}

var _req *GenerateReq

func Setup(request *plugin.CodeGeneratorRequest) {
	_req = NewGenerateReq(request)
}

func NewGenerateReq(request *plugin.CodeGeneratorRequest) *GenerateReq {
	result := &GenerateReq{}
	result.Files = make(map[string]*ProtoFile)
	result.PackageMap = make(map[string]*ProtoFile)
	result.MessageMap = make(map[string]*ProtoMessage)
	result.EnumMap = make(map[string]*ProtoEnum)

	for _, file := range request.ProtoFile {
		pf := NewProtoFile(file)
		pkg := file.GetPackage()
		result.PackageMap[pkg] = pf
		result.Files[file.GetName()] = pf

		if pkg != "" {
			pkg = pkg + "."
		}

		if util.IsStrInSlice(file.GetName(), request.FileToGenerate) {
			pf.IsFileToGenerate = true
		}

		for name, m := range pf.Messages {
			result.MessageMap[pkg+name] = m
		}

		for name, e := range pf.Enums {
			result.EnumMap[pkg+name] = e
		}
	}

	return result
}

func GetProtoFile(filename string) (file *ProtoFile) {
	file = _req.Files[filename]

	if file == nil {
		log.Println("proto file not found: " + filename)
	}
	return
}

func FlattenLocalPackage(msg *MessageData) {
	_, p := GetMessageProtoAndFile(msg.Name)
	if p == nil || p.IsFileToGenerate {
		msg.Name = msg.Name[strings.LastIndex(msg.Name, ".")+1:]
	}

	for _, f := range msg.Fields {
		_, p = GetMessageProtoAndFile(f.DataType)
		if p == nil || p.IsFileToGenerate {
			f.DataType = f.DataType[strings.LastIndex(f.DataType, ".")+1:]
		}
	}
}

func GetMessageProtoAndFile(name string) (msg *ProtoMessage, file *ProtoFile) {
	var pkg string

	msg = _req.MessageMap[name]
	if msg == nil {
		if !util.IsStrInSlice(name, []string{"string", "int", "int64", "bool"}) {
			if _, ok := _req.EnumMap[name]; !ok {
				log.Println("msg not found: " + name)
			}
		}
	}

	file = GetFileFromPackageWithName(name)

	if file == nil {
		log.Println("pkg not found: " + pkg)
	}
	return
}

func GetEnumProtoAndFile(name string) (e *ProtoEnum, file *ProtoFile) {
	var pkg string

	e = _req.EnumMap[name]
	if e == nil {
		return
	}

	file = GetFileFromPackageWithName(name)

	if file == nil {
		log.Println("pkg not found: " + pkg)
	}
	return
}

func GetFileFromPackageWithName(name string) (file *ProtoFile) {
	slices := strings.Split(name, ".")
	temp := ""
	for _, s := range slices {
		temp = temp + "." + s
		if file, found := _req.PackageMap[temp[1:]]; found {
			return file
		}
	}
	return &ProtoFile{}
}

// ProtoFile is a thin wrapper around descriptor.FileDescriptorProto
type ProtoFile struct {
	IsFileToGenerate bool
	Proto            *descriptor.FileDescriptorProto
	Options          map[string]*ProtoOption
	Enums            map[string]*ProtoEnum
	Messages         map[string]*ProtoMessage
	Services         map[string]*ProtoService
}

// NewProtoFile create ProtoFile from descriptor.FileDescriptorProto
func NewProtoFile(proto *descriptor.FileDescriptorProto) *ProtoFile {
	p := &ProtoFile{
		Proto: proto,
	}

	p.Messages = make(map[string]*ProtoMessage)

	initMessages(p, "", proto.GetMessageType())

	p.Services = make(map[string]*ProtoService)
	for _, svr := range proto.Service {
		p.Services[svr.GetName()] = NewProtoService(svr)
	}

	p.Enums = make(map[string]*ProtoEnum)
	for _, obj := range proto.EnumType {
		p.Enums[obj.GetName()] = NewProtoEnum(obj)
	}

	return p
}

func initMessages(p *ProtoFile, namePrefix string, msgs []*descriptor.DescriptorProto) {
	for _, msg := range msgs {
		name := namePrefix + msg.GetName()
		p.Messages[name] = NewProtoMessage(msg)
		nested := msg.GetNestedType()
		if nested != nil {
			initMessages(p, name+".", nested)
		}
	}
}

// ProtoOption is a thin wrapper around descriptor.OptionDescriptorProto
type ProtoOption struct {
}

// ProtoEnum is a thin wrapper around descriptor.EnumDescriptorProto
type ProtoEnum struct {
	Proto *descriptor.EnumDescriptorProto
}

// NewProtoEnum create ProtoEnum from descriptor.EnumDescriptorProto
func NewProtoEnum(proto *descriptor.EnumDescriptorProto) *ProtoEnum {
	return &ProtoEnum{
		Proto: proto,
	}
}

// ProtoMessage is a thin wrapper around descriptor.DescriptorProto (Message descriptor)
type ProtoMessage struct {
	Proto   *descriptor.DescriptorProto
	Options map[string]*ProtoOption
	Fields  map[string]*ProtoField
}

// NewProtoMessage create ProtoMessage from descriptor.DescriptorProto
func NewProtoMessage(proto *descriptor.DescriptorProto) *ProtoMessage {
	return &ProtoMessage{
		Proto: proto,
	}
}

// ProtoField is a thin wrapper around descriptor.FieldDescriptorProto
type ProtoField struct {
	Proto   *descriptor.FieldDescriptorProto
	Options map[string]*ProtoOption
}

// NewProtoField create ProtoField from descriptor.FieldDescriptorProto
func NewProtoField(proto *descriptor.FieldDescriptorProto) *ProtoField {
	return &ProtoField{
		Proto: proto,
	}
}

// ProtoMethod is a thin wrapper around descriptor.MethodDescriptorProto
type ProtoMethod struct {
	Proto   *descriptor.MethodDescriptorProto
	Options map[string]*ProtoOption
}

// NewProtoMethod create ProtoMethod from descriptor.MethodDescriptorProto
func NewProtoMethod(proto *descriptor.MethodDescriptorProto) *ProtoMethod {
	return &ProtoMethod{
		Proto: proto,
	}
}

// ProtoService is a thin wrapper around descriptor.ServiceDescriptorProto
type ProtoService struct {
	Proto   *descriptor.ServiceDescriptorProto
	Options map[string]*ProtoOption
	Methods map[string]*ProtoMethod
}

// NewProtoService create ProtoService from descriptor.ServiceDescriptorProto
func NewProtoService(proto *descriptor.ServiceDescriptorProto) *ProtoService {
	return &ProtoService{
		Proto: proto,
	}
}

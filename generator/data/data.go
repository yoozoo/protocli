package data

import (
	"fmt"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

const (
	//IntFieldType datatype string for interge, it is assumed to be of signed and at least 64 bit
	IntFieldType = "int"
	//Int32FieldType datatype string for interge 32 bit
	Int32FieldType = "int32"
	//Int64FieldType datatype string for interge 64 bit
	Int64FieldType = "int64"
	// BooleanFieldType datatype string for boolean
	BooleanFieldType = "bool"
	// StringFieldType datatype string for string
	StringFieldType = "string"
	// DoubleFieldType datatype string for floating point, it is assumed to be at least of ieee double precision
	DoubleFieldType = "double"

	// PathSeparator the path seperator used to form the full key (ie, key/sub_key )
	PathSeparator = "/"
	// FieldRepeatedLabel is the label for repeated data type
	FieldRepeatedLabel = "LABEL_REPEATED"
	// JavaPackageOption is Java package option constant
	JavaPackageOption = "javaPackageOption"

	// ServiceAuthOption is service auth option
	ServiceAuthOption = 51009
	// ServiceCommonErrorOption is service common_error option
	ServiceCommonErrorOption = 51008
	// ServiceTypeMethodOption is service method option
	ServiceTypeMethodOption = 51006
	// ErrorTypeMethodOption is error return type option
	ErrorTypeMethodOption = 51007
	// FormatFieldOption is the field type validation field option
	FormatFieldOption = 51002
	// RequiredFieldOption is the required type validation field option
	RequiredFieldOption = 51003

	// ComErrMsgName  is common error message name
	ComErrMsgName = "CommonError"

	// path numbers in FileDescriptorProto (describe proto file)
	MessageCommentPath = 4
	EnumCommentPath    = 5
	ServiceCommentPath = 6

	// path numbers in DescriptorProto (describe messages)
	MessageFieldCommentPath  = 2 // field
	MessageNestedCommentPath = 3 // nested
	MessageEnumCommentPath   = 4 // enum

	// path numbers in EnumDescriptorProto (describe enum)
	EnumFieldCommentPath = 2 // field

	// path numbers in ServiceDescriptorProto (describe service)
	ServiceMethodCommentPath = 2
)

//ServiceOptions is the map of field number and field name in service options
var ServiceOptions = map[int32]OptionInfo{
	ServiceCommonErrorOption: OptionInfo{"common_error", (*string)(nil), StringFieldType},
	ServiceAuthOption:        OptionInfo{"auth", (*bool)(nil), BooleanFieldType},
}

// MethodOptions is the map of field number and field name in method options
var MethodOptions = map[int32]OptionInfo{
	ServiceTypeMethodOption: OptionInfo{"service_method", (*string)(nil), StringFieldType},
	ErrorTypeMethodOption:   OptionInfo{"error", (*string)(nil), StringFieldType},
}

// FieldOptions is the map of field number and field name in field options
var FieldOptions = map[int32]OptionInfo{
	FormatFieldOption:   OptionInfo{"val_format", (*string)(nil), StringFieldType},
	RequiredFieldOption: OptionInfo{"val_required", (*bool)(nil), BooleanFieldType},
}

type OptionInfo struct {
	Name       string
	DefaultNil interface{}
	Type       string
}

// EnumField a enum entry for enum datatype
type EnumField struct {
	Name    string // enum entry name
	Value   int32  // enum entry value
	Comment string
}

// EnumData a structure to represent a enum datatype
type EnumData struct {
	Name    string // enum type name
	Comment string
	Fields  []EnumField // enum entries
}

// MessageField a field for the defined message.
type MessageField struct {
	Name     string // message variable name
	DataType string // message variable type
	Key      string // coresponding key name for the variable, default is the same as variable name
	Label    string
	Comment  string
	Options  OptionMap
}

// MessageData a structure to represent a message datatype
type MessageData struct {
	File    string // file where this message is defined
	Name    string // name of the message (class, struct)
	Comment string
	Fields  []*MessageField // message members
}

type Method struct {
	Name       string
	InputType  string
	OutputType string
	HttpMtd    string
	URI        string
	Comment    string
	Options    OptionMap // service method option (default is GET and POST)
}

type ServiceData struct {
	Name            string
	Comment         string
	Methods         []*Method
	Options         OptionMap
	CommonErrorType string
	Service         *descriptor.ServiceDescriptorProto
}

// OptionMap is a structure represents the option declared in a proto file
type OptionMap map[string]string

// CodeGenerator define the interface for code generator
type CodeGenerator interface {
	Init(request *plugin.CodeGeneratorRequest)
	Gen(applicationName string, packageName string, services []*ServiceData, messages []*MessageData, enums []*EnumData, options OptionMap) (map[string]string, error)
	GetLang() string
}

var outputMap = make(map[string]CodeGenerator)

// RegisterCodeGenerator register CodeGenerator with given name
func RegisterCodeGenerator(name string, gen CodeGenerator) {
	if _, ok := outputMap[name]; ok {
		panic(fmt.Errorf("%s code generator already exist", name))
	}

	outputMap[name] = gen
}

// GetCodeGenerator register CodeGenerator with given name
func GetCodeGenerator(name string) CodeGenerator {
	if gen, ok := outputMap[name]; ok {
		return gen
	}

	panic(fmt.Errorf("%s code generator not found", name))
}

// HasCodeGenerator returns if CodeGenerator has been registered
func HasCodeGenerator(name string) bool {
	_, ok := outputMap[name]
	return ok
}

// GetCodeGeneratorMap returns code generator map
func GetCodeGeneratorMap() map[string]CodeGenerator {
	return outputMap
}

// CommentMap map comment to message/service/field
// path => comment
type CommentMap map[string]string

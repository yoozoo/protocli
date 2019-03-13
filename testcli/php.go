package main

import (
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/yoozoo/protocli/generator/data"
	"github.com/yoozoo/protocli/util"
)

type yii2Gen struct {
	result     map[string]string
	enums      []*data.EnumData
	ModuleName string
	NameSpace  string
	bizErrors  []string
	comError   *data.MessageData
}

func (g *yii2Gen) Init(request *plugin.CodeGeneratorRequest) {

}

func (g *yii2Gen) Gen(applicationName string, packageName string, services []*data.ServiceData, messages []*data.MessageData, enums []*data.EnumData, options data.OptionMap) (result map[string]string, err error) {
	return map[string]string{"test": "test"}, nil
}

func init() {
	data.RegisterCodeGenerator("go", &yii2Gen{})
	data.RegisterCodeGenerator("php", &yii2Gen{})
	util.RegisterIncludeFile("protoapi_common.proto", `syntax = "proto3";

	import "google/protobuf/descriptor.proto";

	extend google.protobuf.MethodOptions {
	  string service_method = 51006;
	  string error = 51007;
	}

	extend google.protobuf.ServiceOptions {
	  string common_error = 51008;
	  bool auth = 51009;
	}

	extend google.protobuf.FieldOptions {
	  string val_format = 51002;
	  bool val_required = 51003;
	  int32 min = 51004;
	  int32 max = 51005;
	}

	message CommonError {
	  GenericError genericError = 1;
	  AuthError authError = 2;
	  ValidateError validateError = 3;
	  BindError bindError = 4;
	}

	message GenericError { string message = 1; }

	message AuthError { string message = 1; }

	message BindError { string message = 1; }

	message ValidateError { repeated FieldError errors = 1; }

	message FieldError {
	  string fieldName = 1;
	  ValidateErrorType errorType = 2;
	}

	enum ValidateErrorType {
	  INVALID_EMAIL = 0;
	  FIELD_REQUIRED = 1;
	}
	`)
}

package main

import (
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/yoozoo/protocli/generator/data"
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
	return
}

func init() {
	data.RegisterCodeGenerator("go", &yii2Gen{})
	data.RegisterCodeGenerator("php", &yii2Gen{})
}

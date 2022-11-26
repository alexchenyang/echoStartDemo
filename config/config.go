package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var ConfigInstance ConfStrut

type ConfStrut struct {
	ProjectName string `yaml:"project_name"`
}

// 初始化config包，解析project_conf.yaml, 获取conf全局变量，用于各处
func Init() error {
	fileName := "./conf/project_conf.yaml"
	err := ParseYamlFile(fileName, &ConfigInstance)
	if err != nil {
		return err
	}
	fmt.Printf("parse conf yaml sucess. project name: %v", ConfigInstance.ProjectName)
	return nil
}

func ParseYamlFile(fileName string, resp interface{}) error {
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(fileContent, resp)
}

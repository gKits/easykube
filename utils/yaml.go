package utils

import (
	"io/ioutil"
	"os"

	"k8s.io/apimachinery/pkg/util/yaml"
)

func UnmarshalYaml(yamlPath string, v interface{}) error {
    file, err := os.Open(yamlPath)
    if err != nil {
        return err
    }

    bytes, err := ioutil.ReadAll(file)
    if err != nil {
        return err
    }

    if err = yaml.Unmarshal(bytes, &v); err != nil {
        return err
    }

    return nil
}

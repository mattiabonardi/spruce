package entity

import (
	"bufio"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

type EntityDefinitionManager struct{}

func (h *EntityDefinitionManager) GetDefinition(entityClass string) (EntityDefinition, error) {
	entityDefinition := EntityDefinition{}
	// open file
	filePath := entityClass + ".yaml"
	file, err := os.Open(h.rootDir() + "/resources/entities/definitions/" + filePath)
	if err != nil {
		return entityDefinition, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var yamlContent string
	for scanner.Scan() {
		yamlContent += scanner.Text() + "\n"
	}
	err = yaml.Unmarshal([]byte(yamlContent), &entityDefinition)
	if err != nil {
		return entityDefinition, err
	}
	return entityDefinition, nil
}

func (h *EntityDefinitionManager) rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

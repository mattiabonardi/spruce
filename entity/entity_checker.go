package entity

import (
	"errors"
	"reflect"

	"github.com/mattiabonardi/spruce/models"
	"github.com/mattiabonardi/spruce/utils"
)

type EntityChecker struct{}

func (h EntityChecker) Check(entityDefinition models.EntityDefinition, entity models.Entity) error {
	// check the type of entity attributes
	for name, attribute := range entity.Attributes {
		// get the type in definition
		attributeType := entityDefinition.Attributes[name].Type
		if attributeType == "" {
			return errors.New("Missing type attribute for " + name)
		}
		err := h.checkAttributeType(name, attributeType, attribute)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h EntityChecker) checkAttributeType(attributeName string, attributeTypeDefinition models.Type, attributeValue interface{}) error {
	valueType := reflect.TypeOf(attributeValue)
	switch attributeTypeDefinition {
	case models.String:
		{
			if valueType.Kind() != reflect.String {
				return errors.New(attributeName + " isn't a String")
			} else {
				return nil
			}
		}
	case models.Integer:
		{
			if valueType.Kind() != reflect.Int && valueType.Kind() != reflect.Int8 && valueType.Kind() != reflect.Int16 && valueType.Kind() != reflect.Int32 && valueType.Kind() != reflect.Int64 && valueType.Kind() != reflect.Uint && valueType.Kind() != reflect.Uint8 && valueType.Kind() != reflect.Uint16 && valueType.Kind() != reflect.Uint32 && valueType.Kind() != reflect.Uint64 {
				return errors.New(attributeName + " isn't an Integer")
			} else {
				return nil
			}
		}
	case models.Decimal:
		{
			if valueType.Kind() != reflect.Float32 && valueType.Kind() != reflect.Float64 {
				return errors.New(attributeName + " isn't a Decimal")
			} else {
				return nil
			}
		}
	case models.Boolean:
		{
			if valueType.Kind() != reflect.Bool {
				return errors.New(attributeName + " isn't a Boolean")
			} else {
				return nil
			}
		}
	case models.ObjectId:
		{
			regexManager := utils.RegexManager{}
			if valueType.Kind() == reflect.String && !regexManager.Match(attributeValue.(string), utils.EXADECIMAL_PATTERN) {
				// change value
				return nil
			} else {
				return errors.New(attributeName + " isn't an ObjectId")
			}
		}
	default:
		return errors.New("type not supported")
	}
}

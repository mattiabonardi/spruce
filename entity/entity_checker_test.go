package entity

import (
	"testing"

	"github.com/mattiabonardi/spruce/models"
)

func TestCheckAttributeType(t *testing.T) {
	entityChecker := EntityChecker{}
	// default type
	err := entityChecker.checkAttributeType("notExstenceType", "NotExistenceType", "aaaa")
	if err == nil {
		t.Fatal("Default type check failed")
	}

	// String test
	err = entityChecker.checkAttributeType("tstString", models.String, "aaaa")
	if err != nil {
		t.Fatal("String check failed")
	}
	err = entityChecker.checkAttributeType("tstString", models.String, 1)
	if err == nil {
		t.Fatal("String check failed")
	}

	// Integer test
	err = entityChecker.checkAttributeType("tstInteger", models.Integer, 1)
	if err != nil {
		t.Fatal("Integer check failed")
	}
	err = entityChecker.checkAttributeType("tstInteger", models.Integer, "aaa")
	if err == nil {
		t.Fatal("Integer check failed")
	}

	// Decimal test
	err = entityChecker.checkAttributeType("tstDecimal", models.Decimal, 1.0)
	if err != nil {
		t.Fatal("Decimal check failed")
	}
	err = entityChecker.checkAttributeType("tstDecimal", models.Decimal, "aaa")
	if err == nil {
		t.Fatal("Decimal check failed")
	}
}

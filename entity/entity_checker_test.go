package entity

import (
	"testing"
)

func TestCheckAttributeType(t *testing.T) {
	entityChecker := EntityChecker{}
	// default type
	err := entityChecker.checkAttributeType("notExstenceType", "NotExistenceType", "aaaa")
	if err == nil {
		t.Fatal("Default type check failed")
	}

	// String test
	err = entityChecker.checkAttributeType("tstString", String, "aaaa")
	if err != nil {
		t.Fatal("String check failed")
	}
	err = entityChecker.checkAttributeType("tstString", String, 1)
	if err == nil {
		t.Fatal("String check failed")
	}

	// Integer test
	err = entityChecker.checkAttributeType("tstInteger", Integer, 1)
	if err != nil {
		t.Fatal("Integer check failed")
	}
	err = entityChecker.checkAttributeType("tstInteger", Integer, "aaa")
	if err == nil {
		t.Fatal("Integer check failed")
	}

	// Decimal test
	err = entityChecker.checkAttributeType("tstDecimal", Decimal, 1.0)
	if err != nil {
		t.Fatal("Decimal check failed")
	}
	err = entityChecker.checkAttributeType("tstDecimal", Decimal, "aaa")
	if err == nil {
		t.Fatal("Decimal check failed")
	}

	// Boolean test
	err = entityChecker.checkAttributeType("tstBoolean", Boolean, true)
	if err != nil {
		t.Fatal("Boolean check failed")
	}
	err = entityChecker.checkAttributeType("tstDecimal", Boolean, "aaa")
	if err == nil {
		t.Fatal("Boolean check failed")
	}

	// ObjectId test
	err = entityChecker.checkAttributeType("tstObjectid", ObjectId, "609c6aae7e8d9a2c6d36e82d")
	if err != nil {
		t.Fatal("ObjectId check failed")
	}
	err = entityChecker.checkAttributeType("tstObjectId", ObjectId, "aaa")
	if err == nil {
		t.Fatal("ObjectId check failed")
	}
}

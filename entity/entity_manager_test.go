package entity

import (
	"fmt"
	"testing"
)

func TestGetAll(t *testing.T) {
	entityManager := EntityManager{}
	entities, err := entityManager.GetAll("gender", EntityContext{})

	if err != nil {
		t.Fatal(err)
	}

	if entities[0].Attributes["_id"].Type != "String" {
		t.Fatal("Expected: String got: " + entities[0].Attributes["_id"].Type)
	}

	if entities[0].Attributes["_id"].Value != "1" {
		t.Fatal("Expected: 1 got: " + fmt.Sprintf("%v", entities[0].Attributes["_id"].Value))
	}

	if entities[0].Attributes["description"].Value != "male" {
		t.Fatal("Expected: male got: " + fmt.Sprint(entities[0].Attributes["description"].Value))
	}

	if len(entities) != 2 {
		t.Fatal("Expected: 2 got: " + fmt.Sprint(entities))
	}
}

func TestGetById(t *testing.T) {
	entityManager := EntityManager{}
	entity, err := entityManager.GetById("gender", EntityContext{}, "2")

	if err != nil {
		t.Fatal(err)
	}

	if entity.Attributes["_id"].Type != "String" {
		t.Fatal("Expected: String got: " + entity.Attributes["id"].Type)
	}

	if entity.Attributes["_id"].Value != "2" {
		t.Fatal("Expected: 2 got: " + fmt.Sprintf("%v", entity.Attributes["_id"].Value))
	}

	if entity.Attributes["description"].Value != "female" {
		t.Fatal("Expected: female got: " + fmt.Sprint(entity.Attributes["description"].Value))
	}
}

func TestCreateAndDelete(t *testing.T) {
	entityManager := EntityManager{}
	entity := Entity{}
	entity.Class = "gender"
	attributes := make(map[string]Attribute)
	attributes["_id"] = Attribute{
		Type:  String,
		Value: "3",
	}
	attributes["description"] = Attribute{
		Type:  String,
		Value: "Other",
	}
	entity.Attributes = attributes
	_, err := entityManager.Create("gender", EntityContext{}, entity)
	if err != nil {
		t.Fatal(err)
	}
	_, err = entityManager.Create("gender", EntityContext{}, entity)
	if err == nil {
		t.Fatal("write test failed, tried to write with same id")
	}
	// delete
	err = entityManager.DeleteById("gender", EntityContext{}, "3")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateAndReset(t *testing.T) {
	entityManager := EntityManager{}
	// update
	entity := Entity{}
	entity.Class = "gender"
	attributes := make(map[string]Attribute)
	attributes["_id"] = Attribute{
		Type:  String,
		Value: "2",
	}
	attributes["description"] = Attribute{
		Type:  String,
		Value: "female updated",
	}
	entity.Attributes = attributes
	err := entityManager.Update("gender", EntityContext{}, entity)
	if err != nil {
		t.Fatal(err)
	}
	// reset
	entity.Attributes["description"] = Attribute{
		Type:  String,
		Value: "female",
	}
	err = entityManager.Update("gender", EntityContext{}, entity)
	if err != nil {
		t.Fatal(err)
	}
}

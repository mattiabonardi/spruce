package entity

import (
	"fmt"
	"testing"

	"github.com/mattiabonardi/spruce/models"
)

func TestGetAll(t *testing.T) {
	entityManager := EntityManager{}
	entities := entityManager.GetAll("gender", models.EntityContext{})

	if entities[0].Attributes["_id"].Type != "String" {
		t.Fatal("Expected: String got: " + entities[0].Attributes["id"].Type)
	}

	if entities[0].Attributes["_id"].Value != "1" {
		t.Fatal("Expected: 1 got: " + fmt.Sprintf("%v", entities[0].Attributes["id"].Value))
	}

	if entities[0].Attributes["description"].Value != "male" {
		t.Fatal("Expected: male got: " + fmt.Sprint(entities[0].Attributes["description"].Value))
	}

	if len(entities) != 2 {
		t.Fatal("Expected: 2 got: " + fmt.Sprint(entities))
	}
}

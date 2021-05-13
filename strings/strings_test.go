package strings

import "testing"

func TestExistsInSlice_should_find_a_string(t *testing.T) {
	slice := []string{"a", "b", "c"}
	element := slice[0]

	if exists := ExistsInSlice(element, slice); !exists {
		t.Errorf("should find %s", slice[0])
	}
}

func TestExistsInSlice_should_not_find_a_string(t *testing.T) {
	slice := []string{}
	element := "a"

	if exists := ExistsInSlice(element, slice); exists {
		t.Errorf("should not find %s", slice[0])
	}
}

func TestExistsInSlice_should_find_a_struct_instance(t *testing.T) {
	type myStruct struct {
		value string
	}

	slice := []myStruct{
		{value: "a"},
		{value: "b"},
		{value: "c"},
	}

	element := slice[0]

	if exists := ExistsInSlice(element, slice); !exists {
		t.Errorf("should find %v", slice[0])
	}
}

func TestExistsInSlice_should_not_find_a_struct_instance(t *testing.T) {
	type myStruct struct {
		value string
	}

	slice := []myStruct{
		{value: "a"},
		{value: "b"},
		{value: "c"},
	}

	element := myStruct{value: "d"}

	if exists := ExistsInSlice(element, slice); exists {
		t.Errorf("should not find %v", slice[0])
	}
}

func TestExistsInSlice_should_return_false_to_invalid_slice(t *testing.T) {
	slice := "a"
	element := "a"

	if exists := ExistsInSlice(element, slice); exists {
		t.Error("should return false")
	}
}

func TestExistsInSlice_should_return_false_for_different_types(t *testing.T) {
	type myStruct1 struct {
		value string
	}

	type myStruct2 struct {
		value string
	}

	slice := []myStruct1{
		{value: "a"},
		{value: "b"},
		{value: "c"},
	}

	element := myStruct2{value: "a"}

	if exists := ExistsInSlice(element, slice); exists {
		t.Error("should return false for different types")
	}
}

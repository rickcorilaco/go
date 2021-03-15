package env

import (
	"testing"
)

func TestFromFile(t *testing.T) {
	err := FromFile("example_env.json")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestString(t *testing.T) {
	TestFromFile(t)

	value, err := String("repository.username")
	if err != nil {
		t.Error(err)
		return
	}

	expected := "root"

	if value != expected {
		t.Errorf("expected for '%s' but obtained '%s'", expected, value)
		return
	}
}

func TestMustString(t *testing.T) {
	TestFromFile(t)

	value := MustString("repository.username")
	expected := "root"

	if value != expected {
		t.Errorf("expected for '%s' but obtained '%s'", expected, value)
		return
	}
}

func TestTryString(t *testing.T) {
	TestFromFile(t)

	username := TryString("repository.username")
	expected := "root"

	if username != expected {
		t.Errorf("expected for '%s' but obtained '%s'", expected, username)
		return
	}
}

func TestInt(t *testing.T) {
	TestFromFile(t)

	value, err := Int("repository.port")
	if err != nil {
		t.Error(err)
		return
	}

	expected := int(1234)

	if value != expected {
		t.Errorf("expected for '%d' but obtained '%d'", expected, value)
		return
	}
}

func TestMustInt(t *testing.T) {
	TestFromFile(t)

	value := MustInt("repository.port")

	expected := int(1234)

	if value != expected {
		t.Errorf("expected for '%d' but obtained '%d'", expected, value)
		return
	}
}

func TestTryInt(t *testing.T) {
	TestFromFile(t)

	value := TryInt("repository.port")

	expected := int(1234)

	if value != expected {
		t.Errorf("expected for '%d' but obtained '%d'", expected, value)
		return
	}
}

func TestInt64(t *testing.T) {
	TestFromFile(t)

	value, err := Int64("repository.port")
	if err != nil {
		t.Error(err)
		return
	}

	expected := int64(1234)

	if value != expected {
		t.Errorf("expected for '%d' but obtained '%d'", expected, value)
		return
	}
}

func TestMustInt64(t *testing.T) {
	TestFromFile(t)

	value := MustInt64("repository.port")

	expected := int64(1234)

	if value != expected {
		t.Errorf("expected for '%d' but obtained '%d'", expected, value)
		return
	}
}

func TestTryInt64(t *testing.T) {
	TestFromFile(t)

	value := TryInt64("repository.port")

	expected := int64(1234)

	if value != expected {
		t.Errorf("expected for '%d' but obtained '%d'", expected, value)
		return
	}
}

func TestFloat64(t *testing.T) {
	TestFromFile(t)

	value, err := Float64("payment.min_value")
	if err != nil {
		t.Error(err)
		return
	}

	expected := float64(9.99)

	if value != expected {
		t.Errorf("expected for '%f' but obtained '%f'", expected, value)
		return
	}
}

func TestMustFloat64(t *testing.T) {
	TestFromFile(t)

	value := MustFloat64("payment.min_value")

	expected := float64(9.99)

	if value != expected {
		t.Errorf("expected for '%f' but obtained '%f'", expected, value)
		return
	}
}

func TestTryFloat64(t *testing.T) {
	TestFromFile(t)

	value := TryFloat64("payment.min_value")

	expected := float64(9.99)

	if value != expected {
		t.Errorf("expected for '%f' but obtained '%f'", expected, value)
		return
	}
}

func TestBool(t *testing.T) {
	TestFromFile(t)

	value, err := Bool("repository.log")
	if err != nil {
		t.Error(err)
		return
	}

	expected := bool(true)

	if value != expected {
		t.Errorf("expected for '%t' but obtained '%t'", expected, value)
		return
	}
}

func TestMustBool(t *testing.T) {
	TestFromFile(t)

	value := MustBool("repository.log")

	expected := bool(true)

	if value != expected {
		t.Errorf("expected for '%t' but obtained '%t'", expected, value)
		return
	}
}

func TestTryBool(t *testing.T) {
	TestFromFile(t)

	value := TryBool("repository.log")

	expected := bool(true)

	if value != expected {
		t.Errorf("expected for '%t' but obtained '%t'", expected, value)
		return
	}
}

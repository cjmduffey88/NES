package bus

import "testing"

func TestBus_Read(t *testing.T) {

	memory := bus{}
	result := memory.Read(0x0)
	expected := uint8(0x0)

	if result != expected {
		t.Errorf("Result %q, Expected %q", result, expected)
	}
}

func TestBus_Write(t *testing.T) {

	memory := bus{}

	// Test Writing to Work RAM
	memory.Write(0x0, 0x1)
	result := memory.Read(0x0)
	expected := uint8(0x1)
	if result != expected {
		t.Errorf("Result %q, Expected %q", result, expected)
	}

	// Test Working RAM mirrors
	result = memory.Read(0x0 + 0x800)
	expected = uint8(0x1)
	if result != expected {
		t.Errorf("Result %q, Expected %q", result, expected)
	}

}
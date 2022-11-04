package bus

type bus struct {
	// 0x0000 - 0xFFFF
	memory [0x10000]uint8
}

// Writes data to memory
func (b *bus) Write(address uint16, data uint8) {

	// Write Data
	b.memory[address] = data

	// Write Mirrors
	switch {
	case address < 0x800:
		b.memory[address+0x800] = data
		b.memory[address+0x1000] = data
		b.memory[address+0x1800] = data
	case address < 0x1000:
		b.memory[address-0x800] = data
		b.memory[address+0x800] = data
		b.memory[address+0x1000] = data
	case address < 0x1800:
		b.memory[address-0x1000] = data
		b.memory[address-0x800] = data
		b.memory[address+0x800] = data
	case address < 0x2000:
		b.memory[address-0x1800] = data
		b.memory[address-0x1000] = data
		b.memory[address-0x800] = data
	case address < 0x4000:
		for x := address; x >= 0x2000; x -= 0x8 {
			b.memory[x] = data
		}
		for x := address; x < 0x4000; x += 0x8 {
			b.memory[x] = data
		}
	}

}

// Reads data from memory
func (b *bus) Read(address uint16) uint8 {
	return b.memory[address]
}

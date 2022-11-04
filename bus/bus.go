package bus

type bus struct {
	// 0x0000 - 0xFFFF
	memory [0x10000]uint8
}

// Writes data to memory
func (b *bus) Write(address uint16, data uint8) {

	switch {

	// Work RAM
	case address < 0x800:
		b.memory[address] = data
		b.memory[address+0x800] = data
		b.memory[address+0x1000] = data
		b.memory[address+0x1800] = data
	// Work RAM Mirror
	case address < 0x1000:
		b.memory[address-0x800] = data
		b.memory[address] = data
		b.memory[address+0x800] = data
		b.memory[address+0x1000] = data
	// Work RAM Mirror
	case address < 0x1800:
		b.memory[address-0x1000] = data
		b.memory[address-0x800] = data
		b.memory[address] = data
		b.memory[address+0x800] = data
	// Work RAM Mirror
	case address < 0x2000:
		b.memory[address-0x1800] = data
		b.memory[address-0x1000] = data
		b.memory[address-0x800] = data
		b.memory[address] = data

	}

}

// Reads data from memory
func (b bus) Read(address uint16) uint8 {
	return b.memory[address]
}

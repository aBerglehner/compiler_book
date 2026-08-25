package code

import "testing"

func TestMake(t *testing.T) {
	tests := []struct {
		op       Opcode
		operands []int
		expected []byte
	}{
		// we only show max uint16 that is why there are 2 bytes 2*8bits=16bits
		// first byte in expected are the first 8bits + 8bits
		// 32768 16384 8192 4096 2048 1024 512 256 128 64 32 16 8 4 2 1
		//   1   1      1    1    1    1   1   0   1  1   1  1 1 1 1  1
		{OpConstant, []int{65534}, []byte{byte(OpConstant), 255, 254}},
		{OpConstant, []int{65535}, []byte{byte(OpConstant), 255, 255}},
		{OpConstant, []int{65279}, []byte{byte(OpConstant), 254, 255}},
		{OpConstant, []int{1}, []byte{byte(OpConstant), 0, 1}},
		{OpConstant, []int{257}, []byte{byte(OpConstant), 1, 1}},
		{OpConstant, []int{256}, []byte{byte(OpConstant), 1, 0}},
	}

	for _, tt := range tests {
		instruction := Make(tt.op, tt.operands...)

		if len(instruction) != len(tt.expected) {
			t.Errorf("instruction has wrong length. want=%d, got=%d",
				len(tt.expected), len(instruction))
		}

		for i, b := range tt.expected {
			if instruction[i] != tt.expected[i] {
				t.Errorf("wrong byte at pos %d. want=%d, got=%d",
					i, b, instruction[i])
			}
		}
	}
}

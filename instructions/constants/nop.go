package constants

import "A-Small-JVM/instructions/base"
import "A-Small-JVM/rtda"

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	// do nothing
}

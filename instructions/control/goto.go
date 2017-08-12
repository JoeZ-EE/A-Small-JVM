package control

import "A-Small-JVM/instructions/base"
import "A-Small-JVM/rtda"

type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}

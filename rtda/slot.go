package rtda

import "A-Small-JVM/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}

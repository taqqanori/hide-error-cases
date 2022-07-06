package main

type scope interface {
	getDepth() int
}
type scopeStack struct {
	stack []scope
}

func (ss *scopeStack) push(s scope) {
	ss.stack = append(ss.stack, s)
}

func (ss *scopeStack) popIfDepthMatch(depth int) scope {
	if len(ss.stack) == 0 {
		return nil
	}
	result := (ss.stack)[len(ss.stack)-1]
	if result.getDepth() == depth {
		ss.stack = (ss.stack)[:len(ss.stack)-1]
		return result
	}
	return nil
}

func (ss *scopeStack) peek() scope {
	if len(ss.stack) == 0 {
		return nil
	}
	return (ss.stack)[len(ss.stack)-1]
}

package gr

func rMulti() [][]byte {
	return multiCompile("MULTI")
}

func rExec() [][]byte {
	return multiCompile("EXEC")
}

func rDiscard() [][]byte {
	return multiCompile("DISCARD")
}

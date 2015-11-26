package gr

func rPFAdd(key string, elements ...string) ([][]byte, error) {
	if len(elements) < 1 {
		return nil, NotEnoughParamsErr
	}

	cmds := append([]string{"PFADD", key}, elements...)
	return multiCompile(cmds...), nil
}

func rPFCount(keys ...string) ([][]byte, error) {
	if len(keys) < 1 {
		return nil, NotEnoughParamsErr
	}

	cmds := append([]string{"PFCOUNT"}, keys...)
	return multiCompile(cmds...), nil
}

func rPFMerge(destkey string, sourcekeys ...string) ([][]byte, error) {
	if len(sourcekeys) < 1 {
		return nil, NotEnoughParamsErr
	}

	cmds := append([]string{"PFMERGE", destkey}, sourcekeys...)
	return multiCompile(cmds...), nil
}

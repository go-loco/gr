package gr

type Transaction struct {
	Pipeline
}

func newTransaction(r *Redis) *Transaction {
	p := newPipeline(r)
	p.enqueueStr(rMulti())

	return &Transaction{*p}
}

/////////////
//TRANSACTION
/////////////

func (p *Transaction) Discard() *RespString {
	return p.enqueueStr(rDiscard())
}

func (p *Transaction) Watch() *RespString {
	return p.enqueueStr(rWatch())
}

func (p *Transaction) UnWatch() *RespString {
	return p.enqueueStr(rUnWatch())
}

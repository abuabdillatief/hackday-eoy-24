package services

type Resp struct{
	raw string
}


func (r *Resp) Bytes() []byte {
	if r == nil {
		return nil
	}
	
	return []byte(r.raw)
}
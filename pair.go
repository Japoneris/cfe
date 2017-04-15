//This is the class implementation of keypair
//Voir si la keypair est bien une string.
package cfe



type Pair struct {
    L, R interface{}
}

func (h Pair) GetL() interface{} {return h.L }
func (h Pair) GetR() interface{} {return h.R }
func (h *Pair) Set(l interface{}, r interface{})  {h.L = l; h.R = r }
func (h *Pair) SetL(l interface{})  {h.L = l}
func (h *Pair) SetR(r interface{}) {h.R = r}

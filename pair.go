//This is the class implementation of keypair
//Voir si la keypair est bien une string.
package pair

import 	"fmt"

type Pair struct {
  l, r string
}

func (h Pair) GetL() string {return h.l }
func (h Pair) GetR() string {return h.r }
func (h *Pair) Set(l, r string)  {h.l = l; h.r = r }
func (h *Pair) SetL(l string)  {h.l = l}
func (h *Pair) SetR(r string) {h.r = r}

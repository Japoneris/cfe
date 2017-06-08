// This is the core librairy

package cfe

import (
  "fmt"
  "encoding/binary"
  "math"
  "crypto/rand"
  "math/big"

)

type Cfe struct {
  securityParameter int           // Number of bit
  minPaddingSizeBytes int         // Padding in Bytes
  numBytesInSingleEncryption int  // Max Bytes per encryption
  nmpke *Nmpke                    // Key structure
}
const MINPADDING = 44

func CFE(securityParameter int) *Cfe {
  return &Cfe{
  		securityParameter: securityParameter,
  		minPaddingSizeBytes: MINPADDING,
  		numBytesInSingleEncryption: securityParameter/8 - MINPADDING,
      nmpke : NMPKE(securityParameter),
  	}
}

func (c *Cfe) Keygen(f []Pair, rCT map[int]([]byte)) int {

  var garblingKey int = 0

  rPT := make(map[int][]byte)

  for key, value := range rCT {

    val, err := c.nmpke.Decrypt(value)
    if err != nil {
      fmt.Println("Keygenbug")
      panic(err)
    }
    rPT[key] = val

  }

  rs_required := make([]byte, len(f) * 4)

  for i := 0; i < len(f); i ++ {

    fi := f[i].GetL().(*big.Int).Int64()
    requiredBlock := int(math.Floor(float64(fi) * 4.0 / float64(c.numBytesInSingleEncryption) ))

    requiredElement := int(fi) % (c.numBytesInSingleEncryption/ 4)

    temp := rPT[requiredBlock]

    copy(temp[requiredElement * 4: (requiredElement + 1) * 4],
      rs_required[i * 4 : (i + 1) * 4])

  }
  //fmt.Println("END OF ICI")

  rs := make([]int, len(f))


  for i := 0; i < len(f); i ++ {
    garblingKey += rs[i] * int(f[i].GetR().(*big.Int).Int64())
  }
  return garblingKey
}

func intArraytoByteArray(iarr []int) []byte {

  l := len(iarr)
  barr := make([]byte, l * 4)
  for i := 0; i < l; i ++ {

    b := make([]byte, 4)
    _ = binary.PutVarint(b, int64(iarr[i]))

    copy(b, barr[i * 4 : (i + 1) * 4] )
  }
  return barr
}

func byteArraytoIntArray(barr []byte) []int {
  l := len(barr)/4
  r := make([]int, l)
  for i := 0; i < l; i ++ {
    val, _ := binary.Varint(barr[i * 4 : (i + 1) * 4])
    r[i] = int(val)
  }
  return r
}

func makeRandomByteArray(n int) []byte{
  buff := make([]byte, n)
  _, err := rand.Read(buff)
  check(err)
  return buff
}

func createR(n int) []int {

  return byteArraytoIntArray(makeRandomByteArray(n * 4))

}

func make2d(n, m int) [][]byte {
  twoD := make([][]byte, n)

   for i := 0; i < n; i++ {
       twoD[i] = make([]byte, m)
   }

   return twoD
}

func (c *Cfe) Enc(pt []int) Pair {

    var l = len(pt) //Size of the array to encrypt
    var lb = l * 4
    var crac = c.numBytesInSingleEncryption // Block size

    //R := createR(l) // Create a random array of int of size l
    Rbytes := makeRandomByteArray(lb)
    R := byteArraytoIntArray(Rbytes)

    n := int(math.Ceil(float64(lb) / float64(crac)))
    ct1 := make2d(n, c.securityParameter / 8) // WARNING Normalement, pas de n

    for i := 0; i < n-1; i++ {
        ct1[i], _ = c.nmpke.Encrypt(Rbytes[i * crac: (i+1) * crac])
    }
    ct1[n-1], _ = c.nmpke.Encrypt(Rbytes[(n-1) * crac: ])

    ct2 := make([]int, l)
    for i := 0; i < l; i++ {
      ct2[i] = R[i] + pt[i]
    }

    var CT Pair
    CT.Set(ct1, ct2)

    return CT
}


func Dec(f []Pair, ct2 []int, garblingKey int ) int {

  output := - garblingKey

  for i := 0; i < len(ct2); i++ {
    val := int(f[i].GetR().(*big.Int).Int64())
    output += ct2[i] * val
  }

  return output
}


//Taken
func GenerateRandomBytes(n int) ([]byte, error) {
  b := make([]byte, n)
  _, err := rand.Read(b)

  if err != nil {
      return nil, err
  }

  return b, nil
}

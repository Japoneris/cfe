// This is the core librairy

package cfe

import (
  "fmt"
  "encoding/binary"
  "math"
  "crypto/rand"
  "bytes"
)

type Cfe struct {
  securityParameter int
  minPaddingSizeBytes int
  numBytesInSingleEncryption int
  nmpke *Nmpke
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

func (c *Cfe) keygen(f []Pair, rCT map[int]([]byte)) int {
  var garblingKey int = 0

  rPT := make(map[int][]byte)

  for key, value := range rCT {
    rPT[key], _ = c.nmpke.Decrypt(value)
  }

  rs_required := make([]byte, len(f) * 4)

  for i := 0; i < len(f); i ++ {
    requiredBlock := int(math.Floor(f[i].GetL().(float64) * 4.0 / float64(c.numBytesInSingleEncryption) ))
    requiredElement := f[i].GetL().(int) % (c.numBytesInSingleEncryption/ 4)
    temp := rPT[requiredBlock]
    copy(temp[requiredElement * 4: (requiredElement + 1) * 4],
      rs_required[i * 4 : (i + 1) * 4])

  }

  rs := make([]int, len(f))


  for i := 0; i < len(f); i ++ {
    garblingKey += rs[i] * f[i].GetR().(int)
  }
  return garblingKey
}

func intArraytoByteArray(iarr []int) []byte {
  l := len(iarr)
  barr := make([]byte, l * 4)
  for i := 0; i < l; i ++ {
    buf := new(bytes.Buffer)
    err := binary.Write(buf, binary.LittleEndian, iarr[i])
    copy(buf.Bytes(), barr[i * 4 : (i + 1) * 4] )
    if err != nil {
        fmt.Println("binary.Write failed:", err)
    }
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

  return byteArraytoIntArray(makeRandomByteArray(n))

}

func make2d(n, m int) [][]byte {
  twoD := make([][]byte, n)

   for i := 0; i < n; i++ {
       twoD[i] = make([]byte, m)
   }

   return twoD
}

func (c *Cfe) Enc(pt []int) Pair {

    R := createR(len(pt))

    Rbytes := intArraytoByteArray(R)

    ct1 := make2d(len(Rbytes) / c.numBytesInSingleEncryption, c.securityParameter / 8)

    n := int(math.Ceil(float64(len(Rbytes)) / float64(c.numBytesInSingleEncryption)))

    for i := 0; i < n; i++ {
        ct1[i], _ = c.nmpke.Encrypt(Rbytes[i*c.numBytesInSingleEncryption: (i+1)* c.numBytesInSingleEncryption])
    }

    ct2 := make([]int, len(pt))

    for i := 0; i < len(pt); i++ {
      ct2[i] = R[i] + pt[i]
    }

    var CT Pair
    CT.Set(ct1, ct2)

    return CT
}


func Dec(f []Pair, ct2 []int, garblingKey int ) int {

  output := - garblingKey

  for i := 0; i < len(ct2); i++ {
    val, _ :=  f[i].GetR().(int)
    output += ct2[i] * val
  }

  return output
}


//Taken
func GenerateRandomBytes(n int) ([]byte, error) {
  b := make([]byte, n)
  _, err := rand.Read(b)
  // Note that err == nil only if we read len(b) bytes.
  if err != nil {
      return nil, err
  }

  return b, nil
}

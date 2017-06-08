package main

import (

	"fmt"
	"github.com/japoneris/cfe"
	"os"
	"math"
	"time"
	"strconv"
	"crypto/rand"
	"math/big"
)

var MAXINT = big.NewInt(2147483647) //int32

func main() {

	if len(os.Args) < 3 {
		return
	}
	s, ok0 := strconv.Atoi(os.Args[1])
	fs, ok1 := strconv.Atoi(os.Args[2])
	sp, ok2 := strconv.Atoi(os.Args[3])

	//fmt.Println(s, fs, sp)
	if ok0 != nil || ok1 != nil || ok2 != nil {
		return
	}

	size := int(s)
	functionSize := int(fs)
	securityParameter := int(sp)

	iterations := 1
	minPaddingSizeBytes := 44
	numBytesInSingleEncryption := securityParameter / 8 - minPaddingSizeBytes

	n := 0 // n must be less than 11 + qqchose
	for i := 0; n <= size + minPaddingSizeBytes / 4; i++ {
    n += numBytesInSingleEncryption / 4;
	}

	fe := cfe.CFE(securityParameter)

	PT := make([]int, n)
	// Generate a random int64 array
	for i := 0; i < n; i++ {
		r1, _ := rand.Int(rand.Reader, MAXINT)
		PT[i] = int(r1.Int64())
	}

	encryptionStartTime := time.Now()
	// Encrypt int64 array 
	CT := fe.Enc(PT)

	encryptionTime := time.Since(encryptionStartTime)

	encryptionSize := len(CT.GetR().([]int)) * 4

	f := make([]cfe.Pair, functionSize)

	for i := 0; i < functionSize; i++ {
		r1, _ := rand.Int(rand.Reader, big.NewInt(int64(n)))
		r2,_ := rand.Int(rand.Reader, MAXINT)
		f[i] = cfe.Pair{r1, r2}
		//fmt.Println(f[i])
	}


	for iter := 0; iter < iterations; iter++ {
		keygenStartTime := time.Now()
		partial_rCT2 := make([]int, functionSize)

		partial_rCT1 := make(map[int][]byte)
		isBlockAdded := make(map[int]string)

		left := CT.GetL().([][]byte)
		right := CT.GetR().([]int)


		j := 0
		for i := 0; i < functionSize; i++ {

			val, bg := f[i].GetL().(*big.Int)
			if bg == false {
				panic("bug")
			}
			val1 := float64(val.Int64())
			requiredBlock := int(math.Floor(val1 * 4.0 / float64(numBytesInSingleEncryption)))
			_, ok := isBlockAdded[requiredBlock]

			//fmt.Println("Pass first", iter, i, val1, numBytesInSingleEncryption, requiredBlock)
			if !ok {
				isBlockAdded[requiredBlock] = "exist"
				partial_rCT1[requiredBlock] = left[requiredBlock]
				j++
			}
			fii := (f[i].GetL()).(*big.Int).Int64()
			ria := right[fii]
			partial_rCT2[i] = ria
		}
		//fmt.Println("We passed here", functionSize)
		garblingKey := fe.Keygen(f, partial_rCT1)
		//fmt.Println("We passed Keygen", garblingKey)

		keygenTime := time.Since(keygenStartTime)
		messageClientToAuthoritySize := len(f) * 8 + len(partial_rCT1) * securityParameter / 8
		//fmt.Println("\tbec here")

		decryptionStartTime := time.Now()
		output := cfe.Dec(f, partial_rCT2, garblingKey)
		decryptionTime := time.Since(decryptionStartTime)

		fmt.Println("**********ENCRYPTION**********")
		fmt.Println("Encryption time: ", encryptionTime.Seconds(), "seconds\n")
		fmt.Println("**********KEY GENERATION**********")
		fmt.Println("Upstream bandwidth: ", messageClientToAuthoritySize/1024,"KBytes")
		fmt.Println("Downstream bandwidth: 8Bytes")
		fmt.Println("Key generation time: ", keygenTime.Seconds(),"seconds")
		fmt.Println("Enc Size: ", encryptionSize,"\n")
		fmt.Println("**********DECRYPTION**********")
		fmt.Println("Decryption time: ", decryptionTime.Nanoseconds()/1000, " MICRO seconds\n");

		fmt.Println("\nOutput; ", output, "\n")
	}




/*
	mapair := cfe.NMPKE(500)
	maclefpublique := mapair.PublicKey()
	fmt.Println("Clef Privée\nAdresse", *mapair,"\nValeur", mapair,"\nRefe", &mapair)
	fmt.Println(mapair.PrivateKey())

	txt := []byte("this is my text to encrypt")
	enc, _  := cfe.Encrypt(maclefpublique, txt)
	encs := string(enc)
	fmt.Println("Message chiffré:\n",enc, encs)


	dec, _ := mapair.Decrypt( enc)
	decs := string(dec)
	fmt.Println("Decryption: \n", decs)

	err := mapair.Storekey("test_01")
	fmt.Println(err)

	manouvelle, _ := cfe.Loadkey("test_01")
	fmt.Println("Clef Privée:\n", manouvelle,"\n", manouvelle.PrivateKey())
	dec1, _ := manouvelle.Decrypt( enc)
	decs1 := string(dec1)
	fmt.Println("Decryption: \n", decs1)

	if decs1 == string(txt) {
		fmt.Println(" PASSED ")
	}

*/
}

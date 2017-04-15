// Don't know what is it yet

package cfe

import (
  "fmt"
  "crypto/rand"
  "crypto/rsa"
  "crypto"
  "io/ioutil"
  "encoding/json"
)
type Nmpke struct {
  /*
  private KeyPair keyPair;
  private Key publicKey;
  private Key privateKey;
  private Cipher cipher;
  private SecureRandom rand;

  cipher string
  rand []byte
  */
  publicKey crypto.PublicKey
  privateKey *rsa.PrivateKey
}


func check(e error) {
    if e != nil {
        panic(e)
    }
}


func NMPKE_root() *Nmpke {

  return &Nmpke{}

}

func NMPKE(keysize int) *Nmpke {

    return generateKeys(keysize)
  }

func generateKeys(keysize int) *Nmpke {

    privKey, err := rsa.GenerateKey(rand.Reader, keysize)
    check(err)

    pubkey := privKey.Public()

    return &Nmpke{
      privateKey: privKey,
      publicKey: pubkey,
    }
	}


func (n *Nmpke) PrivateKey() *rsa.PrivateKey {

    pvk := n.privateKey

    return pvk
}

func (n *Nmpke) PublicKey() *rsa.PublicKey {

    pbk, _ := n.publicKey.(*rsa.PublicKey)

    return pbk
}

func Encrypt(publicKey *rsa.PublicKey, plaintext []byte) ([]byte, error) {
  /*func EncryptPKCS1v15(rand io.Reader, pub *PublicKey, msg []byte) ([]byte, error)*/

    enc, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plaintext)

    check(err)

    return enc, err
  }

func (n *Nmpke) Encrypt(plaintext []byte) ([]byte, error) {
  /*func EncryptPKCS1v15(rand io.Reader, pub *PublicKey, msg []byte) ([]byte, error)*/

    enc, err := rsa.EncryptPKCS1v15(rand.Reader, n.PublicKey(), plaintext)

    check(err)

    return enc, err
  }


func (n *Nmpke) Decrypt(ciphertext []byte) ([]byte, error) {

  /*func (priv *PrivateKey) Decrypt(rand io.Reader, ciphertext []byte, opts crypto.DecrypterOpts) (plaintext []byte, err error)*/

  plaintext, err := n.privateKey.Decrypt(rand.Reader, ciphertext, nil)

  check(err)

  return plaintext, err
}


func (n *Nmpke) Storekey(filename string) error {


    pbk, err := json.Marshal(n.publicKey)
    if err == nil {
      err0 := ioutil.WriteFile(filename + ".publickey", pbk, 0600)
      if err0 != nil {
        fmt.Println("Problem for storing publickey")
        return err0
      }
    } else {
      fmt.Println("Problem for storing publickey")

    }


    pvk, err1 := json.Marshal(n.privateKey)
    check(err1)
    /*pvk, ok1 := n.privateKey.([]byte)*/
    err2 := ioutil.WriteFile(filename + ".privatekey", pvk, 0600)
    if err2 != nil {
      fmt.Println("Problem for storing privatekey")
      return err2
    }
    return nil
  }


func Loadkey(filename string) (*Nmpke, error) {

    public, err0 := ioutil.ReadFile(filename + ".publickey")
    check(err0)

    private, err1 := ioutil.ReadFile(filename + ".privatekey")

    check(err1)

    var pvk *rsa.PrivateKey
    err := json.Unmarshal(private, &pvk)

    if err != nil {
      fmt.Println("Problem for Unmarshal")
      return NMPKE_root(), err
    }

    return &Nmpke{
      publicKey: public,
      privateKey: pvk,
    }, nil
  }

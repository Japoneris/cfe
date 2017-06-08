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
  publicKey crypto.PublicKey // Public Generic Key
  privateKey *rsa.PrivateKey // Private RSA Key
}


func check(e error) {
    if e != nil {
        panic(e)
    }
}


func NMPKE_root() *Nmpke {
  // Just create an empty object
  return &Nmpke{}

}

func NMPKE(keysize int) *Nmpke {
    // Given the keysize (in bytes?), Generate Pub and Secret Keys
    return generateKeys(keysize)
  }

func generateKeys(keysize int) *Nmpke {
    // Generate the rsa key
    privKey, err := rsa.GenerateKey(rand.Reader, keysize)
    check(err)
    // Obtain the complementary key
    pubkey := privKey.Public()
    // Create the nmpke object containing both keys
    return &Nmpke{
      privateKey: privKey,
      publicKey: pubkey,
    }
	}


func (n *Nmpke) PrivateKey() *rsa.PrivateKey {
    // Give access to the Private key
    pvk := n.privateKey

    return pvk
}

func (n *Nmpke) PublicKey() *rsa.PublicKey {
    // Give access to the public key (needed if external people wants to encrypt)
    pbk, _ := n.publicKey.(*rsa.PublicKey)

    return pbk
}

func Encrypt(publicKey *rsa.PublicKey, plaintext []byte) ([]byte, error) {
    // Encrypt several plaintext using the publicKey
    ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plaintext)

    check(err)

    return ciphertext, err
  }

func (n *Nmpke) Encrypt(plaintext []byte) ([]byte, error) {
    // Encrypt but using nmpke object
    ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, n.PublicKey(), plaintext)

    check(err)

    return ciphertext, err
  }


func (n *Nmpke) Decrypt(ciphertext []byte) ([]byte, error) {
    // Decrypt several ciphertexts using private key
    plaintext, err := n.privateKey.Decrypt(rand.Reader, ciphertext, nil)

    check(err)

    return plaintext, err
}


func (n *Nmpke) Storekey(filename string) error {
    // Store in a file the keys
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
    err2 := ioutil.WriteFile(filename + ".privatekey", pvk, 0600)
    if err2 != nil {
      fmt.Println("Problem for storing privatekey")
      return err2
    }
    return nil
  }


func Loadkey(filename string) (*Nmpke, error) {
    // Open 2 files containing pubkey and private key. Return the nmpke obj
    public, err0 := ioutil.ReadFile(filename + ".publickey")
    check(err0)

    private, err1 := ioutil.ReadFile(filename + ".privatekey")

    check(err1)

    var pvk *rsa.PrivateKey
    err := json.Unmarshal(private, &pvk) // The Public key is generic
    //The private need to be correctly formated
    if err != nil {
      fmt.Println("Problem for Unmarshal")
      return NMPKE_root(), err
    }

    return &Nmpke{
      publicKey: public,
      privateKey: pvk,
    }, nil
  }

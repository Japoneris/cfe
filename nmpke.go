// Don't know what is it yet

package nmpke

import (
  "fmt"
  "crypto/rand"
  "crypto/rsa"
  "crypto/cipher"
  "io/ioutil"
)
type nmpke struct {
  /*
  private KeyPair keyPair;
  private Key publicKey;
  private Key privateKey;
  private Cipher cipher;
  private SecureRandom rand;
  */
  keyPair key
  publicKey key
  privateKey key
  cipher string
  secureRandom rand
}


func check(e error) {
    if e != nil {
        panic(e)
    }
}

func (sp nmpke) NMPKE() {}

func (sp nmpke) NMPKE(keysize int) {

    return &nmpke{
      keyPair:,
      publicKey :,
      privateKey : ,
      cipher : "",
      secureRandom :,
    }
  }
func (sp nmpke) generateKeys(keysize int) {

/*
private void generateKeys(int keysize){
    rand = new SecureRandom();

    try {
        KeyPairGenerator generator = KeyPairGenerator.getInstance("RSA", "BC");
        generator.initialize(keysize, rand);
        keyPair = generator.generateKeyPair();
        publicKey = keyPair.getPublic();
        privateKey = keyPair.getPrivate();
    } catch (NoSuchAlgorithmException e) {
        e.printStackTrace();
    } catch (NoSuchProviderException e) {
        e.printStackTrace();
    }


}
*/

	b := make([]byte, keysize)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
  // OR

  pr, err = rsa.GenerateKey(random io.Reader, keysize int) //(*PrivateKey, error)

  if err != nil {
    log.Fatal(err)
  }

  privateKey = &pr
  publicKey = pr.Public()
}

func encrypt(plaintext []byte) []byte {
  func EncryptPKCS1v15(rand io.Reader, pub *PublicKey, msg []byte) ([]byte, error)

/*
public byte[] encrypt(byte[] plaintext){
    byte[] ciphertext = null;
    try {
        cipher.init(Cipher.ENCRYPT_MODE, publicKey, rand);
        ciphertext = cipher.doFinal(plaintext);
    } catch (InvalidKeyException e) {
        e.printStackTrace();
    } catch (BadPaddingException e) {
        e.printStackTrace();
    } catch (IllegalBlockSizeException e) {
        e.printStackTrace();
    }

    return ciphertext;
}
*/
}
func decrypt(ciphertext []byte) []byte {

  func (priv *PrivateKey) Decrypt(rand io.Reader, ciphertext []byte, opts crypto.DecrypterOpts) (plaintext []byte, err error)

}
func storekey(filename string) error {
  err := ioutil.WriteFile(filename + ".publickey", publickey, 0600)

  if err != nil {
    return err
  }
  err := ioutil.WriteFile(filename + ".privatekey", privatekey, 0600)

  return err

  /*
  public void storekey(String filename){
      ObjectOutputStream objectOutputStream = null;
      try {
          objectOutputStream = new ObjectOutputStream(new FileOutputStream(filename + ".publickey"));
          objectOutputStream.writeObject(publicKey);
          objectOutputStream.flush();
          objectOutputStream.close();

          objectOutputStream = new ObjectOutputStream(new FileOutputStream(filename + ".privatekey"));
          objectOutputStream.writeObject(privateKey);
          objectOutputStream.flush();
          objectOutputStream.close();
      } catch (FileNotFoundException e) {
          e.printStackTrace();
      } catch (IOException e) {
          e.printStackTrace();
      }
      */

}

func loadkey(filename string) error {

  public, err := ioutil.ReadFile(filename + ".pulickey")
  check(err)

  private, err := ioutil.ReadFile(filename + ".pulickey")
  check(err)

  publicKey = public
  privatekey = private // Convert to the right type ? + assignement

/*
public void loadkey(String filename){

    try {
        ObjectInputStream objectInputStream = new ObjectInputStream(new FileInputStream(filename + ".pulickey"));
        publicKey = (Key) objectInputStream.readObject();
        objectInputStream.close();

        objectInputStream = new ObjectInputStream(new FileInputStream(filename + ".privatekey"));
        privateKey = (Key) objectInputStream.readObject();
        objectInputStream.close();
    } catch (FileNotFoundException e) {
        e.printStackTrace();
    } catch (IOException e) {
        e.printStackTrace();
    } catch (ClassNotFoundException e) {
        e.printStackTrace();
    }

}
*/

}

/*
import org.bouncycastle.jce.provider.BouncyCastleProvider;

import javax.crypto.BadPaddingException;
import javax.crypto.Cipher;
import javax.crypto.IllegalBlockSizeException;
import javax.crypto.NoSuchPaddingException;
import java.io.*;
import java.security.*;


public class NMPKE {

    public NMPKE(){
        Security.addProvider(new BouncyCastleProvider());
    }

    public NMPKE(int keysize){
        Security.addProvider(new BouncyCastleProvider());
        generateKeys(keysize);
        try {
            cipher = Cipher.getInstance("RSA/None/OAEPWithSHA1AndMGF1Padding", "BC");
        } catch (NoSuchAlgorithmException e) {
            e.printStackTrace();
        } catch (NoSuchPaddingException e) {
            e.printStackTrace();
        } catch (NoSuchProviderException e) {
            e.printStackTrace();
        }
    }





    public byte[] decrypt(byte[] ciphertext){
        byte[] plaintext = null;

        try {
            cipher.init(Cipher.DECRYPT_MODE, privateKey);
            plaintext = cipher.doFinal(ciphertext);
        } catch (InvalidKeyException e) {
            e.printStackTrace();
        } catch (BadPaddingException e) {
            e.printStackTrace();
        } catch (IllegalBlockSizeException e) {
            e.printStackTrace();
        }

        return plaintext;
    }



    }



}

*/

// This is the core librairy

package cfe

import (
  "fmt"
  "io/ioutil"
)

type Cfe struct {
  securityParameter int
  minPaddingSizeBytes int
  numBytesInSingleEncryption int
  my_nmpke nmpke
}
const MINPADDING = 44

func CFE(securityParameter int) *cfe {

  return &Cfe{
  		securityParameter: securityParameter,
  		minPaddingSizeBytes: MINPADDING,
  		numBytesInSingleEncryption: securityParameter/8 - MINPADDING,
      my_nmpke : Nmpke.NMPKE(securityParameter)
  	}

  /*
  public CFE(int securityParameter) {
      this.securityParameter = securityParameter;
      minPaddingSizeBytes = 44;
      numBytesInSingleEncryption = securityParameter / 8 - minPaddingSizeBytes;

      nmPKE = new NMPKE(securityParameter);
  */

}

struct pair_int {
  x, y int
}

func (c cfe) keygen(f []pair_int, rCT func(int, []byte) []byte) int {
  var garblingKey int = 0


/*
public long keygen(ArrayList<Pair<Integer, Integer>> f, Map<Integer, byte[]> rCT) {
    long garblingKey = 0;

    Map<Integer, byte[]> rPT = new HashMap<>();

    int it = 0;
    for (Map.Entry<Integer, byte[]> entry : rCT.entrySet()) {
        Integer key = entry.getKey();
        byte[] value = entry.getValue();
        rPT.put(key, nmPKE.decrypt(value));
    }


    byte[] rs_required = new byte[f.size() * 4];

    for (int i = 0; i < f.size(); i++) {
        int requiredBlock = (int) Math.floor((double) f.get(i).getL() / (double) (numBytesInSingleEncryption / 4));
        int requiredElement = f.get(i).getL() % (numBytesInSingleEncryption / 4);
        byte temp[] = rPT.get(requiredBlock);
        System.arraycopy(temp, requiredElement * 4,
                rs_required, i * 4, 4);
    }

    int[] rs = byteArraytoIntArray(rs_required);

    for (int i = 0; i < f.size(); i++) {
        garblingKey += rs[i] * f.get(i).getR();
    }

    return garblingKey;
}
*/

}

func Enc(pt []int) [][]byte, []int {

/*
public Pair<byte[][], int[]> Enc(int[] PT) {
    int R[] = createR(PT.length);

    byte[] Rbytes = intArraytoByteArray(R);

    byte[][] ct1 = new byte[Rbytes.length / numBytesInSingleEncryption][securityParameter / 8];

    for (int i = 0; i < (int) Math.ceil((double) Rbytes.length / (double) numBytesInSingleEncryption); i++) {
        byte[] temp = new byte[numBytesInSingleEncryption];
        System.arraycopy(Rbytes, i * numBytesInSingleEncryption, temp, 0, numBytesInSingleEncryption);
        ct1[i] = nmPKE.encrypt(temp);
    }

    int[] ct2 = new int[PT.length];

    for (int i = 0; i < PT.length; i++) {
        ct2[i] = R[i] + PT[i];
    }

    Pair<byte[][], int[]> CT = new Pair<byte[][], int[]>();

    CT.set(ct1, ct2);

    return CT;
}

*/

}


func Dec(f []pair_int, ct2 []int, garblingKey int ) int {

  /*
  public long Dec(ArrayList<Pair<Integer, Integer>> f, int[] ct2, long garblingKey) {
      long output = -garblingKey;

      for (int i = 0; i < ct2.length; i++) {
          output += ct2[i] * f.get(i).getR();
      }

      return output;
  }

  */
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


func createR(size int) []int {
  var R [size]int


  /*
  private int[] createR(int size) {
      int[] R = new int[size];
      SecureRandom rand = new SecureRandom();
      for (int i = 0; i < size; i++) {
          R[i] = rand.nextInt();
      }
      return R;
  }
  */
}

func byteArraytoIntArray(byteArray []byte) []int {

/*
private int[] byteArraytoIntArray(byte[] byteArray) {
    IntBuffer intBuf = ByteBuffer.wrap(byteArray).order(ByteOrder.BIG_ENDIAN).asIntBuffer();
    int[] array = new int[intBuf.remaining()];
    intBuf.get(array);
    return array;
}
*/
}

func intArraytoByteArray( intArray []int) []byte {
  /*
  private byte[] intArraytoByteArray(int[] intArray) {
      ByteBuffer byteBuffer = ByteBuffer.allocate(intArray.length * 4);
      IntBuffer intBuffer = byteBuffer.asIntBuffer();
      intBuffer.put(intArray);

      return byteBuffer.array();
  }
  */
}

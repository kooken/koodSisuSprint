/* Given a string (which will be the text to be encrypted/decrypted), a boolean
(which will determine whether to encrypt or decrypt) and another string (which
will be the key), either encrypt or decrypt the text using the XOR cipher.
You must return the encrypted text as a base64 encoded value.
Expected function signature
package sprint

import b64 "encoding/base64"

func XOR(text string, encrypt bool, key string) string {
    // solution code here
}
Usage
XOR("This is a secret message.", true, "test")
>> "IA0aB1QMAFQVRQARFxcWAFQIFgcHBBQRWg=="
XOR("IA0aB1QMAFQVRQARFxcWAFQIFgcHBBQRWg==", false, "test")
>> "This is a secret message." */

package sprint

import (
	b64 "encoding/base64"
)

func XOR(text string, encrypt bool, key string) string {
	// Repeated key generator
	repeatKey := func(length int) []byte {
		repeated := make([]byte, length)
		for i := 0; i < length; i++ {
			repeated[i] = key[i%len(key)]
		}
		return repeated
	}

	if encrypt {
		// Encrypt the text
		keyBytes := repeatKey(len(text))
		encrypted := make([]byte, len(text))
		for i := 0; i < len(text); i++ {
			encrypted[i] = text[i] ^ keyBytes[i]
		}
		return b64.StdEncoding.EncodeToString(encrypted)
	} else {
		// Decrypt the text
		decoded, err := b64.StdEncoding.DecodeString(text)
		if err != nil {
			return "" // Handle invalid input gracefully
		}
		keyBytes := repeatKey(len(decoded))
		decrypted := make([]byte, len(decoded))
		for i := 0; i < len(decoded); i++ {
			decrypted[i] = decoded[i] ^ keyBytes[i]
		}
		return string(decrypted)
	}
}

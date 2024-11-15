/* Instructions
Given a string called key and a string slice called values, return the
correct slice value.
The slice values will contain strings in the format of hash:value. You must
use the "djb2" hash of the key to determine the correct value to return. Do
not return the value with the hash and colon!
Expected function signature
package sprint

import "strconv"

func HashMap(key string, values []string) string {
    // solution code here
}
Usage
HashMap("thekey", []string{"5863446:value1", "193493707:value2",
"6954055932175:value3", "210714995797:value4"})
>> "value3" */

package sprint

import (
	"strconv"
)

func HashMap(key string, values []string) string {
	// djb2 hash function
	hash := func(s string) int64 {
		var h int64 = 5381
		for i := 0; i < len(s); i++ {
			h = ((h << 5) + h) + int64(s[i]) // h * 33 + s[i]
		}
		return h
	}

	keyHash := hash(key)

	for _, v := range values {
		colonIndex := -1
		for i := 0; i < len(v); i++ {
			if v[i] == ':' {
				colonIndex = i
				break
			}
		}

		if colonIndex == -1 {
			continue
		}

		hashPart := v[:colonIndex]
		valuePart := v[colonIndex+1:]

		parsedHash, err := strconv.ParseInt(hashPart, 10, 64)
		if err != nil {
			continue // Skip invalid
		}

		// Check if it matches the computed hash
		if parsedHash == keyHash {
			return valuePart // Return the value after the colon
		}
	}

	return "" // Return an empty string if no match is found
}

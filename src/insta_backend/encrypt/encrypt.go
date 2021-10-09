package encrypt


import (
    "crypto/aes"
    "encoding/hex"
)
 
 
func EncryptString(key []byte, plaintext string) string {
 
    c, err := aes.NewCipher(key)
    CheckError(err)
 
    out := make([]byte, len(plaintext))
 
    c.Encrypt(out, []byte(plaintext))
 
    return hex.EncodeToString(out)
}
 
func DecryptString(key []byte, ct string) string {
    ciphertext, _ := hex.DecodeString(ct)
 
    c, err := aes.NewCipher(key)
    CheckError(err)
 
    pt := make([]byte, len(ciphertext))
    c.Decrypt(pt, ciphertext)
 
    return string(pt[:])
}
 
func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}
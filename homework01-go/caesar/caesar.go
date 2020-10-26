package caesar
m = ['a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z']
b = ['A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z']
func EncryptCaesar(plaintext string, shift int) string {
	var ciphertext string

	ciphertext = ""
    for i in plaintext:
        if i in m:
            index = m.index(i)
            ciphertext += m[(index+shift)%len(m)]
        elif i in b:
            index = b.index(i)
            ciphertext += b[(index+shift)%len(b)]
        else:
            ciphertext += i

	return ciphertext
}

func DecryptCaesar(ciphertext string, shift int) string {
	var plaintext string

    plaintext = ""
    for i in ciphertext:
        if i in m:
            index = m.index(i)
            plaintext += m[(index-shift)%len(m)]
        elif i in b:
            index = b.index(i)
            plaintext += b[(index- shift)%len(b)]
        else:
            plaintext += i
    return plaintext

	return plaintext
}
print(encrypt_caesar(input()))
print(decrypt_caesar(input()))

package vigenere
m = ['a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z']
b = ['A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z']
cip = []
pla =[]
indc = []
indp = []
func EncryptVigenere(plaintext string, keyword string) string {
	var ciphertext string

    ciphertext = ""
    for i in keyword:
        if i in m:
            cip.append(int(m.index(i)))
        elif i in b:
            cip.append(int(b.index(i)))
    for i in plaintext:
        if i in m:
            indc.append(m.index(i))
        elif i in b:
            indc.append(b.index(i))
    if len(cip) < len(indc):
        for i in range (len(indc)-len(cip)):
            cip.append(cip[i])
    for i in plaintext:
        if i in m:
            for j in range(len(cip)):
                ciphertext += m[(indc[j] + cip[j])%len(m)]
        elif i in b:
            for j in range (len(cip)):
                ciphertext += b[(indc[j] + cip[j])%len(b)]
        else:
            ciphertext += i
        break

	return ciphertext
}

func DecryptVigenere(ciphertext string, keyword string) string {
	var plaintext string
    plaintext = ""
    for i in keyword:
        if i in m:
            pla.append(int(m.index(i)))
        elif i in b:
            pla.append(int(b.index(i)))
    for i in ciphertext:
        if i in m:
            indp.append(m.index(i))
        elif i in b:
            indp.append(b.index(i))
    if len(pla) < len(indp):
        for i in range (len(indc) - len(pla)):
            pla.append(pla[i])
    for i in ciphertext:
        if i in m:
            for j in range(len(cip)):
                plaintext += m[(indp[j] - pla[j])%len(m)]
        elif i in b:
            for j in range (len(cip)):
                plaintext += b[(indp[j] - pla[j])%len(b)]
        else:
            plaintext += i
        break
    return plaintext

	return plaintext
}
print(encrypt_vigenere('ATTACKATDAWN', 'LEMONLEMONLE'))
print(decrypt_vigenere('LXFOPVEFRNHR','LEMONLEMONLE'))

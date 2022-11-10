package utils

import "golang.org/x/crypto/bcrypt"

// GenPwd the generated password is irreversible due to use of adaptive hash algorithm
func GenPwd(str string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(hash)
}

// ComparePwd by comparing two string hashes to judge whether they are same plaintext
func ComparePwd(str, pwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(pwd), []byte(str)); err != nil {
		return false
	}
	return true
}

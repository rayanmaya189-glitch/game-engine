package main
import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)
func main() {
	hash, _ := bcrypt.GenerateFromPassword([]byte("SuperAdmin@123"), 12)
	fmt.Println(string(hash))
}

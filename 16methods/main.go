package main

// when these functions actually go into the classes then we call them as methods but here we don't have classes we have structs

import "fmt"

func main() {

	fmt.Println("structs in golang")

	// no inheritance in golang ; no super or parent key word

	palak := User{"Palak", "palak@go.dev", true, 19}
	fmt.Println(palak)
	fmt.Printf("palak details are: %+v\n", palak)
	fmt.Printf("Name is %v and email is %v\n", palak.Name, palak.Email)
	palak.GetStatus()
	palak.NewEmail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active", u.Status)
}
func (u User) NewEmail() {
	u.Email = "test.go@dev"
	fmt.Println("New Email of usee is ", u.Email)
	// this will not change the actual email of palak bc whenever you pass on these objects or structs it actually creates a copy of it
	//nd hence for this pointers are defined when u pass alnog the things into multiple functions usually a copy of that function is passed along and thus pointers comes in

} 

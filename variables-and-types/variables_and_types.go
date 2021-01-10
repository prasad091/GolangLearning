package main

import "fmt"

func main() {
	/* unsigned intergers */
	var number8 uint8 = 0   //unsigned  8-bit integers (0 to 255)
	var number16 uint16 = 0 //Unsigned 16-bit integers (0 to 65535)
	var number32 uint32 = 0 //Unsigned 32-bit integers (0 to 4294967295)
	var number64 uint64 = 0 //Unsigned 64-bit integers (0 to 18446744073709551615)
	fmt.Println(number8)
	fmt.Println(number16)
	fmt.Println(number32)
	fmt.Println(number64)
	/* signed integers */
	var num8 int8 = -127                   // Signed 8-bit integers (-128 to 127)
	var num16 int16 = -32768               // Signed 16-bit integers (-32768 to 32767)
	var num32 int32 = -2147483648          // Signed 32-bit integers (-2147483648 to 2147483647)
	var num64 int64 = -9223372036854775808 // Signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

	fmt.Println(num8)
	fmt.Println(num16)
	fmt.Println(num32)
	fmt.Println(num64)
}

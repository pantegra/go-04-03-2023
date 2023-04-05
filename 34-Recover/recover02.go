/* Alta3 Research | RZFeeser
   Refer - using the defer statement  */

   package main

   import(
	   "os"
	   "io"
	   "fmt"
   )
	   
   func CopyFile(dstName, srcName string) (written int64, err error) {
	   src, err := os.Open(srcName)
	   if err != nil {
			fmt.Println("1")
		   return
	   }
	   defer src.Close()
   
	   dst, err := os.Create(dstName)
	   if err != nil {
		fmt.Println("2")
		   return
	   }
	   defer dst.Close()
   
	   return io.Copy(dst, src)
   }
   
   func main() {
	   CopyFile("/tmp/example-copy2.txt", "/tmp/example-copy.txt")
   }
   

package ali

import "strconv"

func timeConversion(s string) string {
        
           
	str := string([]rune{rune(s[0]),rune(s[1])})
	 nub , err := strconv.Atoi(str)
	 if err != nil {
		 fmt.Println(err)
			 os.Exit(1)
	 }     
 if s[len(s)-2:] == "PM" && nub != 12 {
	 nub += 12
 }
 
 if s[len(s)-2:] == "AM" && nub == 12 {
	 nub = 0
 }
	 str1 := fmt.Sprintf("%02d", nub)
	 return str1 + s[2:len(s)-2]
	
}
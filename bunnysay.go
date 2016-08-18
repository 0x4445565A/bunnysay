/*
* bunnysay.go
*
*  Created on: 2016-08-19
*      Author: Brandon M
*
*   This program is free software: you can redistribute it and/or modify
*    it under the terms of the GNU General Public License as published by
*    the Free Software Foundation, either version 3 of the License, or
*    (at your option) any later version.
*
*    This program is distributed in the hope that it will be useful,
*    but WITHOUT ANY WARRANTY; without even the implied warranty of
*    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*    GNU General Public License for more details.
*
*    You should have received a copy of the GNU General Public License
*    along with this program.  If not, see <http://www.gnu.org/licenses/>.
*
*/
package main

import (
  "bufio"
  "fmt"
  "os"
  "flag"
  "strings"
  "golang.org/x/text/width"
)

func padBunnyString(s string, maxLen int) string {
  var half = (maxLen - len(s)) / 2
  // Not using half twice just incase of odd ints.
  return strings.Repeat(" ", half) + s + strings.Repeat(" ", (maxLen - len(s)) - half)
}

func formatBunnyString(s string, maxLen int) string {
  var l string
  var returnString string
  for i := 0; i < len(s); i++ {
    // Find better white space check
    if (string(s[i]) == " ") {
      // Do we have room to look ahead?
      if (i + (maxLen - len(l)) < len(s) - 1) {
        // Try to break on spaces with look ahead when we can
        if !strings.Contains(s[i+1:i + (maxLen - len(l))], " ") {
          l = padBunnyString(l, maxLen - 1)
        }
      }
    }
    l += string(s[i])

    if (len(l) % maxLen == 0 || i+1 == len(s)) {
      l = "｜" + padBunnyString(l, maxLen) + "｜\n"
      returnString += l
      l = ""
    }
  }
  return returnString
}

func printBunnyWide(s string) {
  fmt.Printf("%s", width.Widen.String(s))
}

func printAnimal(animal string, maxLen int) {
  switch animal {
    case "post":
      var postString = strings.Repeat(" ", (maxLen-1)/2)
      postString += "||"
      postString += strings.Repeat(" ", (maxLen-1)/2)
      postString += "\n"
      printBunnyWide(strings.Repeat(postString, 3))
    default:
      bunnyTop := "(\\__/) ||"
      bunnySpaces := strings.Repeat(" ", (maxLen-len(bunnyTop))/2)
      printBunnyWide(bunnySpaces)
      fmt.Println(bunnyTop)
      printBunnyWide(bunnySpaces)
      fmt.Println("(•ㅅ•) ||")
      printBunnyWide(bunnySpaces)
      fmt.Println("/ 　 づ")
  }
}

func bunnyMaxLen() int {
  return 20
}

func printBunnyString(s string, maxLen int, animal string) bool {
  printBunnyWide("|" + strings.Repeat("￣", maxLen)  + "|\n")
  printBunnyWide(formatBunnyString(s, maxLen))
  printBunnyWide("｜" + strings.Repeat("_", maxLen)  + "｜\n")
  printAnimal(animal, maxLen)
  return true
}

func main() {
  var s string
  animalFlag := flag.String("animal", "bunny", "Type of animal to use. Currently there is bunny and post")
  maxFlag := flag.Int("max", 16, "The maximum canvas of sign in characters.")
  flag.Parse()
  scanner := bufio.NewScanner(os.Stdin)
  stat, _ := os.Stdin.Stat()
  if (stat.Mode() & os.ModeCharDevice) == 0 {
    for scanner.Scan() {
      s += scanner.Text()
    }
  } else {
    s = "No data found in STDIN!"
  }
  printBunnyString(s, *maxFlag, *animalFlag)
}

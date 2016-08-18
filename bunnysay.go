package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
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

func printBunnyString(s string) bool {
  fmt.Println(" _____________________")
  fmt.Println("｜                    ｜")
  fmt.Printf("%s", formatBunnyString(s, 20))
  fmt.Println("｜____________________｜")
  fmt.Println("(\\__/) ||")
  fmt.Println("(•ㅅ•) ||")
  fmt.Println("/ 　 づ")
  return true
}

func main() {
  var s string
  scanner := bufio.NewScanner(os.Stdin)
  stat, _ := os.Stdin.Stat()
  if (stat.Mode() & os.ModeCharDevice) == 0 {
    for scanner.Scan() {
      s += scanner.Text()
    }
  } else {
    s = "No data found in STDIN!"
  }
  printBunnyString(s)
}

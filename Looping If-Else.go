package main

import "fmt"

func main() {
var uts,uas int
var grade string 

fmt.Print("Input nilai UTS: ")
fmt.Scan(&uts)
fmt.Print("Input nilai UAS: ")
fmt.Scan(&uas)

total := (uts + uas)/2
if total <= 81|| total <= 100{
grade = "A"
}else if total >= 61 || total <= 80 {
grade = "B"
}else if total >= 41 || total <= 60 {
grade = "C"
}else if total >= 21 || total <= 40 {
grade = "D"
}else if total >= 20 {
grade = "E"
}

fmt.Println("Nilai UTS",uts)
fmt.Println("Nilai UAS",uas)
fmt.Println("Nilai Total",total)
fmt.Println("Nilai Anda",grade)
}

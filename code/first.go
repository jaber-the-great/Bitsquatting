package main

import (
	"fmt"
	"sort"
)

// Specify the number of domains here
// const numOfDomains int = 10

// func findIndex(list []string, name string) int {
// 	// flipped := []string{"akadls.net", "ggogle-analytics.com", "googleusercoftent.com", "googlevkdeo.com", "instagrcm.com", "rilibili.com", "stackoverdlow.com", "traffic-anager.net", "whatsaqp.com", "xandex.net"}
// 	index := sort.StringSlice(list).Search("name")
// 	return index

// }

func main() {

	// Inserting the IP addresses of experiment machines(della and stanley) into arrays
	// var dellaIP [20]string
	// var stanleyIP [20]string
	// prefix := "192.35.222.2"
	// var IP string = prefix
	// var secIP string = prefix
	// for i := 0; i < numOfDomains*2; i++ {

	// 	if i+5 < 10 {
	// 		IP = prefix + "0" + strconv.Itoa(i+5)
	// 	} else {
	// 		IP = prefix + strconv.Itoa(i+5)
	// 	}
	// 	secIP = "\"" + prefix + strconv.Itoa(i+25) + "\""
	// 	IP = "\"" + IP + "\""
	// 	fmt.Print(IP + ",")
	// 	stanleyIP[i] = secIP
	// 	dellaIP[i] = IP
	// }
	//Testing the correctness of IP range assignment
	// fmt.Println(dellaIP)
	// fmt.Println(stanleyIP)

	// Inserting the flipped and correct domains into array
	// var unflipped [numOfDomains]string
	// var flipped [numOfDomains]string

	// unflipped[0] = "instagram.com"
	// flipped[0] = "instagrcm.com"
	// unflipped[1] = "stackoverflow.com"
	// flipped[1] = "stackoverdlow.com"
	// unflipped[2] = "bilibili.com"
	// flipped[2] = "rilibili.com"
	// unflipped[3] = "googleusercontent.com"
	// flipped[3] = "googleusercoftent.com"
	// unflipped[4] = "whatsapp.com"
	// flipped[4] = "whatsaqp.com"
	// unflipped[5] = "googlevideo.com"
	// flipped[5] = "googlevkdeo.com"
	// unflipped[6] = "google-analytics.com"
	// flipped[6] = "ggogle-analytics.com"
	// unflipped[7] = "akadns.net"
	// flipped[7] = "akadls.net"
	// unflipped[8] = "trafficmanager.net"
	// flipped[8] = "traffic-anager.net"
	// unflipped[9] = "yandex.net"
	// flipped[9] = "xandex.net"

	//theflip := []string{"Jaber", "Iran", "Javad", "12314", "alex"}
	// sort.Sort(sort.StringSlice(flipped[:]))
	// fmt.Println(flipped)
	// flipped := []string{"akadls.net", "ggogle-analytics.com", "googleusercoftent.com", "googlevkdeo.com", "instagrcm.com", "rilibili.com", "stackoverdlow.com", "traffic-anager.net", "whatsaqp.com", "xandex.net"}
	// unflipped := []string{"akadns.net", "google-analytics.com", "googleusercontent.com", "googlevideo.com", "instagram.com", "bilibili.com", "stackoverflow.com", "trafficmanager.net", "whatsapp.com", "yandex.net"}
	// Testing the correctness of value assignment
	// for i := 0; i < numOfDomains; i++ {
	// 	fmt.Println("un: " + unflipped[i])
	// 	fmt.Println("flipped: " + flipped[i])
	// }

	// Important: If it can not find the element, the function would
	// return a value as len of the array. The len must be equal to
	// the num of domains, verifying by printing it
	// fmt.Println(sort.StringSlice.Len(flipped))
	// fmt.Println(sort.StringSlice.Len(unflipped))
	jaber := []string{"j1aber","Danesh", "here"}
	sort.Sort(sort.StringSlice(jaber))
	fmt.Println(jaber)
	//index := 3
	//fmt.Println(flipped[index*2+1])
	// fmt.Println(dellaIP)
	// fmt.Println(stanleyIP)
	// fmt.Println(findIndex(flipped))
	index := sort.StringSlice(jaber).Search("ber")
	fmt.Println(index)
	// fmt.Println(unflipped[index])
	// index = sort.StringSlice(flipped).Search("xandex.net")
	// fmt.Println(index)
	// fmt.Println(unflipped[index])

}

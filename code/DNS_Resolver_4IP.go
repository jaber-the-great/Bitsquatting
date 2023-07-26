package main

// Imports
import (
	"os"
	"fmt"
	"net"
	"sync"
	"time"
	"strings"
	"encoding/json"
	"github.com/miekg/dns"
	"sort"
)

// Hold the global log file
var logEncoder *json.Encoder
// Put the number of domains we control here 
const numOfDomains int = 10

// Put the address for the DNS log file here 
var Logfile string = "/var/log/4IP_per_domain.log"

// put the name of flipped names you control in the following array
var flipped = []string{"akadls.net.", "ggogle-analytics.com.", "googleusercoftent.com.", "googlevkdeo.com.", "instagrcm.com.", "rilibili.com.", "stackoverdlow.com.", "traffic-anager.net.", "whatsaqp.com.", "xandex.net."}

// Put the unflipped version of flipped name in the following array. IMPORTANT: The index of flipped and unflipped
// name in these to array should be equal
var unflipped = []string{"akadns.net.", "google-analytics.com.", "googleusercontent.com.", "googlevideo.com.", "instagram.com.", "bilibili.com.", "stackoverflow.com.", "trafficmanager.net.", "whatsapp.com.", "yandex.net."}

// Della and stanley are the machines used in this experiment. Della serves A.com (the unflipped websites)
// Stanley serves B.com (the flipped websites)
var dellaIP =[]string {"192.35.222.205","192.35.222.206","192.35.222.207","192.35.222.208","192.35.222.209","192.35.222.210","192.35.222.211","192.35.222.212","192.35.222.213","192.35.222.214","192.35.222.215","192.35.222.216","192.35.222.217","192.35.222.218","192.35.222.219","192.35.222.220","192.35.222.221","192.35.222.222","192.35.222.223","192.35.222.224"}
var stanleyIP = []string {"192.35.222.225","192.35.222.226","192.35.222.227","192.35.222.228","192.35.222.229","192.35.222.230","192.35.222.231","192.35.222.232","192.35.222.233","192.35.222.234","192.35.222.235","192.35.222.236","192.35.222.237","192.35.222.238","192.35.222.239","192.35.222.240","192.35.222.241","192.35.222.242","192.35.222.243","192.35.222.244"}

var flip_x string  // the query is for flipped domain, x is the IP of unflipped domain
var flip_y string // the query is for flipped domain, y is the IP of flipped domain
var unflip_z string // the query is for unflipped domain, z is the IP of unflipped domain
var unflip_w string // the query is for unflipped domain, w is the IP of flipped domain


// use this structure to bind IP address and ports 
var binds = [] string {"127.0.0.1","192.35.222.16","192.35.222.17"}





// Map of flipped domain to correct domain
var flippedDomains = make(map[string]string)



// Find element in the array:
func Find(a []string, x string) int {
    for i, n := range a {
        if x == n {
            return i
        }
    }
    return len(a)
}




// Makes a reply msg given a request
func createReply(req *dns.Msg) *dns.Msg {

	// Create a response/reply msg
	m := new(dns.Msg)
	m.SetReply(req)

	// Save our precious bandwidth if possible
	m.Compress = true

	// Add a disclaimer extra to every response
	m.Extra = make([]dns.RR, 1)
	m.Extra[0] = &dns.TXT{
		Hdr: dns.RR_Header{
			Name: m.Question[0].Name,
			Rrtype: dns.TypeTXT,
			Class: dns.ClassINET,
			Ttl: 604800,
		},
		Txt: []string{"This server is part of a research project at seclab UCSB: www.seclab.cs.ucsb.edu "},
	}

	// Return it
	return m

}

// Define the format of a log
type LogEntry struct {
	Timestamp string `json:"timestamp"`
	Type uint8 `json:"type"`
	Destination string `json:"dst"`
	Source string `json:"src"`
	Port string `json:"port"`
	QName string `json:"qName"`
	QType uint16 `json:"qType"`
	QClass uint16 `json:"qClass"`
}

// Logs the query to the log file
func logQuestion(w dns.ResponseWriter, req *dns.Msg, queryType uint8) {

	// Grab the local address info
	localAddress, _, err := net.SplitHostPort(w.LocalAddr().String())
	if err != nil {
		panic(err)
	}
	remoteAddress, remotePort, err := net.SplitHostPort(w.RemoteAddr().String())
	if err != nil {
		panic(err)
	}

	// Make a new LogEntry
	log := LogEntry{
		Timestamp: time.Now().Format("2006-01-02T15:04:05"),
		Type: queryType,
		Destination: localAddress,
		Source: remoteAddress,
		Port: remotePort,
		QName: req.Question[0].Name,
		QType: req.Question[0].Qtype,
		QClass: req.Question[0].Qclass,
	}
	
	// Log it as JSON
	err = logEncoder.Encode(log)
	if err != nil {
		panic(err)
	}

}

// Answers flipped requests
// #############################################################################
// For requests of b.com(flipped of a.com), we respond a.com @ x and b.com @ y
// #############################################################################
func FlipServer(w dns.ResponseWriter, req *dns.Msg) {

	// Attempt to log the request to the control domain
	logQuestion(w, req, 3)

	// Create a response/reply msg
	m1 := createReply(req)
// ######################################
// ### TODO: change ns1 and ns2 config.control and put
// The domain names there to pass the bailiwick test
        host_parts_temp := strings.Split(m1.Question[0].Name, ".")

	// Replace the domain with the correct value
	host_parts_temp[len(host_parts_temp) - 3] = flippedDomains[strings.Join(host_parts_temp[len(host_parts_temp) - 3:len(host_parts_temp) - 1], ".")]
	index := sort.StringSlice(flipped).Search(m1.Question[0].Name)
	if index >= numOfDomains {
	//TODO :  handle this issue and make the system robust
		fmt.Println("flipped Name not found in the top ")
	}
	flip_x = dellaIP[index] // for A.com
	flip_y = stanleyIP[index] // for B.com
	if m1.Question[0].Qclass == 1 {
		if m1.Question[0].Qtype == 1 {
			m1.Answer = []dns.RR{
				&dns.A{
					Hdr: dns.RR_Header{
						Name: m1.Question[0].Name,
						Rrtype: dns.TypeA,
						Class: dns.ClassINET,
						Ttl: 1,
					},
					A: net.ParseIP(flip_y),
				},
				&dns.A{
					Hdr: dns.RR_Header{
						Name: unflipped[index],
						Rrtype: dns.TypeA,
						Class: dns.ClassINET,
						Ttl: 1,
					},
					A: net.ParseIP(flip_x),
				},
			}
		} else if m1.Question[0].Qtype == 2 {
			m1.Answer = []dns.RR{
				&dns.NS{
					Hdr: dns.RR_Header{
						Name: m1.Question[0].Name,
						Rrtype: dns.TypeNS,
						Class: dns.ClassINET,
						Ttl: 1,
					},
					Ns: "ns1." + m1.Question[0].Name ,
				},
				&dns.NS{
					Hdr: dns.RR_Header{
						Name: m1.Question[0].Name,
						Rrtype: dns.TypeNS,
						Class: dns.ClassINET,
						Ttl: 1,
					},
					Ns: "ns2." + m1.Question[0].Name ,
				},
			}
		} else if m1.Question[0].Qtype == 6 {
			m1.Answer = []dns.RR{
				&dns.SOA{
					Hdr: dns.RR_Header{
						Name: m1.Question[0].Name,
						Rrtype: dns.TypeSOA,
						Class: dns.ClassINET,
						Ttl: 1,
					},
					Ns: "ns1." + m1.Question[0].Name ,
					Mbox: "msc.jaber@gmail.com" + ".",
					Serial: 1,
					Refresh: 28800,
					Retry: 7200,
					Expire: 604800,
					Minttl: 60,
				},
			}
		} else if m1.Question[0].Qtype == 15 {
			m1.Answer = []dns.RR{
				&dns.MX{
					Hdr: dns.RR_Header{
						Name: m1.Question[0].Name,
						Rrtype: dns.TypeMX,
						Class: dns.ClassINET,
						Ttl: 1,
					},
					Preference: 1,
					Mx: "mx." + m1.Question[0].Name ,
				},
			}
		} 
	}

	
	w.WriteMsg(m1)

}

// Answers control domain requests
// Why should we answer the control domain? 


// #################### TODO  ################################# 
// Answers unflipped domain requests
func UnflippedServer(w dns.ResponseWriter, req *dns.Msg) {

	// Attempt to log the request to the control domain
	logQuestion(w, req, 2)

	// Create a response/reply msg
	m := createReply(req)

        index := Find(unflipped , m.Question[0].Name)
        if index >= numOfDomains {
        // handle this issue and make the system robust
                fmt.Println("flipped Name not found")
        }
        unflip_z = dellaIP[index +10] // for A.com
        unflip_w = stanleyIP[index+10] // for B.com


	// Make a reply
	if m.Question[0].Qclass == 1 {
		if m.Question[0].Qtype == 1 {
			m.Answer = []dns.RR{
				&dns.A{
					Hdr: dns.RR_Header{
						Name: m.Question[0].Name,
						Rrtype: dns.TypeA,
						Class: dns.ClassINET,
						Ttl: 1,
					},
					A: net.ParseIP(unflip_z),
				},
				&dns.A{
					Hdr: dns.RR_Header{
						Name: flipped[index],
						Rrtype: dns.TypeA,
						Class: dns.ClassINET,
						Ttl: 1,
					},
					A: net.ParseIP(unflip_w),
				},
			}
		} else if m.Question[0].Qtype == 2 {
			m.Answer = []dns.RR{
				&dns.NS{
					Hdr: dns.RR_Header{
						Name: m.Question[0].Name,
						Rrtype: dns.TypeNS,
						Class: dns.ClassINET,
						Ttl: 604800,
					},
					Ns: "ns1." + m.Question[0].Name + ".",
				},
				&dns.NS{
					Hdr: dns.RR_Header{
						Name: m.Question[0].Name,
						Rrtype: dns.TypeNS,
						Class: dns.ClassINET,
						Ttl: 604800,
					},
					Ns: "ns2." + m.Question[0].Name + ".",
				},
			}
		} else if m.Question[0].Qtype == 6 {
			m.Answer = []dns.RR{
				&dns.SOA{
					Hdr: dns.RR_Header{
						Name: m.Question[0].Name,
						Rrtype: dns.TypeSOA,
						Class: dns.ClassINET,
						Ttl: 604800,
					},
					Ns: "ns1." + m.Question[0].Name + ".",
					Mbox: "msc.jaber@ucsb.edu" + ".",
					Serial: 1,
					Refresh: 28800,
					Retry: 7200,
					Expire: 604800,
					Minttl: 60,
				},
			}
		} else if m.Question[0].Qtype == 15 {
			m.Answer = []dns.RR{
				&dns.MX{
					Hdr: dns.RR_Header{
						Name: m.Question[0].Name,
						Rrtype: dns.TypeMX,
						Class: dns.ClassINET,
						Ttl: 1,
					},
					Preference: 1,
					Mx: "mx." + m.Question[0].Name + ".",
				},
			}
		} 
	}

	// Send back a normal response
	w.WriteMsg(m)

}

// Answers unknown domain requests
func UnknownServer(w dns.ResponseWriter, req *dns.Msg) {

	// Attempt to log the request to an unknown domain
	logQuestion(w, req, 0)

	// Create a response/reply msg
	m := createReply(req)

	// Refuse it since we don't know how to handle it
	m.Rcode = 5

	// Send back the answer 
	w.WriteMsg(m)

}

// Spawns a listening server given the bind info
func startListening(mux *dns.ServeMux, address string, port string) {

	// Create a dns server
	server := &dns.Server{
		Addr: address + ":" + port,
		Net: "udp",
		Handler: mux,
	}

	// Begin listening and serving requests
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

// Entry Point
func main() {


	// Open the log file specified in the Config
	logFile, err := os.OpenFile(Logfile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600);
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	// Create an encoder
	logEncoder = json.NewEncoder(logFile)

	// Create a ServeMux instance
	mux := dns.NewServeMux()

	// Set the default handler to UnknownServer
	mux.HandleFunc(".", UnknownServer)


	for iter:=0; iter<numOfDomains;iter++{
		// Creating handler for unlipped domains 
		mux.HandleFunc(unflipped[iter],UnflippedServer)
		// Creating handler for flipped domains
		mux.HandleFunc(flipped[iter], FlipServer)
	}


	// Each listener will be spawned in a new goroutine
	var wg sync.WaitGroup

	// Loop each bind and create a new goroutine with a dns server and begin listening
	
	for _,bind := range binds {

		// We have 1 more to wait for
		wg.Add(1);

		// Spawn as a goroutine
		go startListening(mux, bind, "53")

	}

	// Wait for them all to exit
	wg.Wait()

}

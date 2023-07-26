## Ping command, verbose issue
* Different implementations of ping
### The one in IPUtils: https://github.com/iputils/iputils/blob/master/ping/ping.c
* The next line is the verbose part
* printf(_("PING %s (%s) "), rts->hostname, inet_ntoa(rts->whereto.sin_addr));
	if (rts->device || rts->opt_strictsource)
		printf(_("from %s %s: "), inet_ntoa(rts->source.sin_addr), rts->device ? rts->device : "");
	printf(_("%zu(%zu) bytes of data.\n"), rts->datalen, rts->datalen + 8 + rts->optlen + 20);

* The rts->hostname gets the hostname: I pinged jaberxyz.com(with spoofed DNS answer) and the verbose was: 
  * PING b.com (192.168.58.40) 56(84) bytes of data

* First, the ping command checks if the input is an IP address or not by using inet_aton(). If the inet_aton() can convert the input to the binary format of an address which is valid, ping would use that otherwise it would use getaddressinfo() to retrieve the IP address of hostname. 
#### Solution found:
* Based on what is shows in the picutre, we pinged jaberxzy.com and the DNS resolver replied b.com is in 192.168.58.40.
The question was, why in the verbose, we see b.com and not jaberxyz.com. 
The answer is here: https://github.com/iputils/iputils/blob/master/ping/ping.c line 694
When the input of ping is a hostname( not an IP address), it resloves the IP address using **getaddressinfo()** in line 689. After a successful address resolution, it copies the name retieved by getaddressinfo() _ which is b.com and stored. Then, in line 912, it uses &rts->whereto to print the verbose. So, in general, a strncpy() after DNS resolution caused that problem in the verbose.



### This implementation uses: https://github.com/dgibson/iputils/blob/master/ping.c
* uses gethostbyname(), 
#### Solution:
* Same thing in this implementation of ping. The memcpy() updates the hostname after DNS resolution in line 290. This implementation uses **gethostbyname()** for DNS lookup.



* Get hostbyname() is also used in getaddressinfo() and getaddressinfo() is used by nslookup (double check this part)
* gethostbyname() uses gethostent() and I think gethostent() should check for equality of request ID and the ID in the response 
  
### Getaddressinfo:
* explore_fqdn() is a function to look at
  * It prefers IPv6 to IPv4
  * Uses getipnodebyname() and gethostbyname2() and gethostbyname() under 3 different conditions
  
#### Gethostbyname at http://web.mit.edu/ghudson/sipb/pthreads/net/gethostbyname.c:
* Explore the one used by glibc getXXbyYY: https://code.woboq.org/userspace/glibc/nss/getXXbyYY.c.html
* if given hostname is not an IP address, it uses file_find_name() function
* file_find_name calls the gethostent_r() function to retrieve all host entries 
  - [x] Could not find gethostent_r() source code. Look for it. Should be in netdb.h. Look more into the included libraries to find the answer. 
    - [x] Found it http://web.mit.edu/ghudson/sipb/pthreads/net/gethostent.c ; It does not seem it is used by glibC. Use the correct source code 
- gethostbyname() checks the equality of given hostname and the resolved hostname. The issue is that the gethostent_r copies the hostname from response into the result, and then returns it to the gethostbyname():
  - In normal cases, we uses res_search and res_parse_answer
  - result->h_name = p; This is used after a connection error
- [ ] check how gethostent_r() retrieves the hostname in DNS response. Maybe as I guessed, it reads it from Answer section of DNS response
- [ ] Check where these codes check the equality of query ID and response(I did not see anything in the source code). Maybe it does not check cause I saw mutex in gethostent_r(). Check it by running an experiment. 
  
#### getipnodebyname(): https://fossies.org/linux/arla/lib/roken/getipnodebyname.c
* first checks gethostbyname2() and if it does not have it, it would try gethostbyname()
* At the end, it replaces tmp(which is return value of gethostbyname/2) with copyhostent(tmp) and then returns it.
- [ ] Check copyhostent() function


### Dig : https://opensource.apple.com/source/bind9/bind9-42.3/bind9/bin/dig/dig.c.auto.html
* dns_message_currentname(msg, DNS_SECTION_ANSWER, &name)
* 			strncpy((*lookup)->textname, textname,
  sizeof((*lookup)->textname));
  debug("looking up %s", (*lookup)->textname);
* getaddresses function 


### Dig in bind9: /bin/dig/dig.c
* nslookup and host are in the same directory
* 
  
### Telnet: https://elixir.bootlin.com/busybox/latest/source/networking/telnet.c
* xmove_fd(create_and_connect_stream_or_die(host, port), netfd);
* Unlike ping, telnet does not have hostname problem in the verbose. It prints the hostname given as the argument of command
  - [ ] Check this in experiment 


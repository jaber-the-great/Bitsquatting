1. Honey pot
2. Quanta networking devices
3. Drive-by download attack
4. Rouge software
5. Fat finger distance :)
6. QWERTY vs AZERTY vs QWERTZ keyboard 
7. domain name tasting and domain-name front running.
8. Homograph attack
9. Using non latin alphabet which look like latin alphabets in the domain names etc
10. Visually deceptive texts like paypal and paypa1 
11. Using Quantom for oblivious data transfer 
12. Wildcard certificate
13. Cargo cult
14. Adding people to order from doordash in group :)
15. byobu terminal specially for terminal division
16. Work with Ansible files
17. DNS stuff: dig, whois and nslookup --> does the dns querier perform reverse request with something like nslookup to determine the validity of the response? 
18. For test 
19. Enter ".key" after the url of your github(github/your_id.keys) and get a URL for your public keys
20. If there is an space at the beginning of a command, the up arrow can not resolve it, also the history does not resolve it 
21. www --> dab dab dab
22. Dangling links
23. NS records are in-bailiwick
24. Glue record to remove the circular dependency, they can also be found on the DNS response 
25. Domain name registrar
26. AXFR query for zone transfer in DNS
27. In order to gather information concerning the clients behind
recursive resolvers, we enable ECS (EDNS Client Subnet) [10] on
both NS M1 and M2.
28. Dormant attack, open resolver, false positive attacks --> read about them
29. Dropping the scope name is also one of the attacks in typosquatting if the nmp packages. eg: Original: @azure/core  and squatting is: core
30. Dependency confusion attack --> Indicated by high version number 
31. Attack vector and attack surface 
32. PII --> personally identifiable information
33. Data exfiltration/extrusion/exportation(all are same)
34. bug bounty hunting
35. clumsy test
36. reproducible computing 
37. Stealing domains by Hijacking email address in WHOIS:
    1.  Re-registering expired emails 
    2.  Re-registering the expired domain and create the mail account
    3.  Using Email-Hippo to see whether emails are available or not 
38. ubuntu multipass
39. tcmdump and pcap instead of wireshark --> wireshark is better
40. Justin Maness for all of the networking issues on campus(jmaness@ucsb.edu) 
41. chrome://net-internals/#events for checking the DNS and other stuff on chrome browser
42. Nigerian scams
    

## DNS May be hazardous to your health

- [ ] Do I need to remove the names queried by the third party software and local names?
- [ ] Gstatic.com is the most visited, cause serving static content like css, images, javascripts etc --> look for the most popular one right now  
- [ ] What is the difference between www.google.com and www.google.com. and www and google.com in terms of dns? 
- [ ] 

## DNS deep dive
* Simultanious IPV4 and IPV6 query and accept the first response
  * They say the ordering doesn't matter but it matters (the timeout related to UDP when we don't respond the new messages)
* The data centers are running in high temperature to save energy as google recommended 
* 



## npm packages:
* Got the list of packages with 2M package
* Can repead the other work this time considering nmp packages:
    * Query the list of packages over time
    * Observe newly uploaded packages 
    * Check for sqatters 
    * They did it with DNS domains
* For typosqatting:
    * Considere fat finger
    * Reordering of charectars(like o and u)
    * Several stroke on one char
    * Other ways? --> Weng et al

* Two factor authentication for popuar npm packages 
* Jfrog packages to identify malicious nm packages 

## dns lookup in linux
* using libc -> gethostbyname() and gethostbyaddress()
* nslookup does not use libc to retrieve 
* There is, however, a standard C library call called which many programs use: getaddrinfo. But not all applications use this
* strace -e trace=open -f host google.com
* We have different variations of ping(different implementation)
* ping uses nsswitch but host does not 
*  An “interesting” read is diving into how golang applications resolves things. The golang library tries to guess things about nss and fake it. A testcase to show how broken it is can for example be to install nss-myhostname and rm /etc/hosts … The nss module should allow you to resolve localhost and your hostname, but since golang lib doesn’t actually call into the libc and thus nss stack it doesn’t work. Run tcpdump and you’ll see the queries go to your (isp) dns resolver!
* libresolv is the old, traditional and standard way. It is available on every Unix (type man 3 resolver) and includes routines like res_query which does more or less what you want. To query a specific name server, you typically update the global variable _res.nsaddr_list (do note that, apparently, it does not work with IPv6) --> [for IPv6 use](https://stackoverflow.com/questions/388930/using-the-name-resolver-of-resolv-h-with-ipv6)
* The goal of ldns is to simplify DNS programming, it supports recent RFCs like the DNSSEC documents, and allows developers to easily create software conforming to current RFCs, and experimental software for current Internet Drafts. A secondary benefit of using ldns is speed; ldns is written in C it should be a lot faster than Perl. [documentation](https://www.nlnetlabs.nl/documentation/ldns/tutorial1_mx.html)
*  a very common library, but apparently unmaintained, is [adns](http://www.chiark.greenend.org.uk/~ian/adns/)
* [dns library](http://freshmeat.sourceforge.net/projects/poslib)
* some c libraries for dns:
  * adns
  * udns
  * dns.c
  * FireDNS
  * /networking/dnsd.c --> just checks the equality of query string and the query section in the DNS answer
  * getaddressinfo --> based on RFC 3484(year 2003) and RFC 4291 (year 2006 and is about IP v6) but RFC 5452 suggests for equality check which is in 2009
  * Root dns servers: a.root-servers.net to m.root-servers.net
  * The getaddrinfo() function combines the functionality provided by the gethostbyname(3) and getservbyname(3) functions into a single interface, but unlike the latter functions, getaddrinfo() is reentrant and allows programs to eliminate IPv4-versus-IPv6 dependencies.


## nodes and todo from papers:
* Bitflips in routers and switches
  * For transferring packet
  * For stored routing information: Can cause oscillation etc(is it security or robustness)
  * Bitflip in FIB: can be used to redirect traffic and perform DOS attack or black hole
  * Inter AS level: Unwanted hijacking of a prefix
    * Can I measure that?
    * Is the defense different from regular hijacking 
- [ ] Distinguish between bitflips in client side and DNS resolution process
- [ ] Better distinguishing of typos and bitflips(not only fat finger)
* In DNS resolution, lowercase and uppercase letters in domain name are considered the same but not in NMP packages
*  
 

## Nginx
* location of index.html: /var/www/html/
* location of log file: /var/log/nginx/access_log
* changing the format of log:
  * /etc/nginx/nginx.conf
  * added the following line to the log setting:
  * log_format compression '$remote_addr - $remote_user [$time_local] "$host" '
  '"$request" $status $body_bytes_sent '
  '"$http_referer" "$http_user_agent" "$gzip_ratio"';

  ## New Notes:
  * Combo squatting 
  * Skill squatting 

## Honeypot:
* https://bruteforce.gr/honeydrive/

* When asked for TLD(eg .com) from root servers, `RRSIG` was associated with the NS records:
  * An RRSIG record contains thes signature for an RRset with a particular name, class, and type. The RRSIG RR specifies a validity interval for the signature and uses the Algorithm, the Signer's Name, and the Key Tag to identify the DNSKEY RR containing the public key that a validator can use to verify the signature.
  * It's part of DNSSEC:
    * https://simpledns.plus/help/rrsig-records
  * 

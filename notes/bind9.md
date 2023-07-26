## Search these stuff
* Split-horzion dns?
* catalog zone?
* DNSSEC ?
* Response Policy zone?
* response rate limiting?
* recursive query limit


## Upsates
* 5817.   [security]      The rules for acceptance of records into the cache have been tightened to prevent the possibility of poisoning if forwarders send records outside


## Notes and todo from bin9
* TLD, gTLD, ccTLD etc
* Stub resolver(for caching)
* NXDOMAIN --> error, the name does not exist
- [ ] Can referrals be used as an attack vector 
- [ ] DNS over TLS
* DDNS: Dynamic DNS keeps DNS records automatically up to date when an IP address changes
* RNDC: The remote name daemon control program allows the system administrator to control the operation of a name server. The majority of BIND 9 packages or ports come preconfigured with local (loopback address) security preconfigured. If rndc is being invoked from a remote host, further configuration is required.
* The nsupdate tool uses Dynamic DNS (DDNS) features and allows users to dynamically change the contents of the zone file(s). nsupdate access and security may be controlled using named.conf
* If the remote hosts used for either rndc or DDNS lie within a network entirely under the user’s control, the security threat may be regarded as non-existent
* resolvectl flush-caches
* sudo named-chechconf  is used for checking configuration
* You can hide dns server version for security reasons
* TCP is alos used for response size larger than 512 bytes. 
* installing bind9utils is important for rnds stuff
* Port 953 is TCP is used
* sudo journalctl -eu named to view the qurey log
* Systemd-resolved provides the stub resolver on Ubuntu 22.04/20.04. As mentioned in the beginning of this article, a stub resolver is a small DNS client on the end-user’s computer that receives DNS requests from applications such as Firefox and forward requests to a recursive resolver.
* systemd-resolve --status (this command is deprecated, use the next one):
  * resolvectl status
* The master DNS server holds the master copy of the zone file. Changes of DNS records are made on this server. A domain can have one or more DNS zones. Each DNS zone has a zone file which contains every DNS record in that zone. For simplicity’s sake, this article assumes that you want to use a single DNS zone to manage all DNS records for your domain name.
* A zone file can contain 3 types of entries:
  * Comments: start with a semicolon (;)
  * Directives: start with a dollar sign ($)
  * Resource Records: aka DNS records
  * Domain names must end with a dot (.), which is the root domain. When a domain name ends with a dot, it is a fully qualified domain name (FQDN).
  *  yyyymmddss, where yyyy is the four-digit year number, mm is the month, dd is the day, and ss is the sequence number for the day. You must update the serial number when changes are made to the zone file.
  *  named-checkzone db
*  sudo journalctl -eu named to view the qurey log
* Systemd-resolved provides the stub resolver on Ubuntu 22.04/20.04. As mentioned in the beginning of this article, a stub resolver is a small DNS client on the end-user’s computer that receives DNS requests from applications such as Firefox and forward requests to a recursive resolver.
* systemd-resolve --status (this command is deprecated, use the next one):
  * resolvectl status
* To change the default DNS resolver:
  * edit /etc/systemd/resolved.conf
  * set: DNS = \<desired address>
  * systemctl restart systemd-resolved
  * Check the conf by resolvectl
* For your domain to function properly on the internet, you must reply to SOA and NS queries from your DNS server, otherwise some DNS resolvers won't like the way it's set up and fail the lookups.
* AD in dns flags --> (Authentic data)verified by DNSSEC
* CD is dns flags --> (checking disabled) --> accept non DNSSEC responses
* netstat -tulpn(tcp, udp, listening, programs, ) --> the programs which are listening
* lsop -i -P -n ---> like above but shows open files 
* commands to check open ports:
  * lsof
  * ss
  * netstat
  * nmap
  * man 5 services 
* to listen to a port on specific IP address:
  * nc -l -p \<port number> -s \<IP address>
* When I stopped bind at dns server, my spoofed packets were no longer accepted by the client. Even the udp/tcp ports of both computer where open and I could see the packets but they were refused. I checked the ports that linux is listening to them when bind is running and when bind is not running. It turned out that after opening udp port 53 on the interface connecting dns server and client, it worked:
    * nc -ul -p 53 -s \<IP address which was 192.168.56.30>
* In previous setup when bind was up and running and I responded back by my spoof script, it only worked when the domain name was a non-existing wired name like jaberxyz1.com but after disabling bind9 and listening to the udp port. I though this problem was because of the fact that bind tries to get that domain on the internet and since it can not find it, it takes longer and my spoofer can respond and the client accp
* I also thought the dns respond sent back by bind is considered a standard one(with authorization, list of name servers, SOA etc) and that's why spoofing google.com is not working. But I realized that opening and listenning the port for udp/tcp on 53 is the key. In this way, it always accept my spoofed result
* ps -em | grep python ---> for killing python program
* Using bf-dns.go file of bitfl1p.com project:
  * cd \<project directory>
  * go mod init example.com/m
  * go mod tidy
  * go get github.com/miekg/dns
  * go build github.com/miekg/dns
  * go build bf-dns.go
  * ./bf-dns config.yaml
  * If it has some problem with listening to the port, check no other process is using that port(eg bind is running in background should be stopped)
* Due to its arcane user interface and frequently inconsistent behavior, we do not recommend the use of nslookup. Use dig instead.
* A DNS firewall examines DNS traffic and allows some responses to pass through while blocking others. This examination can be based on several criteria, including the name requested, the data (such as an IP address) associated with that name, or the name or IP address of the name server that is authoritative for the requested name. Based on these criteria, a DNS firewall can be configured to discard, modify, or replace the original response, allowing administrators more control over what systems can access or be accessed from their networks.
* DNS Response Policy Zones (RPZ) are a form of DNS firewall in which the firewall rules are expressed within the DNS itself - encoded in an open, vendor-neutral format as records in specially constructed DNS zones.
* A response policy rule in a DNS RPZ(response policy zone) can be triggered as follows:
  * By the query name
  * By the address which would be present in a truthfull response
  * By the name or address of an authoritative name server responsible for publishing the original response
* A response policy action can be one of the following:
  * to synthesize a “domain does not exist” (NXDOMAIN) response
  * to synthesize a “name exists but there are no records of the requested type” (NODATA) response
  * to replace/override the response’s data with specific data (provided within the response policy zone)
  * to exempt the response from further policy processing
*  Can use the firewall and RPZ to block the phishing domain by having a list of phishing websites and responding NXDOMAIN
* Search in the github repo:
  * client->query.qname
  * client->query.origqname
  * exist in lib/ns/query.c and client.c
* rndc flush --> to flush the cache of dns resolver
  


## Checking history of qname and rname eqality check
* [RFC 1035]The next step is to match the response to a current resolver request.
The recommended strategy is to do a preliminary matching using the ID
field in the domain header, and then to verify that the question section
corresponds to the information currently desired.  This requires that
the transmission algorithm devote several bits of the domain ID field to
a request identifier of some sort. 
* lib/ns/query.c
  * dns_name_eqal
  * qname is equal
  * recparam: check the equality of recursion parameters including the qname

- [ ] Where does it check the ID equality? The qname equality should be in the same place
* qname, rrname, rname, origname, cur, origin
* Checing default RPZ
* /lib/ns/client.c
* /lib/dns/gssapictx.c and other gssapis 
* Tkey: Transaction key (maybe the same thing as ID)
* /lib/dns/validator.c (not much)
* /lib/dns/message.c 
* /lib/bind9/check.c (not much)
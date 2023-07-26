  ## Bailiwick check:
  * Bitfl1p.com project uses a control domain as the DNS server for all of the flipped domains it is serving. 
  * This may(with high probability) result in failure of bailiwick check.
  * Not all of the resolvers check for bailiwick but the ones that check, would drop the records which are received from an `out of bailiwick` nameserver.
  * Farsight checks in/out of bailiwick
  * Based on RFC 7719:
    * '   In-bailiwick:
  
      (a) An adjective to describe a name server whose name is either
      subordinate to or (rarely) the same as the zone origin.  In-
      bailiwick name servers require glue records in their parent zone
      (using the first of the definitions of "glue records" in the
      definition above).

      (b) Data for which the server is either authoritative, or else
      authoritative for an ancestor of the owner name.  This sense of
      the term normally is used when discussing the relevancy of glue
      records in a response.  For example, the server for the parent
      zone "example.com" might reply with glue records for
      "ns.child.example.com".  Because the "child.example.com" zone is a
      descendant of the "example.com" zone, the glue records are in-
      bailiwick.

    * Out-of-bailiwick:  The antonym of in-bailiwick.'
* They use ns1.bitfl1p.com and ns2.bitfl1p.com as name server for all of the flipped domains. Based on the above definition, bitfl1p.com would not be In-Bailiwick for jquesy.com or cloudfront.net
* Can not rely on this statistic but it may give a good intuition: 
  * Bitfl1p.com project received 16k dns requests per 10 domain every day(on average)
  * We got 114k dns requests for our 10 domain every day(on average)

### Our method:
* We used two different in bailiwick nameserver for each of our domain
* We set it up in godaddy.com which is our domain name provider
* In godaddy.com, for the flipped domain example.com, we defined the following:
  * ns1.example.com (IP address: `192.35.222.16` @ `frears`)
  * ns2.example.com ((IP address: `192.35.222.17` @ `frears`))
* **How to do it?**
  * Go to manage dns
  * Use `...` beside `Add` button 
  * Choose `Host Names` from the drop down 
  * Click on `Add`
  * choose ns1/ns2 as host name(it would automatically add the domain name to it)
  * Choose `192.35.222.16` for `ns1` and `192.35.222.17` for `ns2`
    * These two IP addresses are public IPs of `frears` which is in unfiltered network
    * `frears` is the DNS server for all of the domains
  * After adding those hosts, go to `nameservers` section
  * Add ns1.example.com and ns2.example.com as the nameservers for example.com
  * It would take up to 48 hours till these changes take effect
### Note on In-bailiwick:
* since ns1.example.com is subdomain of example.com, we would have circular dependency on resolving these names:
  1. To reach example.com, we need to get its address from ns1.example.com
  2. To reach ns1.example.com, we first need to resolve example.com
  3. To eliminate this circular dependency, we use glue records
* Glue record:
  * The domain name provider(godaddy.com) send the glue records to the TLD DNS
    * ns1.example.com is @ 192.35.222.16
    * ns2.example.com is @ 192.35.222.17
  * In this way, the resolver can directly find the IPs of ns1.example.com and ns2.example.com without first finding the IP of example.com

## Identifying source of flips

# Principal Questions and TODO:
## Setting TTL:
- [x] Set the TTL of `A` records as small as several seconds in order to avoid poisoning the internet.
- [ ] Set the TTL of `A` records in a way that we can experiment caching on DNS resolvers:
  * Due to RPZ setup, we can not query every resolver we want
  * Instead, we would look for DNS queries coming from UCSB campus network
  * After receiving those requests, we would immediately query campus DNS resolver for the corresponding flipped domain
  * If the resolver sends another query for the same domain to our authoritative DNS server, it implies that it has not cached our response
  * Otherwise, if the resolver responds to the second query( sent by us) without querying the authoritative DNS server of flipped domain, it implies that it has cached the previous response
- [ ] What should I set for the TTL of `ns` records?
  * Bitfl1p.com used TTL=1 for `A` records and max value for `ns`  records  
- [x] Having high bandwidth for capturing the traffic in case of a spike in the incoming traffic( both for DNS and honeypot)
## DNS responses format:
* Analysis of Bitfl1p.com response format is [here](https://docs.google.com/spreadsheets/d/1BkacV2QkWZHvcwsO6rS5Pcx6IU3dWbAh7pXX5TO-2wA/edit?usp=sharing). It also shows how it should bt based on the Def Con talk and comments in the code
- [ ] Why bitfl1p.com project has a section in DNS server to respond the control domain queries? it responds with `ns1.config.control+"."` and `ns2.config.control + "."`
  - [ ] Should I do the same or completely avoid it?
  - [ ] What is the effect of this part on the dataset?
- [ ] Double check and record all types of responses sent back by `bf-dns.go` from bitfl1p.com
- [ ] In the respons to `b.com`, when bitfl1p.com project is sending back `a.com` is at `x`, it sets the `qname` in the query and answer section of response equal to `a.com` which may result in being dropped by the client.
   * In my own experiment when the client directly queried a DSN server for `b.com` and I responded back `a.com` is at x and the answer was accepted, I set the `qname=b.com` in query section and `qname=a.com` in answer section of the response. (Make sure to check the canonical name issue for this)
## Scalibility of hone pot and DNS server
- [ ] Is the honey pot able to respond to large number of connections?
  - [ ] Check hardware resources( eg RAM is essential for threading and having many open connections)
  - [ ] How to prevent open connection and syn attacks?
  - [ ] Is `nginx` powerful enough to respond to many requests? 
  - [ ] What is the effect of web cache proxy servers in our experiment? How to eliminate it?
  - [ ] Saving the data in NFS rather than memory due to shortage of space specially in case of spike 
- [ ] Regular back up of log files

  


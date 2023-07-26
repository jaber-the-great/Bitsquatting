## Next week to do list:
- [ ] Set up root and TLD DNS server:
  - [x]  Look for existing
  - [ ]  Send and receive req/ans
  - [ ]  Design our own root DNS server etc
  - [ ]  Inject flips
- [ ]  Change the DNS server:
  - [x]  TTL = 1 (both A and ns records)
  - [ ]  Send third type of response 
  - [ ]  I still did not get the answer: Two response in one packet or using different packets 
  - [ ] Control domain stuff
- [ ] Install a light weight honey pot
- [x] Use the NFS on `della` and `stanley`: Increased memory instead of NFS
- [ ] Change the response of web server
- [ ] Talk about https for finger printing 
- [ ] Check chromium source code
- [ ] Check finger printing with Pouneh
- [x] Build chromium on a new client  



## To do and search: 
- [ ] Look into the DNS lookup source codes and see where it checks for equality of DNS query ID and the ID in the response. Then, check whether it updates any table or data record after that or not. Mainly due to the problem in the ping verbose in the following scenario:
  * Ping jaberxyz.com  : pinging b.com at 192.168.50.40 
- [ ]  Forcing/ injecting bitflips in different part of the network for DNS resolution
  - [ ]  Lab environment
  - [ ]  Cloud environment
- [ ]  How to distinguish bots from legit users? specially the distributed botnets: Notes from Penn state paper
- [ ]  Health check of route53 in amazon, can it be related
- [ ]  Checking the cache of DNS resolver if the request is from campus or the resolver is responding
- [ ]  




## Bridge
- [ ]  Finish Dawn Song SoK and the Stanford paper sent by Ilya 
- [ ]  Check the behavior of wardens:
  - [ ]  Why they reveal the unencrypted data
  - [ ]  MPC: What type? ASS, Shamir etc
  - [ ]  Learn solidity
  - [ ]  
  















## The DNS Deep Dive to do
- [ ] 1.2 M domains exposed to two DoS exploits available in MetaSploit(check for that)
- [ ] Bitflip in IO(can be a source but so rare in clinet side)
- [ ] learn a little bit rust and R
- [ ] Working with PyShark and TCP dump

## Bind DNS
- [ ] Why the result is not cached on the local device:
  - [ ] Check caching for legit DSN replies over internet
  - [ ] Check caching for local bind replies(using another machine as resolver)
  - [ ] Caching with my own answer(Scapy)
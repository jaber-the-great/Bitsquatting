#! /usr/bin/env python3
from scapy.all import *
dns_req = IP(dst='192.168.56.30')/UDP(dport=53)/DNS(rd=1, qd=DNSQR(qname='jaberxyz.com'))
answer = sr1(dns_req, verbose=1)

print(answer[DNS].summary())

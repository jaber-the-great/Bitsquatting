# coding: utf-8
conf.iface
conf.iface = "vboxnet0"
pkt = sniff(count =5, iface='enp0s25',filter = 'port 53')
type(pkt)
pkt.summary()
pkt[1].show()
explore()
explore(scapy.layers.dns)
lsc()
ls()
ls(ARP)
get_ipython().run_line_magic('save', '-a new_sate ~0/')
arppkt = Ether()/ARP()
arppkt[ARP].hwsrc = "00:11:22:aa:bb:cc"
arppkt[ARP].pdst = "192.168.56.30"
arppkt[Ether].dst = "ff:ff:ff:ff:ff:ff"
arppkt
dnspkt = IP()/UDP()/DNS()
dnspkt
newdns = IP()
newdns
newdns = newdns/UDP()/DNS()
newdns
sendp(arppkt,iface = "vboxnet0")
# send p is for sending at layer 2
get_ipython().run_line_magic('save', 'alaki ~0/')
# send is for sending regular packet 
# send or sendp(pkt,inter=1,loop=0,iface ): inter is interval between two packets
# loos sends the packet endlessly if not set to 0
pkt[0][Ether]
pkt[0][DNS]
pkt[0][DNS][DNSQR]
pkt[0][DNS][DNSQR].qname
pkt[0].command()
# command() gives the string representation of a packet. you can use 
# sendp("the result of command()") to send a packet 
for packet in pkt:
    if (packet.haslayer(ICMP)):
        print(f"ICMP code: {packet.getlayer(ICMP).code}")
        
# sr() and srp() when you want to send packet and expect a response back
# sr1() and srp1() is just for when you expect/want to get 1 response
sr1(IP(dst="192.168.56.30")/ICMP())
sr(IP(dst="192.168.56.30")/ICMP())
# srloop and srploop: continue to resend the packet after receiving each response
srloop(IP(dst="192.168.56.30")/ICMP(), count = 5)
# to write a python script
from scapy.all import *
from scapy.all import sr1,IP,ICMP
# The python script is not as verbose as the interactive mode
# you would need to pring everythong you want
# IMPORTANT: using prn argument to cusotmize sniff output and to send the responses
def arp_display(pkt):
    if pkt[ARP].op == 1:  # who-has (request)
        return f"Request: {pkt[ARP].psrc} is asking about {pkt[ARP].pdst}"
    if pkt[ARP].op == 2:  # is-at (response)
        return f"*Response: {pkt[ARP].hwsrc} has address {pkt[ARP].psrc}"

sniff(prn=arp_display, filter="arp", store=0, count=10)
# to pass several argument to the prn, you should use nested functions. 
# look into the link for more info: https://thepacketgeek.com/scapy/sniffing-custom-actions/part-2/
# When creatin a message like dns request, we can use random function for src port numebr
# we can define a template packet and extend other packets from that template
# You can define filter as a file and then apply it to sniff function 
myFilter = f"udp port 53 and ip dst 192.168.56.1"
sniff(filter = myFilter, count =2)
get_ipython().run_line_magic('save', 'tillNow ~0/')
# making a DNS query
ans = sr1(IP(dst="8.8.8.8")/UDP(sport=RandShort(), dport=53)/DNS(rd=1,qd=DNSQR(qname="7a645c14a2eaac.d.requestbin.net",qtype="A")))
ans[0]
%save scapy_notes_commands ~0/
get_ipython().run_line_magic('save', 'scapy_notes_commands ~0/')

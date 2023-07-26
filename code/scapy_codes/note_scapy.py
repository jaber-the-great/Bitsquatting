## Notes and command of scapy
## The follwoing line writes the captured packet into a file 
for pkt in foo:
...:     wrpcap("foo_legit_lookup.pcap", pkt, append=True)

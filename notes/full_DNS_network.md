## Implementation of full DNS hierarchy:
### Root DNS server: 
* [Follow these steps](https://flylib.com/books/en/2.684.1/setting_up_a_root_name_server.html)
### TLD server:
* [Follow these steps](https://www.oreilly.com/library/view/wireless-hacks/0596005598/ch04s15.html)
### Authoritative server:
* [Follow these steps](https://www.linuxbabe.com/ubuntu/set-up-authoritative-dns-server-ubuntu-18-04-bind9)
### DNS resolver:
* [Follow these steps](https://www.linuxbabe.com/ubuntu/set-up-local-dns-resolver-ubuntu-20-04-bind9)
### Client machine(Chromium):
* [Building chromium](https://chromium.googlesource.com/chromium/src/+/master/docs/linux/build_instructions.md#Install)
* since many of us use latest version of Ubuntu LTS(22.04 right now); Keep in mind that  you can not build chromium on 22.04 and the latest supported version is 20.04. While it supports 16.04, you may face some issues regarding to the version of some libraries. It's better to go with 20.0


## TODO:
* The network configuration is based on Network_conf.md in the same directory
- [ ] Put all of the steps together as a clean tutorial
- [ ] Discuss running root and TLD server with our own code rather than bind9
- [ ] Wireshark trace for all of of the steps
- [ ] Injecting flips:
  - [x] In client
  - [x] After resolver 
    
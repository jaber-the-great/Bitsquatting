## Finding the available neighbouring domains of top websites
### bf-lookup.go
* The source code which finds the neighbouring domains of a given domain and checks whether the domain is already registered or not
* From project bitfl1p.com
* It has some limitations for being used with NPM packages and debian packages. For example, NPM packages are case sensitive but domain names are not.


### Top-10k.csv file
* Alexa discontinued working as a website ranknig system
* I got the latest version of top 1M websites. 
* Extracted the top 10k websites
* Found the available neighbouring domains of these websites (went till ~7.5k)


### Batch_fb_lookup.sh
* This script goes through a given file(here top-10k.csv) and read the domain names one by one. 
* It runs the bf-lookup:
  * complied version: ./bf-lookup \<domain name>
  * src version: go bf-lookup.go \<domain name>

### Top-10k-respones.txt
* The result of executing Find.sh for about 7500 domains. The * means the domain is not alreay taken.


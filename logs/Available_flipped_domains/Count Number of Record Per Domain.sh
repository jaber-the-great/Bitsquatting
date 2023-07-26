# IMPORTANT: I did it in a better way using a pythin script. I used this script to 
# find the best domains to purchase for our experiment. 
# Requirement: log file (anon-access.json, dns.json etc) from bitflip.com project be in the 
# same directory, or another log file which we want to analyse(edit the name of the file or script)
# Requirement: A file containing the domains used in bitflip.com project given
# as input to this script(of the file containing the list of domains for the experiment).
# Currently it is : "Domain Used in Bitfl1p Project.txt"
# This script finds how many requests each domain has in bitfl1p.com project log file.
# It can be used to the same thing for analysing other log files


#!/bin/bash
input="$1"
cnt=0
var=".txt"
while IFS= read -r line
do
  let cnt=cnt+1
	if test $cnt -gt 0
then
echo $cnt	"$line"

stripped=${line::-1}
count=$(grep -c $stripped anon-access.json)
echo "$line" >> result.txt
echo "$count" >> result.txt
echo "count is: " "$count" 
echo "#############################"
fi
done < "$input"

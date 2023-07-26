#this is the code which iterates through the given file, finds the neighbouring domain names and uses whois
#The file contating the names of domains is the input to this snippet 
#!/bin/bash
input="$1"
cnt=0
while IFS= read -r line
do
  let cnt=cnt+1
	if test $cnt -gt 6780
then
echo $cnt	"$line"
./bf-lookup "$line"
echo "#############################"
fi
done < "$input"

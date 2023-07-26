# Requirenment: executable bf-lookup file in the same directory or if the 
# bf-lookup is the source code, either building the executable file from that
# or compiling and running the source code
#this is the code which iterates through the given file, finds the neighbouring
#domain names and uses whois to check whether the domain is already registerred or not
#The file contating the names of domains is the input to this snippet 
#!/bin/bash
input="$1"
cnt=0
while IFS= read -r line
do
  let cnt=cnt+1
	if test $cnt -gt 0
then
echo $cnt	"$line"
./bf-lookup "$line"
echo "#############################"
fi
done < "$input"

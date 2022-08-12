# go-file-regex
 
use a capture group in RegEx for what you want to list

 ``````bash
#  find all 10 to 20 letter words in moby dick
./go-file-regex -f moby.txt -r '\W([\w]{10,20})\W' -v

# find all digits and find frequencies for each
./go-file-regex -f pi.txt -r '(\d)' -v

# find all digits and find frequencies for each
./go-file-regex -f pi.txt -r '(\d)' -v


# find all first digits and find frequencies for each
./go-file-regex -f pops.json -r '\:\W+?(\d)' -s

# list world populations (w/o country/ region)
./go-file-regex -f pops.json -r '"Year": "2010",\W+"Value": "(\d+)"' -s
 ``````
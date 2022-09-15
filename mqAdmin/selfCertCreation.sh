cd /mnt/mqm
mkdir -p certs
cd certs


#Server create and extract
runmqakm -keydb -create -db server.kdb -type cms -pw passw0rd -stash
runmqakm -cert -create -db server.kdb -pw passw0rd -label ibmwebspheremqqm1 -dn "CN=QM1,O=IBM,C=USA" -size 2048
runmqakm -cert -extract -db server.kdb -pw passw0rd -label ibmwebspheremqqm1 -target server.crt -format ascii

runmqakm -keydb -create -db client.kdb -type cms -pw passw0rd -stash
runmqakm -cert -create -db client.kdb -pw passw0rd -label ibmwebspheremqapp -dn "CN=XXXX,O=IBM,C=USA" -size 2048
runmqakm -cert -extract -db client.kdb -pw passw0rd -label ibmwebspheremqapp`` -target client.crt -format ascii


#Add client cert to server
runmqakm -cert -add -db server.kdb -pw passw0rd -label ibmwebspheremqapp -file client.crt -format ascii

#Add server cert to client 
runmqakm -cert -add -db client.kdb -pw passw0rd -label ibmwebspheremqqm1 -file server.crt -format ascii

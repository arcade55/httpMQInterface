FROM ibmcom/mq:latest


USER mqmadmin

RUN runmqckm -keydb -create -db "/var/mqm/qmgrs/QM1/ssl/QM1.kdb" -pw serverpass -type cms -stash
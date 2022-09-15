# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.18, 1.17, 1-bullseye, 1.18-bullseye, 1.17-bullseye, 1-buster, 1.18-buster, 1.17-buster
ARG VARIANT=1-bullseye
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}
RUN mkdir -p /opt/mqm && chmod a+rx /opt/mqm 
 # Location of the downloadable MQ client package \
ENV RDURL="https://public.dhe.ibm.com/ibmdl/export/pub/software/websphere/messaging/mqdev/redist" \
    RDTAR="IBM-MQC-Redist-LinuxX64.tar.gz" \
    VRMF=9.3.0.0
 # Install the MQ client from the Redistributable package. This also contains the
 # header files we need to compile against. Setup the subset of the package
 # we are going to keep - the genmqpkg.sh script removes unneeded parts
 ENV genmqpkg_incnls=1 \
     genmqpkg_incsdk=1 \
     genmqpkg_inctls=1
 RUN cd /opt/mqm \
   && curl -LO "$RDURL/$VRMF-$RDTAR" \
   && tar -zxf ./*.tar.gz \
   && rm -f ./*.tar.gz \
   && bin/genmqpkg.sh -b /opt/mqm



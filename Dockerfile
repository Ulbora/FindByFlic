FROM ubuntu:16.04

RUN apt-get update  
RUN apt-get install -y ca-certificates
ADD main /main
ADD entrypoint.sh /entrypoint.sh
ADD cert.pem /cert.pem
ADD key.pem /key.pem
ADD static /static
WORKDIR /

EXPOSE 8070
ENTRYPOINT ["/entrypoint.sh"]
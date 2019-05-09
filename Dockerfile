FROM centos:7

COPY server /server

ENTRYPOINT /server

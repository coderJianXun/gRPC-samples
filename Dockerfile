FROM centos
RUN yum install golang -y \
    && yum install dlv -y \
    && yum install binutils -y \
    && yum install vim -y \
    && yum install gdb -y \

#docker build -t test .
#docker run -it --rm test bash
FROM ubuntu

MAINTAINER lwh lwhile521@gmail.com

# RUN apt-get update

COPY ./datavolume /home/datavolume/

# Install Java SDK
RUN ["mkdir", "/usr/lib/jvm/"]
WORKDIR /home/datavolume/
RUN ["tar", "-zxvf","jdk-8u101-linux-x64.tar.gz","-C","/usr/lib/jvm/"]
ENV JAVA_HOME /usr/lib/jvm/jdk1.8.0_101
ENV JRE_HOME &{JAVA_HOME}/jre
ENV CLASSPATH .:${JAVA_HOME}/lib:${JRE_HOME}/lib  
ENV PATH ${JAVA_HOME}/bin:$PATH  

# Install Tomcat8
RUN ["tar", "-zxvf", "apache-tomcat-8.5.5.tar.gz", "-C", "/usr/lib/"]
RUN ["sh", "/usr/lib/apache-tomcat-8.5.5/bin/startup.sh"]

RUN apt-get update
RUN apt-get install curl -y

EXPOSE 65530

FROM ubuntu
EXPOSE 9092
RUN apt-get -y update
RUN apt-get -y install wget
RUN apt-get -y install default-jre
#RUN apt-get -y install zookeeperd
RUN wget https://www-eu.apache.org/dist/kafka/2.2.0/kafka_2.11-2.2.0.tgz && tar xvf kafka_2.11-2.2.0.tgz
RUN rm kafka_2.11-2.2.0.tgz
COPY startKafka.sh startKafka.sh
COPY server.properties ./kafka_2.11-2.2.0/config/server.properties
CMD ["/bin/bash", "/startKafka.sh", "&", "/bin/bash", "&"]


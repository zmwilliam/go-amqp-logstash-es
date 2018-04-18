FROM docker.elastic.co/logstash/logstash:6.2.3

RUN logstash-plugin install logstash-input-rabbitmq

COPY config/logstash.yml /usr/share/logstash/config
COPY pipeline /usr/share/logstash/pipeline

ENV LS_JAVA_OPTS="-Xms128m -Xmx512m"
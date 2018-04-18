
## Golang + Rabbit + Logstash + Elasticsearch

Simple example of a Golang application inserting into Rabbitmq queue, and Logstash consuming the queue and saving to Elasticsearch.

## Running

#### Start Services
  
```

$ docker-compose -f docker/docker-compose.yml up --build -d

```

#### Run Project

```

$ dep ensure

$ go run main.go

```

Now enter the text that will be queued and sent to elasticsearch

#### GET results from Elasticsearch

```

$ curl -X GET http://localhost:9200/applogs/_search

```
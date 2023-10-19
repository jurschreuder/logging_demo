The Ultimate Logger
===

Example for Go and FastAPI (todo now only Go) logging.

# The Fantasy Logging Team

In Golang we use Zap, by Uber as the logger.
It's fast and has a lot of options to create for example json log entries with almost no memory allocations.

[https://github.com/uber-go/zap](here is the github link)

We will write the logs to `/var/log/go_example` (todo now autodiscovers containers)

To ingest the logs and send it to ElasticSearch we use vector (todo now regular old filebeat and logstash).
Vector is a mega fast FileBeat alternative written in Rust.
To install it (auto detects platform):

```
curl --proto '=https' --tlsv1.2 -sSf https://sh.vector.dev | bash
```

The logs are collected in ElasticSearch. This gives you some fancy search options.
To visualize the logs we use Kibana.
This is how you set it all up using Docker:

[https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html](Install ElasticSearch and Kibana)



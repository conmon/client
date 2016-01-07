ConMon Client
=============

Container Monitoring Client

## Purpose ##
ConMon was born because there are not that many end-to-end monitoring solutions for
docker and containers. ConMon client (this component) is responsible for collecting
statistics about current machine and then sending it onwards. It also has an internal
API if you with to extract data yourself from the client.

This project is heavily inspired by Google's [cAdvisor](https://github.com/google/cadvisor). Some of the code might be looking similar (or identical), but that should change over time. Licences and copyright notices would also be added soon

---
# WARNING #
Many features are not yet complete and/or have full functionality

---

ConMon (will) support a range of container technologies including
1. WIP [docker](https://www.docker.com/)
2. TODO [rkt](https://coreos.com/rkt/docs/latest/)
3. TODO [LXC](https://linuxcontainers.org/)

It also allows to export to your favorite database for storing information
1. WIP [ConMon Server](https://github.com/conmon/server)
2. TODO [InfluxDB](https://influxdata.com/time-series-platform/influxdb/)
3. TODO [BigQuery](https://https://cloud.google.com/bigquery/)
4. TODO [ElasticSearch](https://www.elastic.co/products/elasticsearch)
5. TODO [Redis](http://redis.io/)

## Building ##
Coming soon...

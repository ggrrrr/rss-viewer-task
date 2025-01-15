# README

* Take your rss aggregator service take not of the host and port on which it is listening.

* Generators

```sh

# Models for RSS sources
bin/rails generate model  Address url:string

bin/rails generate model rss_item \
    title:text link:text \
    source:text \
    source_url:text \
    description:text \
    publish_date:datetime

rails generate controller sources index delete create

rails generate controller home index

```

# README

This README would normally document whatever steps are necessary to get the
application up and running.

Things you may want to cover:

* Generators

```bash

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

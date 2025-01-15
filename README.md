# RSS Viewer task

## Subtasks

* Create a Golang RSS Reader package, which can parse asynchronous multiple RSS feeds
* Create a Golang Rss Reader service that uses the RSS package
* Create a Rails application

## Code structure

* [be](/be) -- backend services and libraries in GoLang
  * [pkg/rssclient](/be/pkg/rssclient) -- RSS client library which aggregates data
  * [pkg/common](/be/pkg/common) -- help libraries shared between services
    * [auth](./be//pkg/common/auth) -- authorization help tools
    * [system](/be/pkg/common/system/) -- general backend service help tools, start web server/router, adds middleware for logging,auth,CORS etc, handler shutdown,startup
    * [system](/be/pkg/common/web/) -- web helper functions for handling web traffic
  * [svc/rssaggregator](/be/svc/rssaggregator) -- RSS aggregator service
    * [rest](/be/svc/rssaggregator/intternal/rest/) -- handle api calls
    * [app](/be/svc/rssaggregator/intternal/app/) -- actual application layer ( uses pkg/rssclient lib)
    * [cmd](/be/svc/rssaggregator/cmd) -- main entry point for the service
* [rssviewer](/rssviewer/README.md) -- Ruby on Rails application
  * [Home](/rssviewer/app/controllers/home_controller.rb) -- filter and generates RSS home page
  * [RssSources](/rssviewer/app/controllers/rss_sources_controller.rb) -- Handles CRUD for RSS sources list of URLs
  * [RssJob](/rssviewer/app/jobs/rss_job.rb) -- handler fetching data from rss aggregator
  * [Home templates](/rssviewer/app/views/home/) -- Home page template
  * [RssSources templates](/rssviewer/app/views/rss_sources/) -- Templates for RssSources CRUD operations
* [Makefile](/Makefile) -- Make file for automation and CiCd

## Howto start all services

### RSS Reader Service

* Run tests and linter

    ```sh
    make go_lint
    make go_test
    ```

* Build docker image

    ```sh
    make build_svc
    ```

* Run docker

  ```bash
  docker compose up -d rss
  ```

### RAIL APP

* Take your rss aggregator service take not of the host and port on which it is listening.

* create file with `.env` with the following content

    ```sh
    export RSS_URL=http://<HOST:PORT>/v1/parse
    export RSS_JWT=<JWT_TOKEN>
    ```

* From console run the following:

    ```bash
    # Source the env variables
    source .env

    # install all dependencies 
    ./bin/bundler install
    # create DB schema to the configured target database ( by default this is SQLite)
    ./bin/rails db:migrate
    # Run the application
    ./bin/rails server
    # access the page with your browser default is: http://localhost:3000

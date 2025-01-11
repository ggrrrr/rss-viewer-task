# RSS Aggregator service

## API endpoints

* `/v1/parse`
  * method `POST`
  * request payload
    * Format JSON

    ```JSON
    {
        "urls": [
            "url1",
            "url2"
        ]
    }
    ```

    * response
      * Format JSON

      ```JSON
      {
        "items":[
            {
                "title":"Title ",
                "source":"Source title",
                "source_url":"Source link",
                "link":"link",
                "publish_date":"2019-01-23T01:15:00Z",
                "description":"description"
            },
        ]
      }

      ```

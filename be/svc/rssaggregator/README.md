# RSS Aggregator service

## API endpoints

* `/v1/parse`
  * Example curl usage:

    ```bash

    curl -XPOST -H'Authorization: Bearer JWT_TOKEN' -d'{"urls":["asdurl"]}' localhost:8080/v1/parse

    ```

  * Authorisation: HTTP header must be set with JWT token: `Authorization: Bearer <SOME_JWT_TOKEN>`
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

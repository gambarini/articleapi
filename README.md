# Article API

The first endpoint, `POST /articles` handle the receipt of some article data in json format, and store it within the service. The article
will get an id and the date will be the datetime the article was received.

POST Payload
```
    {
          "title": "latest science shows that potato chips are better for you than sugar",
          "body" : "some text, potentially containing simple markup about how potato chips are great",
          "tags" : ["health", "fitness", "science"]
    }
```

The second endpoint `GET /articles/{id}` return the JSON representation of the article.

The final endpoint, `GET /tags/{tagName}/{date}` return the list of articles that have that tag name on the given date and some summary data about that tag for that day.

## Running

Download:

```
    $ go get github.com/gambarini/articleapi
```

You must have the following installed to run the API server:

- golang 1.10 (but any older version should be ok)
- docker

From the articleapi folder execute the run script:

```
    $ run.sh
```

It will take a while to initialize the following:

- MongoDB docker container.
- Start the api server

If it all work, the API should be running on localhost:8000

Just Ctrl+C to terminate the server when you done. The script will cleanup the
containers.

## Considerations

Dates are always in the format `yyyy-mm-dd`

The `POST /articles` endpoint ignores any ID and Date posted in the payload. Both values are
generated by the API based on the time the article is received.

The `GET /tags/{tagName}/{date}` endpoint is not considering high traffic situations. A better
approach for scalling the tag queries would be building it as a separeted service with its own
tags database. Then the tag collection is populated every time a new article is created (services choreograph),
resulting in a eventual consistent tags database with better throughput.

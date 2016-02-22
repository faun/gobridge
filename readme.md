# GoBridge Example Code

GoBridge example code from https://gobridge.gitbooks.io/building-web-apps-with-go/content/en/

### Examples

- [Static File Server](../../tree/static-files)
- [Markdown Generator](../../tree/markdown-generator)
- [Deployment](../../tree/deployment)
- [Routing](../../tree/routing)
- [Middleware](../../tree/middleware)
- [JSON Serialization](../../tree/json)
- [HTML Templates](../../tree/templates)
- [HTML Templates w/Render](../../tree/render)
- [Unit Testing](../../tree/unit-tests)
- [End to End Testing](../../tree/integration-tests)
- [Controllers](../../tree/controllers)
- [Databases](../../tree/databases)

To run these examples, run `go get` followed by: `go run main.go`.

To create a new book, run:

    curl -vX POST http://localhost:8080/books \
    --header "Content-Type: application/json" \
    -d "{\"author\": \"George Orwell\", \"title\": \"1984\"}"

To create another:

    curl -vX POST http://localhost:8080/books \
    --header "Content-Type: application/json" \
    -d "{\"author\": \"Victor Hugo\", \"title\": \"Les Misérables\"}"

To view all books:

    curl -v http://localhost:8080/books --header "Content-Type: application/json"

Or open [http://localhost:8080/books](http://localhost:8080/books) in a browser

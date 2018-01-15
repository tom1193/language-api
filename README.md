# language-api
Entity sentiment anlysis and bing image search.

### Getting started
This app uses [GCP Natural Language API](https://cloud.google.com/natural-language/). You'll need an API key to run your own version of this app. Place API key file in `creds/` for Docker to fetch it.

Docker is used to run the app in a container. Have Docker installed and running before the next steps. To build the image, run the following in the project directory:
```
docker build -t language-api .
```
Run `docker images` to check if the image has been built. The following launches a container off that image and deploys the app at http://localhost:8080/
```
docker run --publish 8080:8080 --name build --rm language-api
```

### Requests

The API currently only accepts GET requests at http://localhost:8080/entity
Submit input text as query parameters in the format of:
```
{text: "Your input text."}
```
### Response

On a successful response the API returns entities, sentiments and corresponding images in the format `[]Entity` where `Entity` is in the format.
```
type Entity struct {
Name string `json:"name,omitempty"`
Order int32 `json:"order,omitempty"`
Sentiment float32 `json:"sentiment,omitempty"`
Images []Image `json:"images,omitempty"`
}
type Image struct {
Iid string `json:"iid,omitempty"`
Url string `json:"url,omitempty"`
}
```
Returns `null` if no entities found.
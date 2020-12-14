# Recommendation Service

This service provides creates a recommendation list of subjects based on the users preferences.

![Poster Recommendation service](./public/poster.png)

## Requirements

`Python 3.8`
Libraries:

- `numpy`
- `pandas`
- `flask`
- `gunicorngs`

## Docker deployment

Build image `docker build  .`

Run the container `docker run   -p 2000:2000 <Image>`

### Endpoints and methods

All information about endpoints and http methods you will find in [Endpoints.md](docs/endpoints/recommendation.md)

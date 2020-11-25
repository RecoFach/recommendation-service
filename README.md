# RecoFach-Recommendation

## Basic functionality

Build image `docker build  .`

Run the container `docker run   -p 2000:2000 <Image>`

Use curl to send `POST` request f.e `curl --location --request POST 'http://0.0.0.0:2000/recommend?Software%20engineering=0&AI=0&Low-level=0&Security=1&Web=0&Theoretical=0' \
--header ': '`

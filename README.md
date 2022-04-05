# Simple Go Actions

1. Created a simple forecasting API using moving averages.
2. Using Dockerfile created a compact Go app. It uses [scratch](https://support.snyk.io/hc/en-us/articles/360004012857-What-are-docker-scratch-based-images-) base image, so it is small in size.
3. Using the Github Actions it is automatically tested and built.
4. Using the Github Actions the changes on `main` branch are published on the API endpoint, that is powered by [Google Cloud Run](https://cloud.google.com/run).

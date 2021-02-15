FROM alpine
WORKDIR /usr/bin
RUN curl -L https://install.meilisearch.com | sh
COPY cmd/timbermill /usr/bin

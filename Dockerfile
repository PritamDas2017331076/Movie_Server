FROM golang:latest
WORKDIR app
ARG LOG_DIR=/app/logs
RUN mkdir -p ${LOG_DIR}
ENV LOG_FILE_LOCATION=${LOG_DIR}/app.log
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install
EXPOSE 8080
VOLUME [${LOG_DIR}]
CMD ["Movie_Server" ,"start"]
#  docker build -t pds99/movie_server:0.0.1 .
# docker push pds99/movie_server:0.0.1
# docker run -d -p 8080:8080 -v ~/logs/go-docker:/app/logs pds99/movie_server:0.0.1
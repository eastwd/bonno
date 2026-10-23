# Welcome to Bonno

## Getting Started

### Requirements

- Go 1.27+
- Revel v1.1 (`github.com/revel/revel`)

### Run locally (without Docker)

    $ (cd tools/revel-cmd && go install github.com/revel/cmd/revel)
    $ DB_DRIVER=sqlite3 DB_DIRECTORY=./conf/app.db revel run -a .

### Start the web server

    $ git clone https://github.com/bandoshintaro/bonno.git
    $ cd bonno
    $ docker-compose up -d

#### store data

    $ docker run -it -v .conf:/app/src/bonno/conf bando/bonno

#### use own movie data

    $ docker run -it -v /your/own/movie/dir:/app/src/bonno/public/douga bando/bonno

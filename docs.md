docs

$ docker build . -t my-app
$ docker run -p 8080:8080 -it --name my-app my-app

$ docker pull ejaycoleman/pusher-dwarf-facade
$ docker run -p 8080:8080 -it --name ejaycoleman/pusher-dwarf-facade
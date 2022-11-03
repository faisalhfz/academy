job "faisal-echo" {
  datacenters = ["dc1"]
  type = "service"

  group "web" {
    count = 1

    network {
      port "http" {
        to = 1323
      }
    }

    task "faisal-echo" {
      driver = "docker"

      config {
        image = "faisalhfz/go-echo-faisal:v1"
        ports = ["http"]
      }

      resources {
        cpu    = 256
        memory = 256
      }
    }

    service {
      name = "faisal-echo"
      port = "http"
      tags = [
        "traefik.enable=true",
        "traefik.http.routers.faisal-echo-demo.rule=Host(\"faisalecho.cupang.efishery.ai\")",
      ]
      check {
        port        = "http"
        type        = "tcp"
        interval    = "15s"
        timeout     = "14s"
      }
    }

  }
}
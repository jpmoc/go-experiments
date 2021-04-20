use "docker" {
  "image" = 123
  "tag" = "local"
}

use skaffold {
  "image" = 123
  "tag" = "local"
  "kaniko" = true
}

"endpoints" = {
  "access_key" = "secret"

  "location" = "online"
  "middle" {
      key01 = "value01"
      key02 = "value02"
  }
  "middle" {
      key11 = "value11"
      key12 = "value12"
  }

  "url" = "http://..."
}

terraform {
  required_providers {
    redis-config = {
      versions = ["0.2"]
      source = "github.com/justsimplify/redis-config"
    }
  }
}

# Following block uses action resource
##############
resource "redis-config" "my-redis-object-1" {
  key   = "k2"
  value = "random value 2"
}
##############


# Following block used data resources
##############
data "redis-config" "redis_k2" {
  depends_on = [redis-config.my-redis-object-1]
  key = "k2"
}

# Returns value of k2
output "k2_conf" {
  value = data.redis-config.redis_k2
}
##############
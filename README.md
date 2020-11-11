### How to
---

Run the following

```bash
# Go to root directory and run
# This creates the terraform provider and puts in the directory 
# ~/.terraform.d/plugins/github.com/justsimplify/redis-config/0.2/darwin_amd64
$ make install

# Navigate to example directory and run
$ make setup
```

- Last command runs and puts redis key and value. Key is `k2` and value is `random value 1`. (Refer `main.tf` file)
- After success, it tries to fetch data and prints the object from `tfstate` as stored. (Using data resource feature)

#### Assumptions
- Redis is up and running and has following config. 
  - host: `0.0.0.0`
  - port: `6379`
  - password: `<empty password>`
- If the above config is different, change those values in `variables.tfvars` file before proceeding.
- If not using `variables.tfvars`, we can also use the following env variables.
  - host: `REDIS_HOST`
  - port: `REDIS_PORT`
  - password: `REDIS_PASSWORD`
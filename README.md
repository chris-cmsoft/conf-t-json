# Conf to Json

Simply converts `.conf` files to json. 

Currently being used to convert configurations such as SSH and Nginx to json, to be validated by Rego policies.

```shell
cat >nginx.conf <<EOL
server {
    location / {
        root /data/www;
    }

    location /images/ {
        root /data;
    }
}
EOL

conf-to-json nginx.conf
```

## Resulting conversions

```none
port 80 -> "port": ["80"]
port -> "port": []
port 80 8080 -> "port": ["80", "8080"]

# Appends multiples
port 80
port 8080
-> "port": ["80", "8080"]
```
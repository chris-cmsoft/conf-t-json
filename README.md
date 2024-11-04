# Conf to Json

Simply converts `.conf` files to json. 

Currently being used to convert configurations such as SSH and Nginx to json, to be validated by Rego policies.

## Usage

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

# From a file
conf-to-json -f nginx.conf
# OR from stdin
cat nginx.conf | conf-to-json
```

## Resulting conversions

All values are output as arrays of values, to standardise single and multiple values. 

```none
`port 80` -> "port": ["80"]
`port` -> "port": []
`port 80 8080` -> "port": ["80", "8080"]

# Appends multiples
`
port 80
port 8080
`
-> "port": ["80", "8080"]
```

# envoy_h2_to_h2c_example
Example of using envoy to convert h2 traffic to h2c

## Installation

1. Install Envoy: https://www.envoyproxy.io/docs/envoy/latest/start/install#install
2. Install golang: https://golang.org/doc/install
3. (Optional) Install curl with http2 support: https://curl.se/docs/http2.html
4. 🎉

## Running

1. `./start.sh` <- This will build and start the h2c app and envoy proxy
1. `curl https://0.0.0.0:61001 --cert ./client_certs/client.crt --key ./client_certs/client.key --cacert ./client_certs/ca.crt -k -v`
1.  See that it uses h2 🎉

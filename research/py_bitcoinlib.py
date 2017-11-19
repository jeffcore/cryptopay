from bitcoin.rpc import Proxy

p = Proxy(service_url="http://bitcoinrpc:pass@192.168.25.17", service_port=8332, btc_conf_file="bitcoin.conf")
info = p.getinfo()

import requests

class Node(object):
    def __init__(self, url=None, port=8332, username=None, password=None):
        self.url = url
        self.port = port
        self.__username = username
        self.__password = password
        self.full_url = self.url + ":" + str(self.port)
        print("New Node Is Born")

    def __str__(self):
        rep = "Node object\n"
        rep += "url: " + str(self.url) + "\n"
        rep += "port: " + str(self.port) + "\n"
        rep += "url: " + str(self.__username) + "\n"
        rep += "password: ******"
        return rep

    def makeCall(self, payload):
        headers = {"content-type": "application/json"}
        res = requests.post(self.full_url, headers=headers, data=payload, auth=(self.__username, self.__password))
        print(res.text)


#
# url     = 'http://192.168.25.17:8332'
# payload = '{"method": "getblockhash", "params":[0], "id": "foo"}'
# payload2 = {"jsonrpc": "1.0", "id":"curltest", "method": "getinfo", "params": [] }
# headers = {"content-type": "application/json"}
# res = requests.post(url, headers=headers, data=payload, auth=('bitcoinrpc', 'pass'))
# #req = urllib.request(url, payload, {'Content-Type': 'application/octet-stream'})
# #res = urllib.urlopen(req)
# print(res.text)


def main():
    bitcoin = Node(url='http://192.168.25.17', port=8332, username='bitcoinrpc', password='pass')
    print(bitcoin)
    payload = '{"method": "decoderawtransaction", "params":[], "id": "foo"}'
    bitcoin.makeCall(payload)

if __name__ == "__main__":
    main()

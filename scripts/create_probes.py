from urllib2 import Request, urlopen

values = """
  {
    "name": "testing",
    "tags": [
      "private",
      "test"
    ],
    "enabled": true
  }
"""

headers = {
  'Content-Type': 'application/json'
}
request = Request('https://url/api/v2/probes', data=values, headers=headers)

response_body = urlopen(request).read

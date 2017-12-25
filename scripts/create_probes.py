from urllib2 import Request, urlopen

values = """
  {
    "name": "JXN",
    "tags": [
      "private",
      "test"
    ],
    "enabled": true,
    "latitude": 32.2988,
    "longitude": 90.1848,
  }
"""

headers = {
  'Content-Type': 'application/json'
}
request = Request('http://localhost:3000/api/v2/probes', data=values, headers=headers)

response_body = urlopen(request).read

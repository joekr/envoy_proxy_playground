import requests as req


if __name__ == "__main__":
  url = "http://localhost:8000/service/test"
  for x in range(10):
    resp = req.get(url)
    print(resp.text)

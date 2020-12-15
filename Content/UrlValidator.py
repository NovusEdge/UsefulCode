from urllib.parse import urlparse

def uri_validator(Url):
    try:
        result = urlparse(Url)
        return all([result.scheme, result.netloc, result.path])
    except:
        return False

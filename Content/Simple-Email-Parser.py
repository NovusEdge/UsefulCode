import re

def emailParser(email):
    
    '''Regex-based method to parse out an email'''
    
    domain = re.findall(r'[\w+\.\w]+',email)[::-1][0]
    username = email[:email.find(domain) - 1]
    return 'Username: {}\nDomain: {}'.format(username,domain)

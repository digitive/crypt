#!/usr/bin/env python3
import secrets
import string
import crypt

def gen_passwrod(length):
    password = ''.join((secrets.choice(string.ascii_letters + string.digits + string.punctuation) for i in range(length)))
    salt = ''.join((secrets.choice(string.ascii_letters) for i in range(2)))
    return (password, salt)

for i in range(8):
    for j in range(20):
        password, salt = gen_passwrod(i+1)
        hash = crypt.crypt(password, salt)
        print(password + ' ' + hash)

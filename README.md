# Make HTTPS Great Again

### **Not all apps are using HTTPS.**
### **Not all hosts are using HSTS.**

Thanks to the EFForg's effort,
we can now force **desktop apps** using HTTPS as much as possible
by setting http proxy to MHGA.

MHGA uses HTTP proxy to redirect non-HTTPS connection to HTTPS
based on [EFForg/https-everywhere](https://github.com/EFForg/https-everywhere)

Example:

![Screenshot](https://user-images.githubusercontent.com/15951701/36340403-191cfd44-1417-11e8-911d-62a26f8d29c1.png)

## How to use

**This project is currently preview version. No binary releases.**

- install Golang, Git (add it to your PATH if you are using Windows)
- running `go get github.com/lightsing/MHGA`
- build it

# Feature

- [x] Redirect HTTP to HTTPS (301 for GET, 307 for POST)

# Todo

- [ ] Cookie modify
- [ ] Faster lookup
- [ ] Binary Release
- [ ] Make Package
# Make HTTPS Great Again

### **Not all apps are using HTTPS.**
### **Not all hosts are using HSTS.**

Thanks to the EFForg's effort,
we can now force **desktop apps** using HTTPS as much as possible
by setting http proxy to MHGA.

MHGA uses HTTP proxy to redirect non-HTTPS connection to HTTPS
based on [EFForg/https-everywhere](https://github.com/EFForg/https-everywhere)

Example:

![Imgur](https://i.imgur.com/YERvcHr.png)

# Feature

- [x] Redirect HTTP to HTTPS (301 for GET, 307 for POST)

# Todo

- [ ] Cookie modify
- [ ] Faster lookup
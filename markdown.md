class: center, bottom


# Building a Container from scratch

???
Present yourself
---
class: middle, right, fit-image
layout: false
background-image: url(http://localhost:8000/images/container.jpg)

# What is a container?

???
Ask audiance what they think is a container

---
class: center, middle

# Why do we need a container?


???
Talk about immutable infrastructure
talk about Pets x cattle

---
# Pets - Legacy

- We given them names, they are unique,
- We love and care and whey they get ill we nurse them.

--

???
We reconfigure the infrastructure

Configuration drift = Even with very good cfg management this will happen

---
# Cattle

- We givem them a number
- If they get ill we

???
Designed for failure
Replecable

---
class: center, middle
# Guess what? Container is not a kernel primitive!

---

# A Container is a group of functionalities provided by the Linux Kernel

- Linux Namespaces
- CGroups

---
class: top, right, fit-image
layout: false
background-image: url(http://localhost:8000/images/talkischeap.jpg)

https://www.linxlabs.com/namespaces-go-part-3-proc-mount/

https://medium.com/@teddyking/namespaces-in-go-basics-e3f0fc1ff69a

https://www.toptal.com/linux/separation-anxiety-isolating-your-system-with-linux-namespaces

https://www.youtube.com/watch?v=MHv6cWjvQjM&t=1316sw

https://lk4d4.darth.io/posts/unpriv1/

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
class: middle, right, fit-image
layout: false
background-image: url(http://localhost:8000/images/lightweight.jpg)
---

class: center, middle

# Why do we need a container?


???
- (WORKS IN MY MACHINE) AVOID CONFIGURATION DRIFT
- (HARDWARE UTILIZATION) BETTER RESOURCE UTILIZATION
- (SPIKE TRAFFIC CASE) FASTER STARTUP

---

class: center, middle
# Guess what? Container is not a kernel primitive!

---
class: middle, right, fit-image
layout: false
background-image: url(http://localhost:8000/images/fancy.jpg)

---
# What you need to know before we start

- In linuxland everything is a file

--

- Linux is the kernel and the rest is just a bunch of needed apps to run ur app

--

- A **syscall** is basically an API request to the kernel

--

- In linux a process can have childs but a child has only one parent (that's why its called, guess what, a process tree)

---
class: top, right, fit-image
layout: false
background-image: url(http://localhost:8000/images/talkischeap.jpg)

---
class: top, right, fit-image
layout: false
background-image: url(http://localhost:8000/images/docker.jpg)

???
- COLLABORATION
- IMMUTABLE INFRASTRUCTURE MADE EASY
- START BUILDING UR APP ENV

https://www.linxlabs.com/namespaces-go-part-3-proc-mount/

https://medium.com/@teddyking/namespaces-in-go-basics-e3f0fc1ff69a

https://www.toptal.com/linux/separation-anxiety-isolating-your-system-with-linux-namespaces

https://www.youtube.com/watch?v=MHv6cWjvQjM&t=1316sw

https://lk4d4.darth.io/posts/unpriv1/

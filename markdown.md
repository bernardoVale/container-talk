class: center, bottom


# Building a Container from scratch

???
Present yourself
---
class: center, bottom

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

--

_Talk is cheap, show me the code - Torvalds, Linus

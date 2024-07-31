import time
import timeit

n = 0

t = time.process_time()
while (n < 1000000000):
    n+=1
print(n)
elapse = time.process_time() - t

print(elapse)
```
# get the first m digits of n
def first_m_of_n(n, m):
    nlen = int(math.log10(n)) + 1
    return int(n / (10 ** (nlen - m)))

print(first_m_of_n(12345, 3))

```

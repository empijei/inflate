# inflate
Just a pipeline wrapper around flate algorithm

# Install 
`go get github.com/empijei/inflate`

# Usage
```sh
# Inflate
cat deflated-file | inflate > inflated-file

# Deflate (default level: 9)
cat inflated-file | inflate -d -l 8 > deflated-file
```
[useless use of cat for clarity]



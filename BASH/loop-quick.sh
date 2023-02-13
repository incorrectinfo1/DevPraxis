# simple loop for making ip list
for ip in $(seq 1 10); do echo 10.11.1.$ip; done

# simple loop with brace expansion
 for i in {1..10}; do echo 10.11.1.$i;done
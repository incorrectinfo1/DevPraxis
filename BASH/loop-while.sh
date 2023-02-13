#! /bin/bash
# similar to loop-quick but more for scripts
# tricky part is remembering -lt "Less than" and -le "less than or equal"

counter=1

while [ $counter -le 10 ]
do
    echo "10.11.1.$counter"
    ((counter++))
done
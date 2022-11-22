// Compile with gcc -m32 -o format format.c
#include <stdio.h>
#include <string.h>

int main() {
char input[151];
char *password = "ThisIsNotMyPassword\n";
printf("Enter the password stored at memory address %p:\n", &password);
fgets(input, 150, stdin);

puts("The password you entered is:");
printf(input);

if(strcmp(input, password) == 0) {
puts("Well done!");
} else {
puts("Failed :(");
}

return(0);
}
#include <stdio.h>
#include <string.h>

int main ()
{
    char input_buffer[15];
    int security_check_token = 0; // A secuirty token

    printf("\nEnter the password : \n");
    gets(input_buffer); // This is the deprecated function. It doesn't count characters entered.

    if(strcmp(input_buffer, "passw0rd"))
    {
        printf ("\n Wrong Password :( \n");
    }
    else
    {
        printf (" Correct Password! :) ");
        security_check_token = 1;
    }
    // Run Code if user is Authenticated
    if(security_check_token)
    {
        printf ("\nSecurity Token Says Privileged\n")
    }


}
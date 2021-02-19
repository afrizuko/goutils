#include <stdio.h>

void printer(const char *text);

class Logger
{
public:
    void printer(const char *text)
    {
        printf("%s\n", text);
    }
};

int main(int argc, char *argv[])
{
    printer("Hello World!");
    Logger l = Logger();
    l.printer("Hello World from l");
    return 0;
}

void printer(const char *text)
{
    printf("%s\n", text);
}
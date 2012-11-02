#include "stdio.h"

void execute(char *ptr, char *src)
{
	char input;
	char loop[30000];
	char *loopptr = loop;

	while (input = *src++)
	{
		switch (input)
		{
		case '>':
			++ptr;
			break;
		case '<':
			--ptr;
			break;
		case '+':
			++*ptr;
			break;
		case '-':
			--*ptr;
			break;
		case '.':
			putchar(*ptr);
			break;
		case ',':
			*ptr=getchar();
			break;
		case '[':

			while ((*loopptr++ = *src++) != ']')
			{
			}

			printf("%s\n", loop);

			while (*ptr)
			{
				execute(ptr, loop);
			}

			break;
		case ']':
			break;

		default :
			break;
		}
	}
}

void main(int argc, char *argv[])
{
	if (argc < 2)
	{
		printf("brainfuck\n");

		return ;
	}

	char array[30000];
	
	execute(array, argv[1]);

	printf("\n");
}

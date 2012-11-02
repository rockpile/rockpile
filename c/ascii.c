#include "stdio.h"

void main(int argc, char *argv[])
{
	if (argc == 2)
	{
		for (int i = 0; i < strlen(argv[1]); ++i)
		{
			printf("%d ", argv[1][i]);
		}

		printf("\n");
	}
}

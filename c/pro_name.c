#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>

extern char **environ;
char *last;

void initProcTitle(int argc, char **argv)
{
	size_t size = 0;

	for (int i = 0; environ[i]; ++i) {
		size += strlen(environ[i] + 1);
	}

	char *raw = new char[size];

	for (int i = 0; environ[i]; ++i) {
		memcpy(raw, environ[i], strlen(environ[i] + 1));
		environ[i] = raw;
		raw += strlen(environ[i]) + 1;
	}

	last = argv[0];

	for (int i = 0; i < argc; ++i) {
		last += strlen(argv[i]) + 1;
	}

	for (int i = 0; environ[i]; ++i) {
		last += strlen(environ[i]) + 1;
	}
}

void setProcTitle(int argc, char **argv, const char *title)
{
	argv[1] = 0;
	char *p = argv[0];
	memset(p, 0x00, last - p);
	strncpy(p, title, last - p);
}

int main(int argc, char *argv[])
{
	initProcTitle(argc, argv);
	setProcTitle(argc, argv, "master");
	sleep(60);
	
	return 0;
}

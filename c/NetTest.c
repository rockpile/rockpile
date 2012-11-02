#include <stdio.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>

int main(int argc, char* argv[])
{
	int sock, c;
	sock = socket(AF_INET, SOCK_STREAM, 0);

	struct sockaddr_in service, client;
	memset((void*)&service, 0, sizeof(struct sockaddr));
	service.sin_family = AF_INET;
   	service.sin_addr.s_addr = inet_addr("192.168.1.149");
        service.sin_port = htons(27015);

	bind(sock, (struct sockaddr *)&service, sizeof(struct sockaddr));

	listen(sock, 51);

	socklen_t size = sizeof(struct sockaddr);
	c = accept(sock, (struct sockaddr *)&client, &size);
	
	char buf[4] = {0};
	int32_t i = 0;
	memcpy((void*)&i, (void*)buf, 4);
	recv(c, (void*)&i, 4, 0);
	printf("%d\n", buf[0]);
	printf("%d\n", i);
	send(c, (const char*)&i, 4, 0);

	return 0;
}

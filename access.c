#include <unistd.h>
#include <fcntl.h>
#include <stdlib.h>
#include <stdio.h>

int main(int argc, char *argv[]) {
	int ret = 0;
	if (argc != 2) {
		printf("usage: access_c PATH\n");
		exit(1);
	}
	faccessat(AT_FDCWD, argv[1], F_OK, AT_SYMLINK_NOFOLLOW);
	perror("");
}

package main

import (
	"strings"
	"unsafe"
)

/*
#define _GNU_SOURCE
#include <arpa/inet.h>
#include <sys/socket.h>
#include <netdb.h>
#include <ifaddrs.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <linux/if_link.h>
#include <string.h>
#include <limits.h>
char* getifaces()
{
    struct ifaddrs *ifaddr, *ifa;
    int family, s;
    char host[NI_MAXHOST];
    int newlen = 0;
    char* interfaces = (char*) malloc(4);
    char* joined = NULL;
    if (getifaddrs(&ifaddr) == -1) {
        free(interfaces);
        perror("getifaddrs");
        exit(EXIT_FAILURE);
    }
    for (ifa = ifaddr; ifa != NULL; ifa = ifa->ifa_next) {
        if (ifa->ifa_addr == NULL)
            continue;
        family = ifa->ifa_addr->sa_family;
        if (family == AF_INET || family == AF_INET6) {
            s = getnameinfo(ifa->ifa_addr,
                    (family == AF_INET) ? sizeof(struct sockaddr_in) :
                                            sizeof(struct sockaddr_in6),
                    host, NI_MAXHOST,
                    NULL, 0, NI_NUMERICHOST);
            newlen = strlen(ifa->ifa_name) + 2;
            joined = malloc(newlen * sizeof(char*));
            strcpy(joined, ifa->ifa_name);
            strcat(joined, ",");
            interfaces = realloc(interfaces, newlen * sizeof(char*));
            strcat(interfaces, joined);
            free(joined);
        }
    }
    freeifaddrs(ifaddr);
    return interfaces;
}
*/
import "C"

// getInterfaces binds to the C getifaces() function.
func (tcp TCPIP) getInterfaces() []string {
	ifacesPTR := C.getifaces()
	var ifaces string = C.GoString(ifacesPTR)
	defer C.free(unsafe.Pointer(ifacesPTR))
	var interfaces []string
	for _, adapter := range strings.Split(ifaces, ",") {
		if len(adapter) < 1 {
			continue
		}
		isDup := false
		for _, ifaceName := range interfaces {
			if ifaceName != adapter {
				continue
			}
			isDup = true
			break
		}
		if !isDup {
			interfaces = append(interfaces, adapter)
		}
	}
	return interfaces
}

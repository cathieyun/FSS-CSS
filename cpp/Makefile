CFLAGS=-fomit-frame-pointer -maes -std=c++11 -msse4.2 -fno-strict-aliasing -pedantic -Wall -Wextra -Wunreachable-code -Wmissing-declarations -Wunused-function -Wno-overlength-strings -Wno-sign-compare -Wno-deprecated-declarations -O3
INC=-I/usr/local/opt/openssl/include -Iinclude
LIB=-L/usr/local/opt/openssl/lib -lcrypto -lssl -lgmpxx -lgmp

fss: fss.o
	g++ -o fss fss.o $(LIB)

fss.o: fss.cpp
	g++ $(CFLAGS) $(INC) -c -o fss.o fss.cpp

clean: 
	rm -f fss *.o
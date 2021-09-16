so:
	gcc -c -fPIC -o mystack.o mystack.c
	gcc -shared -o libmystack.so mystack.o

run:
	docker run outlet_client


lib :
	gcc -c ./oledc/oled.c -lwiringPi
	ar cru ./oledc/liboledc.a ./oledc/oled.o

run: lib
	go run .


lib :
	gcc -c ./oledc/oled.c -o ./oledc/oled.o -lwiringPi
	ar cru ./oledc/liboledc.a ./oledc/oled.o

run: lib
	go run .


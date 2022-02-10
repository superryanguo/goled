lib :
	gcc -c ./oledc/oled.c -o ./oledc/oled.o -lwiringPi
	ar cru ./oledc/liboledc.a ./oledc/oled.o

run: lib
	go run .

fax: lib
	go run ./face/face.go 0 data/haarcascade_frontalface_default.xml


OBJ = cspout_go.o SpoutDirectX.o SpoutGLDXinterop.o SpoutGLextensions.o SpoutMemoryShare.o SpoutReceiver.o SpoutSDK.o SpoutSender.o SpoutSenderNames.o SpoutSharedMemory.o

COPYOBJ = SpoutCopy.o

CC = g++

SPOUT2 = ../../../Spout2
SPOUT2SRC = $(SPOUT2)/SpoutSDK/Source

%.o : $(SPOUT2SRC)/%.cpp
	$(CC) -c -I$(SPOUT2SRC) -Wno-int-to-pointer-cast -Wno-pointer-arith -Wno-conversion-null $(CFLAGS) $(CPPFLAGS) $< -o $@

default: libspout.a

clean:
	-rm -f *.o *.a

libspout.a: $(OBJ) $(COPYOBJ)
	-rm -f libspout.a
	ar q libspout.a $(OBJ) $(COPYOBJ)


Spout_go.o : Spout_go.cpp
	$(CC) -c -I$(SPOUT2SRC) $(CFLAGS) $(CPPFLAGS) $< -o $@

SpoutCopy.o : $(SPOUT2SRC)/SpoutCopy.cpp
	$(CC) -c -msse4.1 $(CFLAGS) $(CPPFLAGS) $< -o $@

MAKE=mingw32-make.exe

cd internal/libspout
$MAKE clean

rm -f cmd/gospout-democlient/*.exe
rm -f cmd/gospout-demoserver/*.exe

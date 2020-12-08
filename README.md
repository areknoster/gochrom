# gochrom
Rozwiązanie 3 zadania laboratoryjnego z numerem 0

Dane: Arkadiusz Noster, 298890
###Uruchomienie
1. Należy mieć zainstalowany język go w wersji conajmniej 1.15 (można pobrać ze [strony języka](https://golang.org/dl/)).
2. w korzeniu projektu należy użyć komendy `go run .`. 
Potrzebne zależności powinny pobrać się automatycznie i program powinien się uruchomić. 
Alternatywnie można zbudować projekt do statycznego pliku binarnego: `go build -o nazwaPlikuWynikowego . ` a następnie ten plik uruchomić
### Użytkowanie
Program jest bardzo intuicyjny i zgodny z opisem zadania. Dolny panel posiada wykres z czterema liniami. 
Kolorowe linie odpowiadają czułości na dane częstotliwości czopków ludzkiego oka. Czarna linia odpowiada widmu koloru.

Wykres widma można zmieniać naciskając prawym przyciskiem myszy na przestrzeń wykresu. 
Wartości długości fali światła są ograniczone do zakresu [400,700], gdyż z takiego zakresu program korzysta do wyliczenia wynikowego koloru.

Modyfikacja wykresu widmowego powoduje przesuwanie znacznika koloru na diagramie w górnym panelu.

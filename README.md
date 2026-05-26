### E-biznes

Zadanie 1 Docker
 
:white_check_mark: 3.0 obraz ubuntu z Pythonem w wersji 3.10 </br>
[Link do commita ](https://github.com/lukasz-malec/E-biznes/commit/67c0c9e0842e71619dbe89cf70901f5cc308c453) </br>
 [Link do obrazu](https://hub.docker.com/repository/docker/lukasz493/zad1_3.0/general)</br></br>
:white_check_mark: 3.5 obraz ubuntu:24.02 z Javą w wersji 8 oraz Kotlinem 
</br>
[Link do commita ](https://github.com/lukasz-malec/E-biznes/commit/b75103711252075b001e70bbd9a7dfa5e21adff8) </br>
 [Link do obrazu](https://hub.docker.com/repository/docker/lukasz493/zad1_3.5/general)</br></br>
:white_check_mark: 4.0 do powyższego należy dodać najnowszego Gradle’a oraz paczkę JDBC 
SQLite w ramach projektu na Gradle (build.gradle) </br>
[Link do commita ](https://github.com/lukasz-malec/E-biznes/commit/3e7d8b4792251b93a880a47d6f865f274d0e8c44) </br>
[Link do obrazu](https://hub.docker.com/repository/docker/lukasz493/zada1_4.0/general)</br>

:white_check_mark: 4.5 stworzyć przykład typu HelloWorld oraz uruchomienie aplikacji
przez CMD oraz gradle </br>
[Link do commita ](https://github.com/lukasz-malec/E-biznes/commit/adf0bd5457fcc841ba6430bfbd3e7a464e468557) </br>
[Link do obrazu](https://hub.docker.com/repository/docker/lukasz493/zad1_4.5/general)</br>

❌5.0 dodać konfigurację docker-compose

Termin: 25.03

##

Zadanie 2 Scala

Należy stworzyć aplikację na frameworku Play lub Scalatra.

:white_check_mark:3.0 Należy stworzyć kontroler do Produktów </br>
[Link do commita ](https://github.com/lukasz-malec/E-biznes/commit/9bedb567429ba1deda34687da44ddd93b7e28aca) </br>

:white_check_mark:3.5 Do kontrolera należy stworzyć endpointy zgodnie z CRUD - dane
pobierane z listy</br>
[Link do commita ](https://github.com/lukasz-malec/E-biznes/commit/806bea2ff50d8b90c8e2ce3856bc693d3095c1b5) </br>

❌4.0 Należy stworzyć kontrolery do Kategorii oraz Koszyka + endpointy
zgodnie z CRUD
</br></br>
❌4.5 Należy aplikację uruchomić na dockerze (stworzyć obraz) oraz dodać
skrypt uruchamiający aplikację via ngrok</br></br>
❌5.0 Należy dodać konfigurację CORS dla dwóch hostów dla metod CRUD

Kontrolery mogą bazować na listach zamiast baz danych. CRUD: show all,
show by id (get), update (put), delete (delete), add (post).


Demo: [Link do nagrania](https://github.com/lukasz-malec/E-biznes/blob/main/ZADANIE2/scala_nagranie.mp4)

##

Zadanie 3 Kotlin

:white_check_mark:3.0 Należy stworzyć aplikację kliencką w Kotlinie we frameworku Ktor,
która pozwala na przesyłanie wiadomości na platformę Discord</br>
[Link do commita ](https://github.com/lukasz-malec/E-biznes/commit/2708fd07b365b7a7fddc1191086a4e13b8bcf11f) </br>

❌3.5 Aplikacja jest w stanie odbierać wiadomości użytkowników z
platformy Discord skierowane do aplikacji (bota)</br></br>

❌4.0 Zwróci listę kategorii na określone żądanie użytkownika</br></br>

❌4.5 Zwróci listę produktów wg żądanej kategorii</br></br>

❌5.0 Aplikacja obsłuży dodatkowo jedną z platform: Slack lub Messenger</br></br>

Aplikację należy uruchomić na dockerze.

Termin: 8.04

Demo: [Link do nagrania](https://github.com/lukasz-malec/E-biznes/tree/main/ZADANIE3/nagranie)

## 

Zadanie 4 GO

Należy stworzyć projekt w echo w Go. Należy wykorzystać gorm do
stworzenia kilka modeli, gdzie pomiędzy dwoma musi być relacja. Należy
zaimplementować proste endpointy do dodawania oraz wyświetlania danych
za pomocą modeli. Jako bazę danych można wybrać dowolną, sugerowałbym
jednak pozostać przy sqlite.

:white_check_mark: 3.0 Należy stworzyć aplikację we frameworki echo w j. Go, która będzie
miała kontroler Produktów zgodny z CRUD</br>
[Link do commita ](https://github.com/lukasz-malec/E-biznes/commit/ed416317a840ca164b77f32096bbb01371b47a31) </br>



❌ 3.5 Należy stworzyć model Produktów wykorzystując gorm oraz
wykorzystać model do obsługi produktów (CRUD) w kontrolerze (zamiast
listy)

❌ 4.0 Należy dodać model Koszyka oraz dodać odpowiedni endpoint

❌ 4.5 Należy stworzyć model kategorii i dodać relację między kategorią,
a produktem

❌ 5.0 pogrupować zapytania w gorm’owe scope'y

Termin: 15.04



## 
Zadanie 5 Frontend

Należy stworzyć aplikację kliencką wykorzystując bibliotekę React.js.
W ramach projektu należy stworzyć trzy komponenty: Produkty, Koszyk
oraz Płatności. Koszyk oraz Płatności powinny wysyłać do aplikacji
serwerowej dane, a w Produktach powinniśmy pobierać dane o produktach
z aplikacji serwerowej. Aplikacja serwera w jednym z trzech języków:
Kotlin, Scala, Go. Dane pomiędzy wszystkimi komponentami powinny być
przesyłane za pomocą React hooks.d

:white_check_mark:3.0 W ramach projektu należy stworzyć dwa komponenty: Produkty oraz
Płatności; Płatności powinny wysyłać do aplikacji serwerowej dane, a w
Produktach powinniśmy pobierać dane o produktach z aplikacji
serwerowej;</br>
[Link do commita](https://github.com/lukasz-malec/E-biznes/commit/4a2a4cd107092e0b38e3d6af844d4c46ca92a804)</br>

:white_check_mark: 3.5 Należy dodać Koszyk wraz z widokiem; należy wykorzystać routing </br>
[Link do commit](https://github.com/lukasz-malec/E-biznes/commit/825cd8fdee1438fee5d5d637851d2390ac28555d)<br>

:white_check_mark: 4.0 Dane pomiędzy wszystkimi komponentami powinny być przesyłane za
pomocą React hooks </br>
[Link do commita](https://github.com/lukasz-malec/E-biznes/commit/9efe576e2a6db7f3448b24fde3255035413cdb14)</br>

:white_check_mark: 4.5 Należy dodać skrypt uruchamiający aplikację serwerową oraz
kliencką na dockerze via docker-compose </br>
[Link do commita ](https://github.com/lukasz-malec/E-biznes/commit/76ab195f4f541aaffad0d42c746e7958d182b4b5)</br>

:white_check_mark: 5.0 Należy wykorzystać axios’a oraz dodać nagłówki pod CORS</br>
[Link do commita](https://github.com/lukasz-malec/E-biznes/commit/675ab9a9367cc6ce88e33181d0fe79d20d15a541)</br></br>

[Demo](https://github.com/user-attachments/assets/fe793182-cd06-4d51-a0e2-26193d5cb864)


##
Zadanie 6 Testy
Należy stworzyć 20 przypadków testowych w jednym z rozwiązań:

- Cypress JS (JS)
- Selenium (Kotlin, Python, Java, JS, Go, Scala)

Testy mają w sumie zawierać minimum 50 asercji (3.5). Mają również
uruchamiać się na platformie Browserstack (5.0). Proszę pamiętać o
stworzeniu darmowego konta via https://education.github.com/pack.

:white_check_mark: 3.0 Należy stworzyć 20 przypadków testowych w CypressJS lub Selenium
(Kotlin, Python, Java, JS, Go, Scala)</br>
[Link do commita ](https://github.com/lukasz-malec/E-biznes/commit/963c06389f53575669c0744494d4cf91e8a03ecb)</br></br>
:white_check_mark: 3.5 Należy rozszerzyć testy funkcjonalne, aby zawierały minimum 50
asercji</br>
[Link do commita](https://github.com/lukasz-malec/E-biznes/commit/4311754146964860ecda4fcf8d4eccf341e4fb47)</br></br>
:white_check_mark: 4.0 Należy stworzyć testy jednostkowe do wybranego wcześniejszego
projektu z minimum 50 asercjami</br>
[Link do commita](https://github.com/lukasz-malec/E-biznes/commit/30cb3408dcacd6546565a670661ab2ff37bf9581)</br></br>
:white_check_mark: 4.5 Należy dodać testy API, należy pokryć wszystkie endpointy z
minimum jednym scenariuszem negatywnym per endpoint</br>
[Link do commita](https://github.com/lukasz-malec/E-biznes/commit/7f2c8ee14d8229798e1b9dd611a057705520aff5)</br></br>
<<<<<<< HEAD
❌5.0 Należy uruchomić testy funkcjonalne na Browserstacku</br></br>

##
Zadanie 7
Należy dodać projekt aplikacji klienckiej oraz serwerowej (jeden
branch, dwa repozytoria) do Sonara w wersji chmurowej
(https://sonarcloud.io/). Należy poprawić aplikacje uzyskując 0 bugów,
0 zapaszków, 0 podatności, 0 błędów bezpieczeństwa. Dodatkowo należy
dodać widżety sonarowe do README w repozytorium dane projektu z
wynikami.

3.0 Należy dodać litera do odpowiedniego kodu aplikacji serwerowej w
hookach gita
3.5 Należy wyeliminować wszystkie bugi w kodzie w Sonarze (kod
aplikacji serwerowej)
4.0 Należy wyeliminować wszystkie zapaszki w kodzie w Sonarze (kod
aplikacji serwerowej)
4.5 Należy wyeliminować wszystkie podatności oraz błędy bezpieczeństwa
w kodzie w Sonarze (kod aplikacji serwerowej)
5.0 Należy wyeliminować wszystkie błędy oraz zapaszki w kodzie
aplikacji klienckiej

https://golangci-lint.run/
https://github.com/pinterest/ktlint
https://scalameta.org/scalafmt/docs/installation.html
=======
:white_check_mark: 5.0 Należy uruchomić testy funkcjonalne na Browserstacku</br>



https://github.com/user-attachments/assets/350d32d6-afc0-45bd-8ccb-a311519cae52


>>>>>>> 3388d996387e258092e3093270594b9a95e497b3

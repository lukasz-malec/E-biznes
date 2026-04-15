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

:white_check_mark:3.0 Należy stworzyć aplikację we frameworki echo w j. Go, która będzie
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
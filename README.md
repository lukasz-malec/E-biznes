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

5.0 dodać konfigurację docker-compose

Termin: 25.03

##

Zadanie 2 Scala

Należy stworzyć aplikację na frameworku Play lub Scalatra.

3.0 Należy stworzyć kontroler do Produktów
3.5 Do kontrolera należy stworzyć endpointy zgodnie z CRUD - dane
pobierane z listy
4.0 Należy stworzyć kontrolery do Kategorii oraz Koszyka + endpointy
zgodnie z CRUD
4.5 Należy aplikację uruchomić na dockerze (stworzyć obraz) oraz dodać
skrypt uruchamiający aplikację via ngrok
5.0 Należy dodać konfigurację CORS dla dwóch hostów dla metod CRUD

Kontrolery mogą bazować na listach zamiast baz danych. CRUD: show all,
show by id (get), update (put), delete (delete), add (post).
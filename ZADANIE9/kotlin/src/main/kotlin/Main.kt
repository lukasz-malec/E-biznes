import io.ktor.client.*
import io.ktor.client.engine.cio.*
import io.ktor.client.request.*
import io.ktor.client.statement.*
import io.ktor.http.*
import kotlinx.coroutines.runBlocking
import io.github.cdimascio.dotenv.dotenv

fun main() = runBlocking {
    val client = HttpClient(CIO) {
        engine {
            requestTimeout = 120_000  // 2 minuty
        }
    }

    val dotenv = dotenv()
    val webhookUrl = dotenv["WEBHOOK_URL"]
    val serwisUrl = "http://localhost:8000/ask"

    println("Wpisz wiadomosc (exit aby zakonczyc)")

    while (true) {
        val message = readln()

        if (message.lowercase() == "exit") {
            println("Zamykanie aplikacji...")
            break
        }

        try {
            println("Czekam na odpowiedz...")

            val aiResponse = client.post(serwisUrl) {
                contentType(ContentType.Application.Json)
                setBody("""{"message": "$message"}""")
            }
            val body = aiResponse.bodyAsText()
            val answer = body
                .substringAfter("\"answer\":\"")
                .substringBefore("\"}")
                .replace("\\n", "\n")

            println("Bot: $answer")

            client.post(webhookUrl) {
                contentType(ContentType.Application.Json)
                setBody("""{"content": "$answer"}""")
            }
            println("Wyslano na Discord!")

        } catch (e: Exception) {
            println("Blad: ${e.message}")
        }
    }

    client.close()
}
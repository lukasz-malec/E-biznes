import io.ktor.client.*
import io.ktor.client.engine.cio.*
import io.ktor.client.request.*
import io.ktor.http.*
import kotlinx.coroutines.runBlocking
import io.github.cdimascio.dotenv.dotenv

fun main() = runBlocking {
    val client = HttpClient(CIO)

    val dotenv = dotenv()
    val webhookUrl = dotenv["WEBHOOK_URL"]

    println("Wpisz wiadomosc (exit aby zakonczyc)")

    while (true) {
        val message = readln()

        if (message.lowercase() == "exit") {
            println("Zamykanie aplikacji...")
            break
        }

        try {
            val response = client.post(webhookUrl) {
                contentType(ContentType.Application.Json)
                setBody("""
                    {
                        "content": "$message"
                    }
                """.trimIndent())
            }
            println("Wyslano!")
        } catch (e: Exception) {
            println("Blad: ${e.message}")
        }
    }

    client.close()
}

import org.eclipse.jetty.server.Server
import org.eclipse.jetty.webapp.WebAppContext

object Main {
  def main(args: Array[String]): Unit = {
    val server = new Server(8080)

    val context = new WebAppContext()
    context.setContextPath("/")
    context.setResourceBase("src/main/webapp")

    context.setInitParameter(
      "org.eclipse.jetty.servlet.Default.dirAllowed",
      "false"
    )

    server.setHandler(context)

    server.start()
    println("Server: http://localhost:8080")
    server.join()
  }
}
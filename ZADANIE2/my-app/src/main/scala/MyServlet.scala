import org.scalatra._

class MyServlet extends ScalatraServlet {

  get("/") {
    "Hello Scalatrasudhjkdh!"
  }

  get("/hello/:name") {
    s"Hello ${params("name")}!"
  }
}
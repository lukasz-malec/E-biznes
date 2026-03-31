import org.scalatra._
import org.json4s.DefaultFormats
import org.json4s.jackson.Serialization.write
import scala.collection.mutable.ListBuffer


class ProductController extends ScalatraServlet {

  implicit val formats = DefaultFormats

  private val products = ListBuffer(
    Product(1, "Laptop", 3500.0),
    Product(2, "Telefon", 2000.0),
    Product(3, "Klawiatura", 150.0),
    Product(4, "Monitor", 600.0),
    Product(5, "Dysk SSD", 850.0)
  )

  get("/products") {
    contentType = "application/json"
    write(products)
  }

  get("/") {
    "dziala"
  }

}
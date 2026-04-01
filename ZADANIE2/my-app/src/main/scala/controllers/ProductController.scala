import org.scalatra._
import org.json4s.{DefaultFormats, Formats}
import org.json4s.jackson.Serialization.write
import org.json4s.jackson.JsonMethods._
import scala.collection.mutable.ListBuffer

class ProductController extends ScalatraServlet {

  implicit val formats: Formats = DefaultFormats

  private val products = ListBuffer(
    Product(1, "Laptop", 3500.0),
    Product(2, "Telefon", 2000.0),
    Product(3, "PS5", 2700.0)
  )

    // strona domowa
    get("/"){
      "działa"
    }



  // CRUD 


  // show all
  get("/products") {
    contentType = "application/json"
    write(products)
  }


  // show by id
  get("/products/:id") {
    contentType = "application/json"
    val id = params("id").toInt
    write(products.find(_.id == id).getOrElse(halt(404)))
  }


  // add
  post("/products") {
    val json = parse(request.body)
    val product = json.extract[Product]

    products += product
    contentType = "application/json"
    write(product)
  }


  // update
  put("/products/:id") {
    val id = params("id").toInt
    val json = parse(request.body)
    val updated = json.extract[Product]

    val idx = products.indexWhere(_.id == id)
    if (idx == -1) halt(404)

    products.update(idx, updated)
    contentType = "application/json"
    write(updated)
  }

  // delete
  delete("/products/:id") {
  val id = params("id").toInt
  products --= products.filter(_.id == id)
  "deleted"
  }
}




// 


// POST
// {
//   "id": 4,
//   "name" : "Myszka",
//   "price" : 200.0
// }



// PUT
//{
//   "id": 3,
//   "name": "PS5 PRO",
//   "price": 3200.0
// }